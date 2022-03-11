// Code generated by protoc-gen-lava. DO NOT EDIT.
// versions:
// - protoc-gen-lava v0.1.0
// - protoc         v3.19.4
// source: proto/login/code.proto

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

func InitCodeClient(srv string, opts ...func(cfg *grpcc.Cfg)) CodeClient {
	var cfg = grpcc.DefaultCfg(opts...)
	var cli = &codeClient{grpcc.NewClient(srv, cfg)}
	xerror.Exit(dix.ProviderNs(cfg.GetReg(), cli))
	return cli
}

func init() {
	var mthList []xgen.GrpcRestHandler
	mthList = append(mthList, xgen.GrpcRestHandler{
		Input:        &SendCodeRequest{},
		Output:       &SendCodeResponse{},
		Service:      "login.Code",
		Name:         "SendCode",
		Method:       "POST",
		Path:         "/user/code/send-code",
		DefaultUrl:   false,
		ClientStream: false,
		ServerStream: false,
	})

	mthList = append(mthList, xgen.GrpcRestHandler{
		Input:        &VerifyRequest{},
		Output:       &VerifyResponse{},
		Service:      "login.Code",
		Name:         "Verify",
		Method:       "POST",
		Path:         "/user/code/verify",
		DefaultUrl:   false,
		ClientStream: false,
		ServerStream: false,
	})

	mthList = append(mthList, xgen.GrpcRestHandler{
		Input:        &IsCheckImageCodeRequest{},
		Output:       &IsCheckImageCodeResponse{},
		Service:      "login.Code",
		Name:         "IsCheckImageCode",
		Method:       "POST",
		Path:         "/user/code/is-check-image-code",
		DefaultUrl:   false,
		ClientStream: false,
		ServerStream: false,
	})

	mthList = append(mthList, xgen.GrpcRestHandler{
		Input:        &VerifyImageCodeRequest{},
		Output:       &VerifyImageCodeResponse{},
		Service:      "login.Code",
		Name:         "VerifyImageCode",
		Method:       "POST",
		Path:         "/user/code/verify-image-code",
		DefaultUrl:   false,
		ClientStream: false,
		ServerStream: false,
	})

	mthList = append(mthList, xgen.GrpcRestHandler{
		Input:        &GetSendStatusRequest{},
		Output:       &GetSendStatusResponse{},
		Service:      "login.Code",
		Name:         "GetSendStatus",
		Method:       "POST",
		Path:         "/user/code/get-send-status",
		DefaultUrl:   false,
		ClientStream: false,
		ServerStream: false,
	})

	xgen.Add(RegisterCodeServer, mthList)
}

func RegisterCodeSrvServer(srv interface {
	Mux() *runtime.ServeMux
	Conn() grpc.ClientConnInterface
	RegisterService(desc *grpc.ServiceDesc, impl interface{})
}, impl CodeServer) {
	srv.RegisterService(&Code_ServiceDesc, impl)

	_ = RegisterCodeHandlerClient(context.Background(), srv.Mux(), NewCodeClient(srv.Conn()))

}
