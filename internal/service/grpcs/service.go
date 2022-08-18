package grpcs

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"reflect"
	"strings"

	"github.com/fullstorydev/grpchan"
	"github.com/gofiber/adaptor/v2"
	fiber2 "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/pubgo/dix"
	"github.com/pubgo/funk/assert"
	"github.com/pubgo/funk/logx"
	"github.com/pubgo/funk/recovery"
	xtry "github.com/pubgo/funk/xtry"
	"github.com/pubgo/x/stack"
	"go.uber.org/zap"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/pubgo/lava/core/lifecycle"
	"github.com/pubgo/lava/core/runmode"
	"github.com/pubgo/lava/core/signal"
	fiber_builder2 "github.com/pubgo/lava/internal/pkg/fiber_builder"
	grpc_builder2 "github.com/pubgo/lava/internal/pkg/grpc_builder"
	"github.com/pubgo/lava/internal/pkg/syncx"
	"github.com/pubgo/lava/logging/logutil"
	"github.com/pubgo/lava/service"
)

func New() service.Service { return newService() }

func newService() *serviceImpl {
	return &serviceImpl{
		grpcSrv:    grpc_builder2.New(),
		httpSrv:    fiber_builder2.New(),
		handlers:   grpchan.HandlerMap{},
		httpMiddle: func(_ *fiber2.Ctx) error { return nil },
	}
}

var _ service.Service = (*serviceImpl)(nil)

type serviceImpl struct {
	getLifecycle lifecycle.GetLifecycle
	lc           lifecycle.Lifecycle
	app          *service.Web
	cfg          *Cfg
	log          *zap.Logger
	grpcSrv      grpc_builder2.Builder
	httpSrv      fiber_builder2.Builder
	handlers     grpchan.HandlerMap

	unaryInt   grpc.UnaryServerInterceptor
	streamInt  grpc.StreamServerInterceptor
	httpMiddle func(_ *fiber2.Ctx) error
	registers  []service.GatewayRegister
}

func (s *serviceImpl) RegisterServer(register interface{}, impl interface{}) {
	defer recovery.Exit()

	assert.If(impl == nil, "[impl] is nil")
	assert.If(register == nil, "[register] is nil")
	reflect.ValueOf(register).Call([]reflect.Value{reflect.ValueOf(s), reflect.ValueOf(impl)})
}

func (s *serviceImpl) RegisterGateway(register ...service.GatewayRegister) {
	s.registers = append(s.registers, register...)
}

func (s *serviceImpl) Run() {
	defer s.stop()
	s.start()
	signal.Wait()
}

func (s *serviceImpl) Start() { s.start() }
func (s *serviceImpl) Stop()  { s.stop() }

func (s *serviceImpl) init() {
	s.handlers.ForEach(func(_ *grpc.ServiceDesc, svc interface{}) { dix.Inject(svc) })
	var p = dix.Inject(new(struct {
		Middlewares  []service.Middleware
		GetLifecycle lifecycle.GetLifecycle
		Lifecycle    lifecycle.Lifecycle
		Log          *zap.Logger
		App          *service.Web
		Cfg          *Cfg
	}))

	s.getLifecycle = p.GetLifecycle
	s.lc = p.Lifecycle
	s.log = p.Log.Named("server")
	s.app = p.App
	s.cfg = p.Cfg

	var middlewares []service.Middleware
	for _, m := range p.Middlewares {
		middlewares = append(middlewares, m)
	}

	s.unaryInt = s.handlerUnaryMiddle(middlewares)
	s.streamInt = s.handlerStreamMiddle(middlewares)
	s.httpMiddle = s.handlerHttpMiddle(middlewares)

	s.grpcSrv.UnaryInterceptor(s.unaryInt)
	s.grpcSrv.StreamInterceptor(s.streamInt)
	// grpc serve初始化
	assert.Must(s.grpcSrv.Build(s.cfg.Grpc))

	// 初始化 handlers
	s.handlers.ForEach(func(desc *grpc.ServiceDesc, svr interface{}) {
		s.grpcSrv.Get().RegisterService(desc, svr)

		if h, ok := svr.(service.Close); ok {
			s.lc.AfterStop(h.Close)
		}

		if h, ok := svr.(service.WebHandler); ok {
			h.Router(s.app)
		}

		if h, ok := svr.(service.Init); ok {
			assert.Must(h.Init())
		}
	})

	s.app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowCredentials: true,
		//AllowHeaders: "Origin, Content-Type, Accept",
	}))

	var conn = assert.Must1(grpc.Dial(fmt.Sprintf("localhost:%d", runmode.GrpcPort), grpc.WithTransportCredentials(insecure.NewCredentials())))
	s.lc.BeforeStop(conn.Close)
	grpcMux := runtime.NewServeMux()
	for i := range s.registers {
		assert.Must(s.registers[i](context.Background(), grpcMux, conn))
	}

	wrappedGrpc := grpcweb.WrapServer(s.grpcSrv.Get(),
		grpcweb.WithAllowNonRootResource(true),
		grpcweb.WithOriginFunc(func(origin string) bool { return true }))

	var prefix = fmt.Sprintf("/%s/grpc", runmode.Project)
	s.app.All(prefix, adaptor.HTTPHandler(http.StripPrefix(prefix, h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			s.grpcSrv.Get().ServeHTTP(w, r)
			return
		}

		if wrappedGrpc.IsGrpcWebSocketRequest(r) {
			wrappedGrpc.HandleGrpcWebsocketRequest(w, r)
			return
		}

		if wrappedGrpc.IsGrpcWebRequest(r) {
			wrappedGrpc.HandleGrpcWebRequest(w, r)
			return
		}

		grpcMux.ServeHTTP(w, r)
	}), &http2.Server{}))))

	// 网关初始化
	assert.Must(s.httpSrv.Build(s.cfg.Api))
	s.httpSrv.Get().Use(s.httpMiddle)
	s.httpSrv.Get().Mount("/", s.app.App)

	if s.cfg.PrintRoute {
		for _, stacks := range s.httpSrv.Get().Stack() {
			for _, s := range stacks {
				logx.Info(
					"service route",
					"name", s.Name,
					"path", s.Path,
					"method", s.Method,
				)
			}
		}
	}
}

