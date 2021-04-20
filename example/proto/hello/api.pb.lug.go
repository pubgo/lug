// Code generated by protoc-gen-lug. DO NOT EDIT.
// source: example/proto/hello/api.proto

package hello

import (
	"reflect"

	"github.com/pubgo/lug/client/grpcc"
	"github.com/pubgo/lug/xgen"
	"github.com/pubgo/xerror"
	"google.golang.org/grpc"
)

func GetTestApiClient(srv string, optFns ...func(service string) []grpc.DialOption) func() (TestApiClient, error) {
	client := grpcc.GetClient(srv, optFns...)
	return func() (TestApiClient, error) {
		c, err := client.Get()
		return &testApiClient{c}, xerror.WrapF(err, "srv: %s", srv)
	}
}

func GetTestApiV2Client(srv string, optFns ...func(service string) []grpc.DialOption) func() (TestApiV2Client, error) {
	client := grpcc.GetClient(srv, optFns...)
	return func() (TestApiV2Client, error) {
		c, err := client.Get()
		return &testApiV2Client{c}, xerror.WrapF(err, "srv: %s", srv)
	}
}

func init() {

	var mthList []xgen.GrpcRestHandler

	mthList = append(mthList, xgen.GrpcRestHandler{
		Service:       "hello.TestApi",
		Name:          "Version",
		Method:        "POST",
		Path:          "/hello/test_api/version",
		ClientStream:  "False" == "True",
		ServerStreams: "False" == "True",
	})

	mthList = append(mthList, xgen.GrpcRestHandler{
		Service:       "hello.TestApi",
		Name:          "VersionTest",
		Method:        "GET",
		Path:          "/v1/example/versiontest",
		ClientStream:  "False" == "True",
		ServerStreams: "False" == "True",
	})

	xgen.Add(reflect.ValueOf(RegisterTestApiServer), mthList)

	xgen.Add(reflect.ValueOf(RegisterTestApiHandlerFromEndpoint), nil)

	mthList = append(mthList, xgen.GrpcRestHandler{
		Service:       "hello.TestApiV2",
		Name:          "Version1",
		Method:        "POST",
		Path:          "/v2/example/version",
		ClientStream:  "False" == "True",
		ServerStreams: "False" == "True",
	})

	mthList = append(mthList, xgen.GrpcRestHandler{
		Service:       "hello.TestApiV2",
		Name:          "VersionTest1",
		Method:        "POST",
		Path:          "/v2/example/versiontest",
		ClientStream:  "False" == "True",
		ServerStreams: "False" == "True",
	})

	xgen.Add(reflect.ValueOf(RegisterTestApiV2Server), mthList)

	xgen.Add(reflect.ValueOf(RegisterTestApiV2HandlerFromEndpoint), nil)

}
