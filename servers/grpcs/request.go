package grpcs

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pubgo/lava/service"
	"google.golang.org/grpc"
)

var _ service.Request = (*rpcRequest)(nil)

type rpcRequest struct {
	handler       grpc.UnaryHandler
	handlerStream grpc.StreamHandler
	stream        grpc.ServerStream
	srv           interface{}
	service       string
	method        string
	url           string
	contentType   string
	header        *service.RequestHeader
	payload       interface{}
}

func (r *rpcRequest) Kind() string                   { return "grpc" }
func (r *rpcRequest) Client() bool                   { return false }
func (r *rpcRequest) Header() *service.RequestHeader { return r.header }
func (r *rpcRequest) Payload() interface{}           { return r.payload }
func (r *rpcRequest) ContentType() string            { return r.contentType }
func (r *rpcRequest) Service() string                { return r.service }
func (r *rpcRequest) Operation() string              { return r.method }
func (r *rpcRequest) Endpoint() string               { return r.url }
func (r *rpcRequest) Stream() bool                   { return r.stream != nil }

var _ service.Request = (*httpRequest)(nil)

type httpRequest struct {
	ctx *fiber.Ctx
}

func (r *httpRequest) Kind() string                   { return "http" }
func (r *httpRequest) Operation() string              { return r.ctx.Route().Path }
func (r *httpRequest) Client() bool                   { return false }
func (r *httpRequest) Header() *service.RequestHeader { return &r.ctx.Request().Header }
func (r *httpRequest) Payload() interface{}           { return r.ctx.Body() }

func (r *httpRequest) ContentType() string {
	return string(r.ctx.Request().Header.ContentType())
}

func (r *httpRequest) Service() string  { return r.ctx.OriginalURL() }
func (r *httpRequest) Endpoint() string { return r.ctx.OriginalURL() }
func (r *httpRequest) Stream() bool     { return false }
