// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: proto/services/metadata/metadata.proto

package metadata

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

const (
	Metadata_ListServices_FullMethodName   = "/lava.service.Metadata/ListServices"
	Metadata_GetServiceDesc_FullMethodName = "/lava.service.Metadata/GetServiceDesc"
)

// MetadataClient is the client API for Metadata service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MetadataClient interface {
	// ListServices list the full name of all services.
	ListServices(ctx context.Context, in *ListServicesRequest, opts ...grpc.CallOption) (*ListServicesReply, error)
	// GetServiceDesc get the full fileDescriptorSet of service.
	GetServiceDesc(ctx context.Context, in *GetServiceDescRequest, opts ...grpc.CallOption) (*GetServiceDescReply, error)
}

type metadataClient struct {
	cc grpc.ClientConnInterface
}

func NewMetadataClient(cc grpc.ClientConnInterface) MetadataClient {
	return &metadataClient{cc}
}

func (c *metadataClient) ListServices(ctx context.Context, in *ListServicesRequest, opts ...grpc.CallOption) (*ListServicesReply, error) {
	out := new(ListServicesReply)
	err := c.cc.Invoke(ctx, Metadata_ListServices_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metadataClient) GetServiceDesc(ctx context.Context, in *GetServiceDescRequest, opts ...grpc.CallOption) (*GetServiceDescReply, error) {
	out := new(GetServiceDescReply)
	err := c.cc.Invoke(ctx, Metadata_GetServiceDesc_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetadataServer is the server API for Metadata service.
// All implementations should embed UnimplementedMetadataServer
// for forward compatibility
type MetadataServer interface {
	// ListServices list the full name of all services.
	ListServices(context.Context, *ListServicesRequest) (*ListServicesReply, error)
	// GetServiceDesc get the full fileDescriptorSet of service.
	GetServiceDesc(context.Context, *GetServiceDescRequest) (*GetServiceDescReply, error)
}

// UnimplementedMetadataServer should be embedded to have forward compatible implementations.
type UnimplementedMetadataServer struct {
}

func (UnimplementedMetadataServer) ListServices(context.Context, *ListServicesRequest) (*ListServicesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListServices not implemented")
}
func (UnimplementedMetadataServer) GetServiceDesc(context.Context, *GetServiceDescRequest) (*GetServiceDescReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServiceDesc not implemented")
}

// UnsafeMetadataServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MetadataServer will
// result in compilation errors.
type UnsafeMetadataServer interface {
	mustEmbedUnimplementedMetadataServer()
}

func RegisterMetadataServer(s grpc.ServiceRegistrar, srv MetadataServer) {
	s.RegisterService(&Metadata_ServiceDesc, srv)
}

func _Metadata_ListServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListServicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServer).ListServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Metadata_ListServices_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServer).ListServices(ctx, req.(*ListServicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metadata_GetServiceDesc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServiceDescRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServer).GetServiceDesc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Metadata_GetServiceDesc_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServer).GetServiceDesc(ctx, req.(*GetServiceDescRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Metadata_ServiceDesc is the grpc.ServiceDesc for Metadata service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Metadata_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "lava.service.Metadata",
	HandlerType: (*MetadataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListServices",
			Handler:    _Metadata_ListServices_Handler,
		},
		{
			MethodName: "GetServiceDesc",
			Handler:    _Metadata_GetServiceDesc_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/services/metadata/metadata.proto",
}
