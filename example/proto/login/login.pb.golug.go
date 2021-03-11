// Code generated by protoc-gen-golug. DO NOT EDIT.
// source: example/proto/login/login.proto

package login

import (
	"bytes"
	"reflect"

	"github.com/gofiber/fiber/v2"
	"github.com/pubgo/golug/client/grpclient"
	"github.com/pubgo/golug/gutils"
	"github.com/pubgo/golug/xgen"
	"github.com/pubgo/x/xutil"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var _ = gutils.Decode

func init() {
	var mthList []xgen.GrpcRestHandler
	mthList = append(mthList, xgen.GrpcRestHandler{
		Service:       "login.Login",
		Name:          "Login",
		Method:        "POST",
		Path:          "/user/login/login",
		ClientStream:  "False" == "True",
		ServerStreams: "False" == "True",
	})

	mthList = append(mthList, xgen.GrpcRestHandler{
		Service:       "login.Login",
		Name:          "Authenticate",
		Method:        "POST",
		Path:          "/user/login/authenticate",
		ClientStream:  "False" == "True",
		ServerStreams: "False" == "True",
	})

	xgen.Add(reflect.ValueOf(RegisterLoginServer), mthList)
	xgen.Add(reflect.ValueOf(RegisterLoginHandlerFromEndpoint), nil)
}

func GetLoginClient(srv string, opts ...grpc.DialOption) func() (LoginClient, error) {
	client := grpclient.Client(srv, opts...)
	return func() (LoginClient, error) {
		c, err := client.Get()
		return &loginClient{c}, err
	}
}
