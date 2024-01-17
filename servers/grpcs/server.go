package grpcs

import (
	"context"
	"errors"
	"fmt"
	"github.com/fullstorydev/grpchan/httpgrpc"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/pubgo/funk/errors/errutil"
	"github.com/pubgo/funk/generic"
	"github.com/pubgo/funk/proto/errorpb"
	"github.com/pubgo/lava/pkg/grpcutil"
	"github.com/pubgo/lava/pkg/httputil"
	"github.com/pubgo/lava/pkg/larking"
	"github.com/pubgo/lava/pkg/wsproxy"
	"github.com/pubgo/opendoc/opendoc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"

	"github.com/fullstorydev/grpchan"
	_ "github.com/fullstorydev/grpchan/httpgrpc"
	"github.com/fullstorydev/grpchan/inprocgrpc"
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/pubgo/funk/assert"
	"github.com/pubgo/funk/async"
	"github.com/pubgo/funk/config"
	"github.com/pubgo/funk/log"
	"github.com/pubgo/funk/recovery"
	"github.com/pubgo/funk/running"
	"github.com/pubgo/funk/stack"
	"github.com/pubgo/funk/vars"
	"github.com/pubgo/funk/version"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"

	"github.com/pubgo/lava/core/debug"
	"github.com/pubgo/lava/core/lifecycle"
	"github.com/pubgo/lava/core/metrics"
	"github.com/pubgo/lava/core/signal"
	"github.com/pubgo/lava/internal/consts"
	"github.com/pubgo/lava/internal/logutil"
	"github.com/pubgo/lava/internal/middlewares/middleware_accesslog"
	"github.com/pubgo/lava/internal/middlewares/middleware_metric"
	"github.com/pubgo/lava/internal/middlewares/middleware_recovery"
	"github.com/pubgo/lava/internal/middlewares/middleware_service_info"
	"github.com/pubgo/lava/lava"
)

func New() lava.Service { return newService() }

func newService() *serviceImpl {
	return &serviceImpl{
		reg: make(grpchan.HandlerMap),
		cc:  new(inprocgrpc.Channel),
	}
}

var _ lava.Service = (*serviceImpl)(nil)

type serviceImpl struct {
	lc         lifecycle.Getter
	httpServer *fiber.App
	grpcServer *grpc.Server
	log        log.Logger
	reg        grpchan.HandlerMap
	cc         *inprocgrpc.Channel
	initList   []func()
	conf       *Config
}

func (s *serviceImpl) Run() {
	defer s.stop()
	s.start()
	signal.Wait()
}

func (s *serviceImpl) Start() { s.start() }
func (s *serviceImpl) Stop()  { s.stop() }

