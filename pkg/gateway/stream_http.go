package gateway

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"github.com/pubgo/funk/errors"
	"github.com/pubgo/funk/generic"
	"github.com/pubgo/lava/pkg/gateway/internal/routertree"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type streamHTTP struct {
	method     *methodWrapper
	path       *routertree.MatchOperation
	handler    *fiber.Ctx
	ctx        context.Context
	header     metadata.MD
	params     url.Values
	sentHeader bool
}

var _ grpc.ServerStream = (*streamHTTP)(nil)

func (s *streamHTTP) SetHeader(md metadata.MD) error {
	if s.sentHeader {
		return errors.WrapStack(fmt.Errorf("already sent headers"))
	}
	s.header = metadata.Join(s.header, md)
	return nil
}

func (s *streamHTTP) SendHeader(md metadata.MD) error {
	if s.sentHeader {
		return errors.WrapCaller(fmt.Errorf("already sent headers"))
	}
	s.header = metadata.Join(s.header, md)
	s.sentHeader = true

	for k, v := range s.header {
		for i := range v {
			s.handler.Response().Header.Set(k, v[i])
		}
	}

	return nil
}

func (s *streamHTTP) SetTrailer(md metadata.MD) {
	s.header = metadata.Join(s.header, md)
}

func (s *streamHTTP) Context() context.Context {
	return grpc.NewContextWithServerTransportStream(
		s.ctx,
		&serverTransportStream{ServerStream: s, method: s.method.grpcFullMethod},
	)
}

func (s *streamHTTP) SendMsg(m interface{}) error {
	if generic.IsNil(m) {
		return errors.New("stream http send msg got nil")
	}

	reply, ok := m.(proto.Message)
	if !ok {
		return errors.New("stream http send proto msg got unknown type message")
	}

	if fRsp, ok := s.handler.Response().BodyWriter().(http.Flusher); ok {
		defer fRsp.Flush()
	}

	cur := reply.ProtoReflect()
	for _, fd := range getReqBodyDesc(s.path) {
		cur = cur.Mutable(fd).Message()
	}
	msg := cur.Interface()

	reqName := msg.ProtoReflect().Descriptor().FullName()
	handler := s.method.srv.opts.responseInterceptors[reqName]
	if handler != nil {
		return errors.Wrapf(handler(s.handler, msg), "failed to handler response data by %s", reqName)
	}

	b, err := protojson.Marshal(msg)
	if err != nil {
		return errors.Wrap(err, "failed to marshal response by proto-json")
	}

	_, err = s.handler.Write(b)
	return errors.WrapCaller(err)
}

func (s *streamHTTP) RecvMsg(m interface{}) error {
	if generic.IsNil(m) {
		return errors.New("stream http recv msg got nil")
	}

	args, ok := m.(proto.Message)
	if !ok {
		return errors.New("stream http recv proto msg got unknown type message")
	}

	var method = s.handler.Method()

	if method == http.MethodPut ||
		method == http.MethodPost ||
		method == http.MethodDelete ||
		method == http.MethodPatch {
		cur := args.ProtoReflect()
		for _, fd := range getRspBodyDesc(s.path) {
			cur = cur.Mutable(fd).Message()
		}
		msg := cur.Interface()

		reqName := msg.ProtoReflect().Descriptor().FullName()
		handler := s.method.srv.opts.requestInterceptors[reqName]
		if handler != nil {
			return errors.Wrapf(handler(s.handler, msg), "failed to handler request data by %s", reqName)
		}

		if method == http.MethodPut ||
			method == http.MethodPost ||
			method == http.MethodPatch {
			if len(s.handler.Body()) == 0 {
				return errors.WrapCaller(fmt.Errorf("request body is nil, operation=%s", reqName))
			}
		}

		if len(s.handler.Body()) > 0 {
			if err := protojson.Unmarshal(s.handler.Body(), msg); err != nil {
				return errors.Wrap(err, "failed to unmarshal body by proto-json")
			}
		}
	}

	if len(s.params) > 0 {
		if err := PopulateQueryParameters(args, s.params, utilities.NewDoubleArray(nil)); err != nil {
			return errors.Wrapf(err, "failed to set query params, params=%v", s.params)
		}
	}

	return nil
}
