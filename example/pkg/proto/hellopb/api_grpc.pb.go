// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package hellopb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	structpb "google.golang.org/protobuf/types/known/structpb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TestApiClient is the client API for TestApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TestApiClient interface {
	// Version rpc
	Version(ctx context.Context, in *TestReq, opts ...grpc.CallOption) (*TestApiOutput, error)
	Version1(ctx context.Context, in *structpb.Value, opts ...grpc.CallOption) (*TestApiOutput1, error)
	// VersionTest rpc
	VersionTest(ctx context.Context, in *TestReq, opts ...grpc.CallOption) (*TestApiOutput, error)
	// VersionTest rpc custom
	VersionTestCustom(ctx context.Context, in *TestReq, opts ...grpc.CallOption) (*TestApiOutput, error)
}

type testApiClient struct {
	cc grpc.ClientConnInterface
}

func NewTestApiClient(cc grpc.ClientConnInterface) TestApiClient {
	return &testApiClient{cc}
}

func (c *testApiClient) Version(ctx context.Context, in *TestReq, opts ...grpc.CallOption) (*TestApiOutput, error) {
	out := new(TestApiOutput)
	err := c.cc.Invoke(ctx, "/hello.TestApi/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testApiClient) Version1(ctx context.Context, in *structpb.Value, opts ...grpc.CallOption) (*TestApiOutput1, error) {
	out := new(TestApiOutput1)
	err := c.cc.Invoke(ctx, "/hello.TestApi/Version1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testApiClient) VersionTest(ctx context.Context, in *TestReq, opts ...grpc.CallOption) (*TestApiOutput, error) {
	out := new(TestApiOutput)
	err := c.cc.Invoke(ctx, "/hello.TestApi/VersionTest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testApiClient) VersionTestCustom(ctx context.Context, in *TestReq, opts ...grpc.CallOption) (*TestApiOutput, error) {
	out := new(TestApiOutput)
	err := c.cc.Invoke(ctx, "/hello.TestApi/VersionTestCustom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestApiServer is the server API for TestApi service.
// All implementations should embed UnimplementedTestApiServer
// for forward compatibility
type TestApiServer interface {
	// Version rpc
	Version(context.Context, *TestReq) (*TestApiOutput, error)
	Version1(context.Context, *structpb.Value) (*TestApiOutput1, error)
	// VersionTest rpc
	VersionTest(context.Context, *TestReq) (*TestApiOutput, error)
	// VersionTest rpc custom
	VersionTestCustom(context.Context, *TestReq) (*TestApiOutput, error)
}

// UnimplementedTestApiServer should be embedded to have forward compatible implementations.
type UnimplementedTestApiServer struct {
}

func (UnimplementedTestApiServer) Version(context.Context, *TestReq) (*TestApiOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedTestApiServer) Version1(context.Context, *structpb.Value) (*TestApiOutput1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version1 not implemented")
}
func (UnimplementedTestApiServer) VersionTest(context.Context, *TestReq) (*TestApiOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VersionTest not implemented")
}
func (UnimplementedTestApiServer) VersionTestCustom(context.Context, *TestReq) (*TestApiOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VersionTestCustom not implemented")
}

// UnsafeTestApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TestApiServer will
// result in compilation errors.
type UnsafeTestApiServer interface {
	mustEmbedUnimplementedTestApiServer()
}

func RegisterTestApiServer(s grpc.ServiceRegistrar, srv TestApiServer) {
	s.RegisterService(&TestApi_ServiceDesc, srv)
}

func _TestApi_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestApiServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.TestApi/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestApiServer).Version(ctx, req.(*TestReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestApi_Version1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(structpb.Value)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestApiServer).Version1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.TestApi/Version1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestApiServer).Version1(ctx, req.(*structpb.Value))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestApi_VersionTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestApiServer).VersionTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.TestApi/VersionTest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestApiServer).VersionTest(ctx, req.(*TestReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestApi_VersionTestCustom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestApiServer).VersionTestCustom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.TestApi/VersionTestCustom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestApiServer).VersionTestCustom(ctx, req.(*TestReq))
	}
	return interceptor(ctx, in, info, handler)
}

// TestApi_ServiceDesc is the grpc.ServiceDesc for TestApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TestApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hello.TestApi",
	HandlerType: (*TestApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _TestApi_Version_Handler,
		},
		{
			MethodName: "Version1",
			Handler:    _TestApi_Version1_Handler,
		},
		{
			MethodName: "VersionTest",
			Handler:    _TestApi_VersionTest_Handler,
		},
		{
			MethodName: "VersionTestCustom",
			Handler:    _TestApi_VersionTestCustom_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello/api.proto",
}

// TestApiV2Client is the client API for TestApiV2 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TestApiV2Client interface {
	Version1(ctx context.Context, in *TestReq, opts ...grpc.CallOption) (*TestApiOutput, error)
	VersionTest1(ctx context.Context, in *TestReq, opts ...grpc.CallOption) (*TestApiOutput, error)
}

type testApiV2Client struct {
	cc grpc.ClientConnInterface
}

func NewTestApiV2Client(cc grpc.ClientConnInterface) TestApiV2Client {
	return &testApiV2Client{cc}
}

func (c *testApiV2Client) Version1(ctx context.Context, in *TestReq, opts ...grpc.CallOption) (*TestApiOutput, error) {
	out := new(TestApiOutput)
	err := c.cc.Invoke(ctx, "/hello.TestApiV2/Version1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testApiV2Client) VersionTest1(ctx context.Context, in *TestReq, opts ...grpc.CallOption) (*TestApiOutput, error) {
	out := new(TestApiOutput)
	err := c.cc.Invoke(ctx, "/hello.TestApiV2/VersionTest1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestApiV2Server is the server API for TestApiV2 service.
// All implementations should embed UnimplementedTestApiV2Server
// for forward compatibility
type TestApiV2Server interface {
	Version1(context.Context, *TestReq) (*TestApiOutput, error)
	VersionTest1(context.Context, *TestReq) (*TestApiOutput, error)
}

// UnimplementedTestApiV2Server should be embedded to have forward compatible implementations.
type UnimplementedTestApiV2Server struct {
}

func (UnimplementedTestApiV2Server) Version1(context.Context, *TestReq) (*TestApiOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version1 not implemented")
}
func (UnimplementedTestApiV2Server) VersionTest1(context.Context, *TestReq) (*TestApiOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VersionTest1 not implemented")
}

// UnsafeTestApiV2Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TestApiV2Server will
// result in compilation errors.
type UnsafeTestApiV2Server interface {
	mustEmbedUnimplementedTestApiV2Server()
}

func RegisterTestApiV2Server(s grpc.ServiceRegistrar, srv TestApiV2Server) {
	s.RegisterService(&TestApiV2_ServiceDesc, srv)
}

func _TestApiV2_Version1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestApiV2Server).Version1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.TestApiV2/Version1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestApiV2Server).Version1(ctx, req.(*TestReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestApiV2_VersionTest1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestApiV2Server).VersionTest1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.TestApiV2/VersionTest1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestApiV2Server).VersionTest1(ctx, req.(*TestReq))
	}
	return interceptor(ctx, in, info, handler)
}

// TestApiV2_ServiceDesc is the grpc.ServiceDesc for TestApiV2 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TestApiV2_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hello.TestApiV2",
	HandlerType: (*TestApiV2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version1",
			Handler:    _TestApiV2_Version1_Handler,
		},
		{
			MethodName: "VersionTest1",
			Handler:    _TestApiV2_VersionTest1_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello/api.proto",
}