func (s *serviceImpl) DixInject(
	handlers []lava.GrpcRouter,
	httpRouters []lava.HttpRouter,
	dixMiddlewares []lava.Middleware,
	getLifecycle lifecycle.Getter,
	lifecycle lifecycle.Lifecycle,
	metric metrics.Metric,
	log log.Logger,
	conf *Config,
	docs []*opendoc.Swagger,
	empty []*lava.EmptyRouter,
) {
	_ = empty
	s.conf = conf
	if conf.HttpPort == nil {
		conf.HttpPort = generic.Ptr(running.HttpPort)
	}

	if conf.GrpcPort == nil {
		conf.GrpcPort = generic.Ptr(running.GrpcPort)
	}

	if conf.BaseUrl == "" {
		conf.BaseUrl = "/" + version.Project()
	}

	fiber.SetParserDecoder(fiber.ParserConfig{
		IgnoreUnknownKeys: true,
		ZeroEmpty:         true,
		ParserType:        parserTypes,
	})

	s.lc = getLifecycle

	conf = config.MergeR(defaultCfg(), conf).Unwrap()
	conf.BaseUrl = "/" + strings.Trim(conf.BaseUrl, "/")

	var doc = opendoc.New(func(swag *opendoc.Swagger) {
		swag.Config.Title = "service title "
		swag.Description = "this is description"
		swag.License = &opendoc.License{
			Name: "Apache License 2.0",
			URL:  "https://github.com/pubgo/opendoc/blob/master/LICENSE",
		}

		swag.Contact = &opendoc.Contact{
			Name:  "barry",
			URL:   "https://github.com/pubgo/opendoc",
			Email: "kooksee@163.com",
		}

		swag.TermsOfService = "https://github.com/pubgo"
	})
	if len(docs) > 0 {
		doc = docs[0]
	}
	doc.SetRootPath(conf.BaseUrl)

	middlewares := lava.Middlewares{
		middleware_service_info.New(),
		middleware_metric.New(metric),
		middleware_accesslog.New(log),
		middleware_recovery.New(),
	}
	middlewares = append(middlewares, dixMiddlewares...)

	log = log.WithName("grpc-server")
	s.log = log

	httpServer := fiber.New(fiber.Config{
		EnableIPValidation: true,
		EnablePrintRoutes:  conf.EnablePrintRoutes,
		AppName:            version.Project(),
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			if err == nil {
				return nil
			}

			code := fiber.StatusBadRequest
			errPb := errutil.ParseError(err)
			if errPb == nil || errPb.Code.Code == 0 {
				return nil
			}

			errPb.Trace.Operation = ctx.Route().Path
			code = errutil.GrpcCodeToHTTP(codes.Code(errPb.Code.Code))
			ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
			return ctx.Status(code).JSON(errPb.Code)
		},
	})

	if conf.EnableCors {
		httpServer.Use(cors.New(cors.Config{
			AllowOriginsFunc: func(origin string) bool {
				return true
			},
			AllowOrigins: "*",
			AllowMethods: strings.Join([]string{
				fiber.MethodGet,
				fiber.MethodPost,
				fiber.MethodPut,
				fiber.MethodDelete,
				fiber.MethodPatch,
				fiber.MethodHead,
				fiber.MethodOptions,
			}, ","),
			AllowHeaders:     "",
			AllowCredentials: true,
			ExposeHeaders:    "",
			MaxAge:           0,
		}))
	}

	httpServer.Mount("/debug", debug.App())

	app := fiber.New()
	app.Use(handlerHttpMiddle(middlewares))
	for _, h := range httpRouters {
		//srv := doc.WithService()
		//for _, an := range h.Annotation() {
		//	switch a := an.(type) {
		//	case *annotation.Openapi:
		//		if a.ServiceName != "" {
		//			srv.SetName(a.ServiceName)
		//		}
		//	}
		//}

		var g = app.Group("", handlerHttpMiddle(h.Middlewares()))
		h.Router(&lava.Router{
			R:   g,
			Doc: doc.WithService(),
		})

		if m, ok := h.(lava.Close); ok {
			lifecycle.BeforeStop(m.Close)
		}

		if m, ok := h.(lava.Init); ok {
			s.initList = append(s.initList, m.Init)
		}
	}

	for _, handler := range handlers {
		//srv := doc.WithService()
		//for _, an := range h.Annotation() {
		//	switch a := an.(type) {
		//	case *annotation.Openapi:
		//		if a.ServiceName != "" {
		//			srv.SetName(a.ServiceName)
		//		}
		//	}
		//}

		h, ok := handler.(lava.HttpRouter)
		if !ok {
			continue
		}

		var g = app.Group("", handlerHttpMiddle(h.Middlewares()))
		h.Router(&lava.Router{
			R:   g,
			Doc: doc.WithService(),
		})

		if m, ok := h.(lava.Close); ok {
			lifecycle.BeforeStop(m.Close)
		}

		if m, ok := h.(lava.Init); ok {
			s.initList = append(s.initList, m.Init)
		}
	}

	httpServer.Mount(conf.BaseUrl, app)

	grpcGateway := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.HTTPBodyMarshaler{
			Marshaler: &runtime.JSONPb{
				MarshalOptions: protojson.MarshalOptions{
					EmitUnpopulated: true,
				},
				UnmarshalOptions: protojson.UnmarshalOptions{
					DiscardUnknown: true,
				},
			},
		}),
		runtime.SetQueryParameterParser(new(DefaultQueryParser)),
		runtime.WithIncomingHeaderMatcher(func(s string) (string, bool) {
			return strings.ToLower(s), true
		}),
		runtime.WithOutgoingHeaderMatcher(func(s string) (string, bool) {
			return strings.ToUpper(s), true
		}),
		runtime.WithMetadata(func(ctx context.Context, request *http.Request) metadata.MD {
			path, ok := runtime.HTTPPathPattern(ctx)
			if !ok {
				return nil
			}
			return metadata.Pairs("http_path", path, "http_method", request.Method, "http_url", request.URL.Path)
		}),
		runtime.WithErrorHandler(func(ctx context.Context, mux *runtime.ServeMux, marshal runtime.Marshaler, w http.ResponseWriter, request *http.Request, err error) {
			md, ok := runtime.ServerMetadataFromContext(ctx)
			if ok && w != nil {
				for k, v := range md.HeaderMD {
					for i := range v {
						w.Header().Add(k, v[i])
					}
				}

				for k, v := range md.TrailerMD {
					for i := range v {
						w.Header().Add(k, v[i])
					}
				}
			}

			sts, ok := status.FromError(err)
			if !ok || sts == nil {
				runtime.DefaultHTTPErrorHandler(ctx, mux, marshal, w, request, err)
				return
			}

			w.Header().Set("Content-Type", marshal.ContentType(sts))

			const fallback = `{"code": 13, "message": "failed to marshal error message"}`
			var pb *errorpb.ErrCode
			if len(sts.Details()) > 0 {
				if code, ok := sts.Details()[0].(*errorpb.Error); ok {
					pb = code.Code
				}
			} else {
				pb = &errorpb.ErrCode{
					Message:    sts.Message(),
					StatusCode: errorpb.Code(sts.Code()),
					Name:       "lava.grpc.status",
					Details:    sts.Proto().Details,
				}
			}

			buf, mErr := marshal.Marshal(pb)
			if mErr != nil {
				grpclog.Infof("Failed to marshal error message %q: %v", s, mErr)
				w.WriteHeader(http.StatusInternalServerError)
				if _, err := io.WriteString(w, fallback); err != nil {
					grpclog.Infof("Failed to write response: %v", err)
				}
				return
			}

			w.WriteHeader(runtime.HTTPStatusFromCode(codes.Code(pb.Code)))
			if _, err := w.Write(buf); err != nil {
				grpclog.Infof("Failed to write response: %v", err)
			}
		}),
	)

	srvMidMap := make(map[string][]lava.Middleware)
	for _, h := range handlers {
		desc := h.ServiceDesc()
		assert.If(desc == nil, "desc is nil")

		srvMidMap[desc.ServiceName] = append(srvMidMap[desc.ServiceName], middlewares...)
		srvMidMap[desc.ServiceName] = append(srvMidMap[desc.ServiceName], h.Middlewares()...)

		if m, ok := h.(lava.Close); ok {
			lifecycle.BeforeStop(m.Close)
		}

		if m, ok := h.(lava.Initializer); ok {
			s.initList = append(s.initList, m.Initialize)
		}

		if m, ok := h.(lava.Init); ok {
			s.initList = append(s.initList, m.Init)
		}

		s.reg.RegisterService(desc, h)
		s.cc.RegisterService(desc, h)
		if m, ok := h.(lava.GrpcGatewayRouter); ok {
			assert.Exit(m.RegisterGateway(context.Background(), grpcGateway, s.cc))
		}
	}

	s.cc = s.cc.WithServerUnaryInterceptor(handlerUnaryMiddle(srvMidMap))
	s.cc = s.cc.WithServerStreamInterceptor(handlerStreamMiddle(srvMidMap))

	// grpc server初始化
	grpcServer := conf.GrpcConfig.Build(
		grpc.ChainUnaryInterceptor(handlerUnaryMiddle(srvMidMap)),
		grpc.ChainStreamInterceptor(handlerStreamMiddle(srvMidMap))).Unwrap()

	mux := assert.Must1(larking.NewMux(
		larking.UnaryServerInterceptorOption(handlerUnaryMiddle(srvMidMap)),
		larking.StreamServerInterceptorOption(handlerStreamMiddle(srvMidMap)),
	))

	for _, h := range handlers {
		mux.RegisterService(h.ServiceDesc(), h)
		grpcServer.RegisterService(h.ServiceDesc(), h)
	}

	httpgrpc.HandleServices(
		func(pattern string, handler func(http.ResponseWriter, *http.Request)) {
			httpServer.Post(pattern, httputil.HTTPHandler(http.HandlerFunc(handler)))
		},
		assert.Must1(url.JoinPath(conf.BaseUrl, "api")),
		s.reg,
		handlerUnaryMiddle(srvMidMap),
		handlerStreamMiddle(srvMidMap),
	)

	wrappedGrpc := grpcweb.WrapServer(grpcServer,
		grpcweb.WithWebsockets(true),
		grpcweb.WithAllowNonRootResource(true),
		grpcweb.WithWebsocketOriginFunc(func(req *http.Request) bool { return true }),
		grpcweb.WithCorsForRegisteredEndpointsOnly(false),
		grpcweb.WithOriginFunc(func(origin string) bool { return true }))

	gwPrefix := assert.Must1(url.JoinPath(conf.BaseUrl, "gw"))
	httpServer.Group(gwPrefix+"/*", httputil.HTTPHandler(http.StripPrefix(gwPrefix, mux)))

	apiPrefix := assert.Must1(url.JoinPath(conf.BaseUrl, "api"))
	s.log.Info().Str("path", apiPrefix).Msg("service grpc gateway base path")
	httpServer.Group(apiPrefix+"/*", httputil.HTTPHandler(http.StripPrefix(apiPrefix, wsproxy.WebsocketProxy(grpcGateway))))

	grpcWebApiPrefix := assert.Must1(url.JoinPath(conf.BaseUrl, "grpc-web"))
	s.log.Info().Str("path", grpcWebApiPrefix).Msg("service grpc web base path")
	httpServer.Group(grpcWebApiPrefix+"/*", adaptor.HTTPHandler(h2c.NewHandler(http.StripPrefix(grpcWebApiPrefix,
		http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			if wrappedGrpc.IsAcceptableGrpcCorsRequest(request) {
				writer.WriteHeader(http.StatusNoContent)
				return
			}

			if wrappedGrpc.IsGrpcWebSocketRequest(request) {
				wrappedGrpc.HandleGrpcWebsocketRequest(writer, request)
				return
			}

			if wrappedGrpc.IsGrpcWebRequest(request) {
				wrappedGrpc.HandleGrpcWebRequest(writer, request)
				return
			}

			if grpcutil.IsGRPCRequest(request) {
				grpcServer.ServeHTTP(writer, request)
			}
		})), new(http2.Server))))

	s.httpServer = httpServer
	s.grpcServer = grpcServer

	vars.RegisterValue(fmt.Sprintf("%s-grpc-server-config", version.Project()), &conf)
}

