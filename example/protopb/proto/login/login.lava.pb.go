// Code generated by protoc-gen-lava. DO NOT EDIT.
// versions:
// - protoc-gen-lava v0.1.0
// - protoc         v3.19.4
// source: proto/login/login.proto

package login

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

func InitLoginClient(srv string, opts ...func(cfg *grpcc.Cfg)) LoginClient {
	var cfg = grpcc.DefaultCfg(opts...)
	var cli = &loginClient{grpcc.NewClient(srv, cfg)}
	xerror.Exit(dix.ProviderNs(cfg.GetReg(), cli))
	return cli
}

func init() {
	var mthList []xgen.GrpcRestHandler
	mthList = append(mthList, xgen.GrpcRestHandler{
		Input:        &LoginRequest{},
		Output:       &LoginResponse{},
		Service:      "login.Login",
		Name:         "Login",
		Method:       "POST",
		Path:         "/user/login/login",
		DefaultUrl:   false,
		ClientStream: false,
		ServerStream: false,
	})

	mthList = append(mthList, xgen.GrpcRestHandler{
		Input:        &AuthenticateRequest{},
		Output:       &AuthenticateResponse{},
		Service:      "login.Login",
		Name:         "Authenticate",
		Method:       "POST",
		Path:         "/user/login/authenticate",
		DefaultUrl:   false,
		ClientStream: false,
		ServerStream: false,
	})

	xgen.Add(RegisterLoginServer, mthList)
}

func RegisterLoginSrvServer(srv interface {
	Mux() *runtime.ServeMux
	Conn() grpc.ClientConnInterface
	RegisterService(desc *grpc.ServiceDesc, impl interface{})
}, impl LoginServer) {
	srv.RegisterService(&Login_ServiceDesc, impl)

	_ = RegisterLoginHandlerClient(context.Background(), srv.Mux(), NewLoginClient(srv.Conn()))

}
