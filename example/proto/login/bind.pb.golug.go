// Code generated by protoc-gen-golug. DO NOT EDIT.
// source: example/proto/login/bind.proto

package login

import (
	"reflect"

	"github.com/pubgo/golug/golug_data"
	"github.com/pubgo/golug/golug_entry"
)

func init() {
	var _mth []golug_entry.GrpcRestHandler
	for _, m := range ss.GetMethod() {
		_mth = append(_mth, golug_entry.GrpcRestHandler{
			Name:          m.GetName(),
			Method:        m.HttpMethod,
			Path:          m.HttpPath,
			ClientStream:  m.CS,
			ServerStreams: m.SS,
		})
	}
	golug_data.Add(reflect.ValueOf(RegisterBindTelephoneServer), _mth)
}
