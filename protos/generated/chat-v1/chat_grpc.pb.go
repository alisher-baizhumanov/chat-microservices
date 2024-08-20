// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.6
// source: chat.proto

package chat_v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ChatServiceV1Client is the client API for ChatServiceV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatServiceV1Client interface {
	CreateChat(ctx context.Context, in *CreateChatIn, opts ...grpc.CallOption) (*CreateChatOut, error)
	DeleteChat(ctx context.Context, in *DeleteChatIn, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SendMessage(ctx context.Context, in *SendMessageIn, opts ...grpc.CallOption) (*SendMessageOut, error)
}

type chatServiceV1Client struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceV1Client(cc grpc.ClientConnInterface) ChatServiceV1Client {
	return &chatServiceV1Client{cc}
}

func (c *chatServiceV1Client) CreateChat(ctx context.Context, in *CreateChatIn, opts ...grpc.CallOption) (*CreateChatOut, error) {
	out := new(CreateChatOut)
	err := c.cc.Invoke(ctx, "/chat_v1.ChatServiceV1/CreateChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceV1Client) DeleteChat(ctx context.Context, in *DeleteChatIn, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/chat_v1.ChatServiceV1/DeleteChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceV1Client) SendMessage(ctx context.Context, in *SendMessageIn, opts ...grpc.CallOption) (*SendMessageOut, error) {
	out := new(SendMessageOut)
	err := c.cc.Invoke(ctx, "/chat_v1.ChatServiceV1/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServiceV1Server is the server API for ChatServiceV1 service.
// All implementations must embed UnimplementedChatServiceV1Server
// for forward compatibility
type ChatServiceV1Server interface {
	CreateChat(context.Context, *CreateChatIn) (*CreateChatOut, error)
	DeleteChat(context.Context, *DeleteChatIn) (*emptypb.Empty, error)
	SendMessage(context.Context, *SendMessageIn) (*SendMessageOut, error)
	mustEmbedUnimplementedChatServiceV1Server()
}

// UnimplementedChatServiceV1Server must be embedded to have forward compatible implementations.
type UnimplementedChatServiceV1Server struct {
}

func (UnimplementedChatServiceV1Server) CreateChat(context.Context, *CreateChatIn) (*CreateChatOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChat not implemented")
}
func (UnimplementedChatServiceV1Server) DeleteChat(context.Context, *DeleteChatIn) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteChat not implemented")
}
func (UnimplementedChatServiceV1Server) SendMessage(context.Context, *SendMessageIn) (*SendMessageOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedChatServiceV1Server) mustEmbedUnimplementedChatServiceV1Server() {}

// UnsafeChatServiceV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServiceV1Server will
// result in compilation errors.
type UnsafeChatServiceV1Server interface {
	mustEmbedUnimplementedChatServiceV1Server()
}

func RegisterChatServiceV1Server(s grpc.ServiceRegistrar, srv ChatServiceV1Server) {
	s.RegisterService(&ChatServiceV1_ServiceDesc, srv)
}

func _ChatServiceV1_CreateChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateChatIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceV1Server).CreateChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat_v1.ChatServiceV1/CreateChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceV1Server).CreateChat(ctx, req.(*CreateChatIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatServiceV1_DeleteChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteChatIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceV1Server).DeleteChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat_v1.ChatServiceV1/DeleteChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceV1Server).DeleteChat(ctx, req.(*DeleteChatIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatServiceV1_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceV1Server).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat_v1.ChatServiceV1/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceV1Server).SendMessage(ctx, req.(*SendMessageIn))
	}
	return interceptor(ctx, in, info, handler)
}

// ChatServiceV1_ServiceDesc is the grpc.ServiceDesc for ChatServiceV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatServiceV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chat_v1.ChatServiceV1",
	HandlerType: (*ChatServiceV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateChat",
			Handler:    _ChatServiceV1_CreateChat_Handler,
		},
		{
			MethodName: "DeleteChat",
			Handler:    _ChatServiceV1_DeleteChat_Handler,
		},
		{
			MethodName: "SendMessage",
			Handler:    _ChatServiceV1_SendMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chat.proto",
}
