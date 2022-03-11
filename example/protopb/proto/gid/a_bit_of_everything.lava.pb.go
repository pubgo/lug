// Code generated by protoc-gen-lava. DO NOT EDIT.
// versions:
// - protoc-gen-lava v0.1.0
// - protoc         v3.19.4
// source: proto/gid/a_bit_of_everything.proto

package gid

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

func InitLoginServiceClient(srv string, opts ...func(cfg *grpcc.Cfg)) LoginServiceClient {
	var cfg = grpcc.DefaultCfg(opts...)
	var cli = &loginServiceClient{grpcc.NewClient(srv, cfg)}
	xerror.Exit(dix.ProviderNs(cfg.GetReg(), cli))
	return cli
}

func init() {
	var mthList []xgen.GrpcRestHandler
	mthList = append(mthList, xgen.GrpcRestHandler{
		Input:        &LoginRequest{},
		Output:       &LoginReply{},
		Service:      "gid.LoginService",
		Name:         "Login",
		Method:       "POST",
		Path:         "/v1/example/login",
		DefaultUrl:   false,
		ClientStream: false,
		ServerStream: false,
	})

	mthList = append(mthList, xgen.GrpcRestHandler{
		Input:        &LogoutRequest{},
		Output:       &LogoutReply{},
		Service:      "gid.LoginService",
		Name:         "Logout",
		Method:       "POST",
		Path:         "/v1/example/logout",
		DefaultUrl:   false,
		ClientStream: false,
		ServerStream: false,
	})

	xgen.Add(RegisterLoginServiceServer, mthList)
}

func RegisterLoginServiceSrvServer(srv interface {
	Mux() *runtime.ServeMux
	Conn() grpc.ClientConnInterface
	RegisterService(desc *grpc.ServiceDesc, impl interface{})
}, impl LoginServiceServer) {
	srv.RegisterService(&LoginService_ServiceDesc, impl)

	_ = RegisterLoginServiceHandlerClient(context.Background(), srv.Mux(), NewLoginServiceClient(srv.Conn()))

}