func (s *serviceImpl) start() {
	defer recovery.Exit()

	logutil.OkOrFailed(s.log, "running before service starts", func() error {
		defer recovery.Exit()
		for _, run := range s.lc.GetBeforeStarts() {
			s.log.Info().Msgf("running %s", stack.CallerWithFunc(run.Handler))
			run.Handler()
		}
		return nil
	})

	logutil.OkOrFailed(s.log, "init handler before service starts", func() error {
		defer recovery.Exit()
		for _, ii := range s.initList {
			s.log.Info().Msgf("init handler %s", stack.CallerWithFunc(ii))
			ii()
		}
		return nil
	})

	s.log.Info().
		Int("grpc-port", *s.conf.GrpcPort).
		Int("http-port", *s.conf.HttpPort).
		Msg("create network listener")
	grpcLn := assert.Must1(net.Listen("tcp", fmt.Sprintf(":%d", *s.conf.GrpcPort)))
	httpLn := assert.Must1(net.Listen("tcp", fmt.Sprintf(":%d", *s.conf.HttpPort)))

	logutil.OkOrFailed(s.log, "service starts", func() error {
		// 启动grpc服务
		async.GoDelay(func() error {
			s.log.Info().Msg("[grpc] Server Starting")
			logutil.LogOrErr(s.log, "[grpc] Server Stop", func() error {
				defer recovery.Exit()
				if err := s.grpcServer.Serve(grpcLn); err != nil &&
					!errors.Is(err, http.ErrServerClosed) &&
					!errors.Is(err, net.ErrClosed) {
					return err
				}
				return nil
			})
			return nil
		})

		// 启动grpc网关
		async.GoDelay(func() error {
			s.log.Info().Msg("[http] Server Starting")
			logutil.LogOrErr(s.log, "[http] Server Stop", func() error {
				defer recovery.Exit()
				if err := s.httpServer.Listener(httpLn); err != nil &&
					!errors.Is(err, http.ErrServerClosed) &&
					!errors.Is(err, net.ErrClosed) {
					return err
				}
				return nil
			})
			return nil
		})
		return nil
	})

	logutil.OkOrFailed(s.log, "running after service starts", func() error {
		for _, run := range s.lc.GetAfterStarts() {
			logutil.LogOrErr(
				s.log,
				fmt.Sprintf("running %s", stack.CallerWithFunc(run.Handler)),
				func() error { run.Handler(); return nil },
			)
		}
		return nil
	})
}

func (s *serviceImpl) stop() {
	defer recovery.Exit()

	logutil.OkOrFailed(s.log, "running before service stops", func() error {
		for _, run := range s.lc.GetBeforeStops() {
			logutil.LogOrErr(
				s.log,
				fmt.Sprintf("running %s", stack.CallerWithFunc(run.Handler)),
				func() error { run.Handler(); return nil },
			)
		}
		return nil
	})

	logutil.LogOrErr(s.log, "[grpc] Server GracefulStop", func() error {
		s.grpcServer.GracefulStop()
		return nil
	})

	logutil.LogOrErr(s.log, "[http] Server Shutdown", func() error {
		return s.httpServer.ShutdownWithTimeout(consts.DefaultTimeout)
	})

	logutil.OkOrFailed(s.log, "running after service stops", func() error {
		for _, run := range s.lc.GetAfterStops() {
			logutil.LogOrErr(
				s.log,
				fmt.Sprintf("running %s", stack.CallerWithFunc(run.Handler)),
				func() error { run.Handler(); return nil },
			)
		}
		return nil
	})
}
