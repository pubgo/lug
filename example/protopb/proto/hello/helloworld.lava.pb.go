// Code generated by protoc-gen-lava. DO NOT EDIT.
// versions:
// - protoc-gen-lava v0.1.0
// - protoc         v3.19.4
// source: proto/hello/helloworld.proto

package hello

import (
	context "context"
	runtime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	dix "github.com/pubgo/dix"
	grpcc "github.com/pubgo/lava/clients/grpcc"
	xgen "github.com/pubgo/lava/xgen"
	xerror "github.com/pubgo/xerror"
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

func InitGreeterClient(srv string, opts ...func(cfg *grpcc.Cfg)) GreeterClient {
	var cfg = grpcc.DefaultCfg(opts...)
	var cli = &greeterClient{grpcc.NewClient(srv, cfg)}
	xerror.Exit(dix.ProviderNs(cfg.GetReg(), cli))
	return cli
}

func init() {
	var mthList []xgen.GrpcRestHandler
	mthList = append(mthList, xgen.GrpcRestHandler{
		Input:        &HelloRequest{},
		Output:       &HelloReply{},
		Service:      "hello.Greeter",
		Name:         "SayHello",
		Method:       "GET",
		Path:         "/say/{name}",
		DefaultUrl:   false,
		ClientStream: false,
		ServerStream: false,
	})

	xgen.Add(RegisterGreeterServer, mthList)
}

func RegisterGreeterSrvServer(srv interface {
	Mux() *runtime.ServeMux
	Conn() grpc.ClientConnInterface
	RegisterService(desc *grpc.ServiceDesc, impl interface{})
}, impl GreeterServer) {
	srv.RegisterService(&Greeter_ServiceDesc, impl)

	_ = RegisterGreeterHandlerClient(context.Background(), srv.Mux(), NewGreeterClient(srv.Conn()))

}
