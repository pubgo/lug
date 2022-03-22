// Code generated by protoc-gen-lava. DO NOT EDIT.
// versions:
// - protoc-gen-lava v0.1.0
// - protoc         v3.19.4
// source: proto/gid/id.proto

package gid

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

func InitIdClient(srv string, opts ...func(cfg *grpcc.Cfg)) {

	opts = append(opts, grpcc.WithNewClientFunc(func(cc grpc.ClientConnInterface) interface{} {
		return NewIdClient(cc)
	}))
	grpcc.InitClient(srv, append(opts, grpcc.WithClientType((*IdClient)(nil)))...)
}

func RegisterId(srv service_type.Service, impl IdServer) {
	var desc service_type.Desc
	desc.Handler = impl
	desc.ServiceDesc = Id_ServiceDesc
	desc.GrpcClientFn = NewIdClient

	desc.GrpcGatewayFn = func(mux *runtime.ServeMux) error {
		return RegisterIdHandlerServer(context.Background(), mux, impl)
	}

	srv.RegisterService(desc)
}

func InitABitOfEverythingServiceClient(srv string, opts ...func(cfg *grpcc.Cfg)) {

	opts = append(opts, grpcc.WithNewClientFunc(func(cc grpc.ClientConnInterface) interface{} {
		return NewABitOfEverythingServiceClient(cc)
	}))
	grpcc.InitClient(srv, append(opts, grpcc.WithClientType((*ABitOfEverythingServiceClient)(nil)))...)
}

func RegisterABitOfEverythingService(srv service_type.Service, impl ABitOfEverythingServiceServer) {
	var desc service_type.Desc
	desc.Handler = impl
	desc.ServiceDesc = ABitOfEverythingService_ServiceDesc
	desc.GrpcClientFn = NewABitOfEverythingServiceClient

	desc.GrpcGatewayFn = func(mux *runtime.ServeMux) error {
		return RegisterABitOfEverythingServiceHandlerServer(context.Background(), mux, impl)
	}

	srv.RegisterService(desc)
}

func InitCamelCaseServiceNameClient(srv string, opts ...func(cfg *grpcc.Cfg)) {

	opts = append(opts, grpcc.WithNewClientFunc(func(cc grpc.ClientConnInterface) interface{} {
		return NewCamelCaseServiceNameClient(cc)
	}))
	grpcc.InitClient(srv, append(opts, grpcc.WithClientType((*CamelCaseServiceNameClient)(nil)))...)
}

func RegisterCamelCaseServiceName(srv service_type.Service, impl CamelCaseServiceNameServer) {
	var desc service_type.Desc
	desc.Handler = impl
	desc.ServiceDesc = CamelCaseServiceName_ServiceDesc
	desc.GrpcClientFn = NewCamelCaseServiceNameClient

	desc.GrpcGatewayFn = func(mux *runtime.ServeMux) error {
		return RegisterCamelCaseServiceNameHandlerServer(context.Background(), mux, impl)
	}

	srv.RegisterService(desc)
}

func InitAnotherServiceWithNoBindingsClient(srv string, opts ...func(cfg *grpcc.Cfg)) {

	opts = append(opts, grpcc.WithNewClientFunc(func(cc grpc.ClientConnInterface) interface{} {
		return NewAnotherServiceWithNoBindingsClient(cc)
	}))
	grpcc.InitClient(srv, append(opts, grpcc.WithClientType((*AnotherServiceWithNoBindingsClient)(nil)))...)
}

func RegisterAnotherServiceWithNoBindings(srv service_type.Service, impl AnotherServiceWithNoBindingsServer) {
	var desc service_type.Desc
	desc.Handler = impl
	desc.ServiceDesc = AnotherServiceWithNoBindings_ServiceDesc
	desc.GrpcClientFn = NewAnotherServiceWithNoBindingsClient

	desc.GrpcGatewayFn = func(mux *runtime.ServeMux) error {
		return RegisterAnotherServiceWithNoBindingsHandlerServer(context.Background(), mux, impl)
	}

	srv.RegisterService(desc)
}
