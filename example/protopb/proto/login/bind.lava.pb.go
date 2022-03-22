// Code generated by protoc-gen-lava. DO NOT EDIT.
// versions:
// - protoc-gen-lava v0.1.0
// - protoc         v3.19.4
// source: proto/login/bind.proto

package login

import (
	context "context"
	runtime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	grpcc "github.com/pubgo/lava/clients/grpcc"
	service_type "github.com/pubgo/lava/service/service_type"
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

func InitBindTelephoneClient(srv string, opts ...func(cfg *grpcc.Cfg)) {

	opts = append(opts, grpcc.WithNewClientFunc(func(cc grpc.ClientConnInterface) interface{} {
		return NewBindTelephoneClient(cc)
	}))
	grpcc.InitClient(srv, append(opts, grpcc.WithClientType((*BindTelephoneClient)(nil)))...)
}

func RegisterBindTelephone(srv service_type.Service, impl BindTelephoneServer) {
	var desc service_type.Desc
	desc.Handler = impl
	desc.ServiceDesc = BindTelephone_ServiceDesc
	desc.GrpcClientFn = NewBindTelephoneClient

	desc.GrpcGatewayFn = func(mux *runtime.ServeMux) error {
		return RegisterBindTelephoneHandlerServer(context.Background(), mux, impl)
	}

	srv.RegisterService(desc)
}
