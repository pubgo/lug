// Code generated by protoc-gen-lava. DO NOT EDIT.
// versions:
// - protoc-gen-lava v0.1.0
// - protoc         v3.19.4
// source: proto/user/user.proto

package gid

import (
	context "context"
	runtime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	grpcc_builder "github.com/pubgo/lava/clients/grpcc/grpcc_builder"
	inject "github.com/pubgo/lava/inject"
	service "github.com/pubgo/lava/service"
	fx "go.uber.org/fx"
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

func InitUserClient(addr string, alias ...string) {

	var name = ""
	if len(alias) > 0 {
		name = alias[0]
	}
	conn := grpcc_builder.NewClient(addr)

	inject.Register(fx.Provide(fx.Annotated{
		Target: func() UserClient { return NewUserClient(conn) },
		Name:   name,
	}))
}

func RegisterUser(srv service.Service, impl UserServer) {
	srv.RegService(service.Desc{
		Handler:     impl,
		ServiceDesc: User_ServiceDesc,
	})

	srv.RegGateway(func(ctx context.Context, mux *runtime.ServeMux, cc grpc.ClientConnInterface) error {
		return RegisterUserHandlerClient(ctx, mux, NewUserClient(cc))
	})

}

func InitABitOfEverythingServiceClient(addr string, alias ...string) {

	var name = ""
	if len(alias) > 0 {
		name = alias[0]
	}
	conn := grpcc_builder.NewClient(addr)

	inject.Register(fx.Provide(fx.Annotated{
		Target: func() ABitOfEverythingServiceClient { return NewABitOfEverythingServiceClient(conn) },
		Name:   name,
	}))
}

func RegisterABitOfEverythingService(srv service.Service, impl ABitOfEverythingServiceServer) {
	srv.RegService(service.Desc{
		Handler:     impl,
		ServiceDesc: ABitOfEverythingService_ServiceDesc,
	})

	srv.RegGateway(func(ctx context.Context, mux *runtime.ServeMux, cc grpc.ClientConnInterface) error {
		return RegisterABitOfEverythingServiceHandlerClient(ctx, mux, NewABitOfEverythingServiceClient(cc))
	})

}

func InitCamelCaseServiceNameClient(addr string, alias ...string) {

	var name = ""
	if len(alias) > 0 {
		name = alias[0]
	}
	conn := grpcc_builder.NewClient(addr)

	inject.Register(fx.Provide(fx.Annotated{
		Target: func() CamelCaseServiceNameClient { return NewCamelCaseServiceNameClient(conn) },
		Name:   name,
	}))
}

func RegisterCamelCaseServiceName(srv service.Service, impl CamelCaseServiceNameServer) {
	srv.RegService(service.Desc{
		Handler:     impl,
		ServiceDesc: CamelCaseServiceName_ServiceDesc,
	})

	srv.RegGateway(func(ctx context.Context, mux *runtime.ServeMux, cc grpc.ClientConnInterface) error {
		return RegisterCamelCaseServiceNameHandlerClient(ctx, mux, NewCamelCaseServiceNameClient(cc))
	})

}

func InitAnotherServiceWithNoBindingsClient(addr string, alias ...string) {

	var name = ""
	if len(alias) > 0 {
		name = alias[0]
	}
	conn := grpcc_builder.NewClient(addr)

	inject.Register(fx.Provide(fx.Annotated{
		Target: func() AnotherServiceWithNoBindingsClient { return NewAnotherServiceWithNoBindingsClient(conn) },
		Name:   name,
	}))
}

func RegisterAnotherServiceWithNoBindings(srv service.Service, impl AnotherServiceWithNoBindingsServer) {
	srv.RegService(service.Desc{
		Handler:     impl,
		ServiceDesc: AnotherServiceWithNoBindings_ServiceDesc,
	})

	srv.RegGateway(func(ctx context.Context, mux *runtime.ServeMux, cc grpc.ClientConnInterface) error {
		return RegisterAnotherServiceWithNoBindingsHandlerClient(ctx, mux, NewAnotherServiceWithNoBindingsClient(cc))
	})

}
