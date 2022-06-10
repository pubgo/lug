// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package gidpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// IdClient is the client API for Id service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IdClient interface {
	// Generate 生成ID
	Generate(ctx context.Context, in *GenerateRequest, opts ...grpc.CallOption) (*GenerateResponse, error)
	// Types id类型
	Types(ctx context.Context, in *TypesRequest, opts ...grpc.CallOption) (*TypesResponse, error)
}

type idClient struct {
	cc grpc.ClientConnInterface
}

func NewIdClient(cc grpc.ClientConnInterface) IdClient {
	return &idClient{cc}
}

func (c *idClient) Generate(ctx context.Context, in *GenerateRequest, opts ...grpc.CallOption) (*GenerateResponse, error) {
	out := new(GenerateResponse)
	err := c.cc.Invoke(ctx, "/gid.Id/Generate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *idClient) Types(ctx context.Context, in *TypesRequest, opts ...grpc.CallOption) (*TypesResponse, error) {
	out := new(TypesResponse)
	err := c.cc.Invoke(ctx, "/gid.Id/Types", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IdServer is the server API for Id service.
// All implementations should embed UnimplementedIdServer
// for forward compatibility
type IdServer interface {
	// Generate 生成ID
	Generate(context.Context, *GenerateRequest) (*GenerateResponse, error)
	// Types id类型
	Types(context.Context, *TypesRequest) (*TypesResponse, error)
}

// UnimplementedIdServer should be embedded to have forward compatible implementations.
type UnimplementedIdServer struct {
}

func (UnimplementedIdServer) Generate(context.Context, *GenerateRequest) (*GenerateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Generate not implemented")
}
func (UnimplementedIdServer) Types(context.Context, *TypesRequest) (*TypesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Types not implemented")
}

// UnsafeIdServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IdServer will
// result in compilation errors.
type UnsafeIdServer interface {
	mustEmbedUnimplementedIdServer()
}

func RegisterIdServer(s grpc.ServiceRegistrar, srv IdServer) {
	s.RegisterService(&Id_ServiceDesc, srv)
}

func _Id_Generate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdServer).Generate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gid.Id/Generate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdServer).Generate(ctx, req.(*GenerateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Id_Types_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TypesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdServer).Types(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gid.Id/Types",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdServer).Types(ctx, req.(*TypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Id_ServiceDesc is the grpc.ServiceDesc for Id service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Id_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gid.Id",
	HandlerType: (*IdServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Generate",
			Handler:    _Id_Generate_Handler,
		},
		{
			MethodName: "Types",
			Handler:    _Id_Types_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "example/gid/id.proto",
}
