// Code generated by protoc-gen-lava. DO NOT EDIT.
// versions:
// - protoc-gen-lava v0.1.0
// - protoc         v3.19.4
// source: proto/hello/transport.proto

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

func InitTransportClient(srv string, opts ...func(cfg *grpcc.Cfg)) TransportClient {
	var cfg = grpcc.DefaultCfg(opts...)
	var cli = &transportClient{grpcc.NewClient(srv, cfg)}
	xerror.Exit(dix.ProviderNs(cfg.GetReg(), cli))
	return cli
}

func init() {
	var mthList []xgen.GrpcRestHandler
	mthList = append(mthList, xgen.GrpcRestHandler{
		Input:        &Message{},
		Output:       &Message{},
		Service:      "hello.Transport",
		Name:         "TestStream",
		Method:       "POST",
		Path:         "/hello/transport/test-stream",
		DefaultUrl:   true,
		ClientStream: true,
		ServerStream: true,
	})

	mthList = append(mthList, xgen.GrpcRestHandler{
		Input:        &Message{},
		Output:       &Message{},
		Service:      "hello.Transport",
		Name:         "TestStream1",
		Method:       "POST",
		Path:         "/hello/transport/test-stream1",
		DefaultUrl:   true,
		ClientStream: true,
		ServerStream: false,
	})

	mthList = append(mthList, xgen.GrpcRestHandler{
		Input:        &Message{},
		Output:       &Message{},
		Service:      "hello.Transport",
		Name:         "TestStream2",
		Method:       "POST",
		Path:         "/hello/transport/test-stream2",
		DefaultUrl:   true,
		ClientStream: false,
		ServerStream: true,
	})

	mthList = append(mthList, xgen.GrpcRestHandler{
		Input:        &Message{},
		Output:       &Message{},
		Service:      "hello.Transport",
		Name:         "TestStream3",
		Method:       "GET",
		Path:         "/v1/Transport/TestStream3",
		DefaultUrl:   false,
		ClientStream: false,
		ServerStream: false,
	})

	xgen.Add(RegisterTransportServer, mthList)
}

func RegisterTransportSrvServer(srv interface {
	Mux() *runtime.ServeMux
	Conn() grpc.ClientConnInterface
	RegisterService(desc *grpc.ServiceDesc, impl interface{})
}, impl TransportServer) {
	srv.RegisterService(&Transport_ServiceDesc, impl)

}
