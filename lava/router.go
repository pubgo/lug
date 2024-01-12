package lava

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/pubgo/opendoc/opendoc"
	"google.golang.org/grpc"
)

type EmptyRouter struct{}

type GrpcRouter interface {
	Middlewares() []Middleware
	ServiceDesc() *grpc.ServiceDesc
}

type GrpcGatewayRouter interface {
	GrpcRouter
	RegisterGateway(ctx context.Context, mux *runtime.ServeMux, conn grpc.ClientConnInterface) error
}

type HttpRouter interface {
	Middlewares() []Middleware
	Router(router *Router)
	//Annotation() []Annotation
}

type Router struct {
	R   fiber.Router
	Doc *opendoc.Service
}
