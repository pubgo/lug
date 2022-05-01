// Code generated by protoc-gen-lava. DO NOT EDIT.
// versions:
// - protoc-gen-lava v0.1.0
// - protoc         v3.19.4
// source: proto/login/login.proto

package login

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

func InitLoginClient(addr string, alias ...string) {

	var name = ""
	if len(alias) > 0 {
		name = alias[0]
	}
	conn := grpcc_builder.NewClient(addr)

	inject.Register(fx.Provide(fx.Annotated{
		Target: func() LoginClient { return NewLoginClient(conn) },
		Name:   name,
	}))
}

func RegisterLogin(srv service.Service, impl LoginServer) {
	srv.RegService(service.Desc{
		Handler:     impl,
		ServiceDesc: Login_ServiceDesc,
	})

	srv.RegGateway(func(ctx context.Context, mux *runtime.ServeMux, cc grpc.ClientConnInterface) error {
		return RegisterLoginHandlerClient(ctx, mux, NewLoginClient(cc))
	})

}
