// Code generated by protoc-gen-lava. DO NOT EDIT.
// versions:
// - protoc-gen-lava v0.1.0
// - protoc         v3.19.4
// source: proto/hello/helloworld.proto

package hello

import (
	context "context"
	runtime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	grpcc "github.com/pubgo/lava/clients/grpcc"
	"github.com/pubgo/lava/clients/grpcc/grpcc_config"
	service "github.com/pubgo/lava/service"
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

func InitGreeterClient(srv string, opts ...func(cfg *grpcc_config.Cfg)) {

	opts = append(opts, grpcc.WithNewClientFunc(func(cc grpc.ClientConnInterface) interface{} { return NewGreeterClient(cc) }))
	grpcc.InitClient(srv, append(opts, grpcc.WithClientType((*GreeterClient)(nil)))...)
}

func RegisterGreeter(srv service.Service, impl GreeterServer) {
	var desc service.Desc
	desc.Handler = impl
	desc.ServiceDesc = Greeter_ServiceDesc
	desc.GrpcClientFn = NewGreeterClient

	desc.GrpcGatewayFn = func(mux *runtime.ServeMux) error {
		return RegisterGreeterHandlerServer(context.Background(), mux, impl)
	}

	srv.RegisterService(desc)
}
