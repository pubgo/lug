// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: gid/id.proto

package gidpb

import (
	context "context"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Id_Generate_FullMethodName       = "/gid.Id/Generate"
	Id_TypeStream_FullMethodName     = "/gid.Id/TypeStream"
	Id_Types_FullMethodName          = "/gid.Id/Types"
	Id_Chat_FullMethodName           = "/gid.Id/Chat"
	Id_Chat1_FullMethodName          = "/gid.Id/Chat1"
	Id_UploadDownload_FullMethodName = "/gid.Id/UploadDownload"
)

// IdClient is the client API for Id service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IdClient interface {
	// Generate 生成ID
	Generate(ctx context.Context, in *GenerateRequest, opts ...grpc.CallOption) (*GenerateResponse, error)
	// 返回流
	TypeStream(ctx context.Context, in *TypesRequest, opts ...grpc.CallOption) (Id_TypeStreamClient, error)
	// Types id类型
	Types(ctx context.Context, in *TypesRequest, opts ...grpc.CallOption) (*TypesResponse, error)
	// 聊天
	Chat(ctx context.Context, opts ...grpc.CallOption) (Id_ChatClient, error)
	// ws: chat1
	Chat1(ctx context.Context, opts ...grpc.CallOption) (Id_Chat1Client, error)
	UploadDownload(ctx context.Context, in *UploadFileRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error)
}

type idClient struct {
	cc grpc.ClientConnInterface
}

func NewIdClient(cc grpc.ClientConnInterface) IdClient {
	return &idClient{cc}
}

func (c *idClient) Generate(ctx context.Context, in *GenerateRequest, opts ...grpc.CallOption) (*GenerateResponse, error) {
	out := new(GenerateResponse)
	err := c.cc.Invoke(ctx, Id_Generate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *idClient) TypeStream(ctx context.Context, in *TypesRequest, opts ...grpc.CallOption) (Id_TypeStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Id_ServiceDesc.Streams[0], Id_TypeStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &idTypeStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Id_TypeStreamClient interface {
	Recv() (*TypesResponse, error)
	grpc.ClientStream
}

type idTypeStreamClient struct {
	grpc.ClientStream
}

func (x *idTypeStreamClient) Recv() (*TypesResponse, error) {
	m := new(TypesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *idClient) Types(ctx context.Context, in *TypesRequest, opts ...grpc.CallOption) (*TypesResponse, error) {
	out := new(TypesResponse)
	err := c.cc.Invoke(ctx, Id_Types_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *idClient) Chat(ctx context.Context, opts ...grpc.CallOption) (Id_ChatClient, error) {
	stream, err := c.cc.NewStream(ctx, &Id_ServiceDesc.Streams[1], Id_Chat_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &idChatClient{stream}
	return x, nil
}

type Id_ChatClient interface {
	Send(*ChatMessage) error
	Recv() (*ChatMessage, error)
	grpc.ClientStream
}

type idChatClient struct {
	grpc.ClientStream
}

func (x *idChatClient) Send(m *ChatMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *idChatClient) Recv() (*ChatMessage, error) {
	m := new(ChatMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *idClient) Chat1(ctx context.Context, opts ...grpc.CallOption) (Id_Chat1Client, error) {
	stream, err := c.cc.NewStream(ctx, &Id_ServiceDesc.Streams[2], Id_Chat1_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &idChat1Client{stream}
	return x, nil
}

type Id_Chat1Client interface {
	Send(*ChatMessage) error
	Recv() (*ChatMessage, error)
	grpc.ClientStream
}

type idChat1Client struct {
	grpc.ClientStream
}

func (x *idChat1Client) Send(m *ChatMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *idChat1Client) Recv() (*ChatMessage, error) {
	m := new(ChatMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *idClient) UploadDownload(ctx context.Context, in *UploadFileRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error) {
	out := new(httpbody.HttpBody)
	err := c.cc.Invoke(ctx, Id_UploadDownload_FullMethodName, in, out, opts...)
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
	// 返回流
	TypeStream(*TypesRequest, Id_TypeStreamServer) error
	// Types id类型
	Types(context.Context, *TypesRequest) (*TypesResponse, error)
	// 聊天
	Chat(Id_ChatServer) error
	// ws: chat1
	Chat1(Id_Chat1Server) error
	UploadDownload(context.Context, *UploadFileRequest) (*httpbody.HttpBody, error)
}

// UnimplementedIdServer should be embedded to have forward compatible implementations.
type UnimplementedIdServer struct {
}

func (UnimplementedIdServer) Generate(context.Context, *GenerateRequest) (*GenerateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Generate not implemented")
}
func (UnimplementedIdServer) TypeStream(*TypesRequest, Id_TypeStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method TypeStream not implemented")
}
func (UnimplementedIdServer) Types(context.Context, *TypesRequest) (*TypesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Types not implemented")
}
func (UnimplementedIdServer) Chat(Id_ChatServer) error {
	return status.Errorf(codes.Unimplemented, "method Chat not implemented")
}
func (UnimplementedIdServer) Chat1(Id_Chat1Server) error {
	return status.Errorf(codes.Unimplemented, "method Chat1 not implemented")
}
func (UnimplementedIdServer) UploadDownload(context.Context, *UploadFileRequest) (*httpbody.HttpBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadDownload not implemented")
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
		FullMethod: Id_Generate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdServer).Generate(ctx, req.(*GenerateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Id_TypeStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TypesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(IdServer).TypeStream(m, &idTypeStreamServer{stream})
}

type Id_TypeStreamServer interface {
	Send(*TypesResponse) error
	grpc.ServerStream
}

type idTypeStreamServer struct {
	grpc.ServerStream
}

func (x *idTypeStreamServer) Send(m *TypesResponse) error {
	return x.ServerStream.SendMsg(m)
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
		FullMethod: Id_Types_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdServer).Types(ctx, req.(*TypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Id_Chat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(IdServer).Chat(&idChatServer{stream})
}

type Id_ChatServer interface {
	Send(*ChatMessage) error
	Recv() (*ChatMessage, error)
	grpc.ServerStream
}

type idChatServer struct {
	grpc.ServerStream
}

func (x *idChatServer) Send(m *ChatMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *idChatServer) Recv() (*ChatMessage, error) {
	m := new(ChatMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Id_Chat1_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(IdServer).Chat1(&idChat1Server{stream})
}

type Id_Chat1Server interface {
	Send(*ChatMessage) error
	Recv() (*ChatMessage, error)
	grpc.ServerStream
}

type idChat1Server struct {
	grpc.ServerStream
}

func (x *idChat1Server) Send(m *ChatMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *idChat1Server) Recv() (*ChatMessage, error) {
	m := new(ChatMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Id_UploadDownload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdServer).UploadDownload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Id_UploadDownload_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdServer).UploadDownload(ctx, req.(*UploadFileRequest))
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
		{
			MethodName: "UploadDownload",
			Handler:    _Id_UploadDownload_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TypeStream",
			Handler:       _Id_TypeStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Chat",
			Handler:       _Id_Chat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Chat1",
			Handler:       _Id_Chat1_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "gid/id.proto",
}