func (s *serviceImpl) RegisterService(desc *grpc.ServiceDesc, impl interface{}) {
	assert.Assert(desc == nil, "[desc] is nil")
	assert.Assert(impl == nil, "[impl] is nil")
	s.handlers.RegisterService(desc, impl)
}

func (s *serviceImpl) Provider(provider interface{}) {
	dix.Provider(provider)
}

func (s *serviceImpl) start() {
	s.init()

	logutil.OkOrPanic(s.log, "service before-start", func() error {
		for _, run := range s.getLifecycle.GetBeforeStarts() {
			s.log.Sugar().Infof("before-start running %s", stack.Func(run))
			assert.Must(xtry.Try(run.Handler), stack.Func(run))
		}
		return nil
	})

	grpcLn := assert.Must1(net.Listen("tcp", fmt.Sprintf(":%d", runmode.GrpcPort)))
	httpLn := assert.Must1(net.Listen("tcp", fmt.Sprintf(":%d", runmode.HttpPort)))

	logutil.OkOrPanic(s.log, "service start", func() error {
		// 启动grpc服务
		syncx.GoDelay(func() {
			s.log.Info("[grpc] Server Starting")
			logutil.LogOrErr(s.log, "[grpc] Server Stop", func() error {
				if err := s.grpcSrv.Get().Serve(grpcLn); err != nil &&
					!errors.Is(err, http.ErrServerClosed) &&
					!errors.Is(err, net.ErrClosed) {
					return err
				}
				return nil
			})
		})

		// 启动grpc网关
		syncx.GoDelay(func() {
			s.log.Info("[grpc-gw] Server Starting")
			logutil.LogOrErr(s.log, "[grpc-gw] Server Stop", func() error {
				if err := s.httpSrv.Get().Listener(httpLn); err != nil &&
					!errors.Is(err, http.ErrServerClosed) &&
					!errors.Is(err, net.ErrClosed) {
					return err
				}
				return nil
			})
		})
		return nil
	})

	logutil.OkOrPanic(s.log, "service after-start", func() error {
		for _, run := range s.getLifecycle.GetAfterStarts() {
			s.log.Sugar().Infof("after-start running %s", stack.Func(run))
			assert.Must(xtry.Try(run.Handler), stack.Func(run))
		}
		return nil
	})
}

func (s *serviceImpl) stop() {
	logutil.OkOrErr(s.log, "service before-stop", func() error {
		for _, run := range s.getLifecycle.GetBeforeStops() {
			s.log.Sugar().Infof("before-stop running %s", stack.Func(run))
			assert.Must(xtry.Try(run.Handler), stack.Func(run))
		}
		return nil
	})

	logutil.LogOrErr(s.log, "[grpc-gw] Shutdown", func() error {
		assert.Must(s.httpSrv.Get().Shutdown())
		return nil
	})

	logutil.LogOrErr(s.log, "[grpc] GracefulStop", func() error {
		s.grpcSrv.Get().GracefulStop()
		return nil
	})

	logutil.OkOrErr(s.log, "service after-stop", func() error {
		for _, run := range s.getLifecycle.GetAfterStops() {
			s.log.Sugar().Infof("after-stop running %s", stack.Func(run))
			assert.Must(xtry.Try(run.Handler), stack.Func(run))
		}
		return nil
	})
}
