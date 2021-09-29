// Code generated by protoc-gen-lug. DO NOT EDIT.
// versions:
// - protoc-gen-lug v0.1.0
// - protoc         v3.17.3
// source: proto/gid/echo_service.proto

package gid

import (
	grpcc "github.com/pubgo/lug/plugins/grpcc"
	xgen "github.com/pubgo/lug/xgen"
	xerror "github.com/pubgo/xerror"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

func GetEchoServiceClient(srv string, opts ...func(cfg *grpcc.Cfg)) func(func(cli EchoServiceClient)) error {
	client := grpcc.GetClient(srv, opts...)
	return func(fn func(cli EchoServiceClient)) (err error) {
		defer xerror.RespErr(&err)

		c, err := client.Get()
		if err != nil {
			return xerror.WrapF(err, "srv: %s", srv)
		}

		fn(&echoServiceClient{c})
		return
	}
}
func init() {
	var mthList []xgen.GrpcRestHandler
	mthList = append(mthList, xgen.GrpcRestHandler{
		Input:        &SimpleMessage{},
		Output:       &SimpleMessage{},
		Service:      "gid.EchoService",
		Name:         "Echo",
		Method:       "POST",
		Path:         "/v1/example/echo/{id}",
		DefaultUrl:   false,
		ClientStream: false,
		ServerStream: false,
	})
	mthList = append(mthList, xgen.GrpcRestHandler{
		Input:        &SimpleMessage{},
		Output:       &SimpleMessage{},
		Service:      "gid.EchoService",
		Name:         "EchoBody",
		Method:       "POST",
		Path:         "/v1/example/echo_body",
		DefaultUrl:   false,
		ClientStream: false,
		ServerStream: false,
	})
	mthList = append(mthList, xgen.GrpcRestHandler{
		Input:        &SimpleMessage{},
		Output:       &SimpleMessage{},
		Service:      "gid.EchoService",
		Name:         "EchoDelete",
		Method:       "DELETE",
		Path:         "/v1/example/echo_delete",
		DefaultUrl:   false,
		ClientStream: false,
		ServerStream: false,
	})
	mthList = append(mthList, xgen.GrpcRestHandler{
		Input:        &DynamicMessageUpdate{},
		Output:       &DynamicMessageUpdate{},
		Service:      "gid.EchoService",
		Name:         "EchoPatch",
		Method:       "PATCH",
		Path:         "/v1/example/echo_patch",
		DefaultUrl:   false,
		ClientStream: false,
		ServerStream: false,
	})
	mthList = append(mthList, xgen.GrpcRestHandler{
		Input:        &SimpleMessage{},
		Output:       &SimpleMessage{},
		Service:      "gid.EchoService",
		Name:         "EchoUnauthorized",
		Method:       "GET",
		Path:         "/v1/example/echo_unauthorized",
		DefaultUrl:   false,
		ClientStream: false,
		ServerStream: false,
	})
	xgen.Add(reflect.ValueOf(RegisterEchoServiceServer), mthList)
	xgen.Add(reflect.ValueOf(RegisterEchoServiceHandlerServer), nil)
}
