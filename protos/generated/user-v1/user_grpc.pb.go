// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.2
// source: user.proto

package user_v1

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

// UserServiceV1Client is the client API for UserServiceV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceV1Client interface {
	Create(ctx context.Context, in *CreateIn, opts ...grpc.CallOption) (*CreateOut, error)
	Get(ctx context.Context, in *GetIn, opts ...grpc.CallOption) (*GetOut, error)
	Update(ctx context.Context, in *UpdateIn, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Delete(ctx context.Context, in *DeleteIn, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type userServiceV1Client struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceV1Client(cc grpc.ClientConnInterface) UserServiceV1Client {
	return &userServiceV1Client{cc}
}

func (c *userServiceV1Client) Create(ctx context.Context, in *CreateIn, opts ...grpc.CallOption) (*CreateOut, error) {
	out := new(CreateOut)
	err := c.cc.Invoke(ctx, "/user_v1.UserServiceV1/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceV1Client) Get(ctx context.Context, in *GetIn, opts ...grpc.CallOption) (*GetOut, error) {
	out := new(GetOut)
	err := c.cc.Invoke(ctx, "/user_v1.UserServiceV1/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceV1Client) Update(ctx context.Context, in *UpdateIn, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/user_v1.UserServiceV1/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceV1Client) Delete(ctx context.Context, in *DeleteIn, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/user_v1.UserServiceV1/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceV1Server is the server API for UserServiceV1 service.
// All implementations must embed UnimplementedUserServiceV1Server
// for forward compatibility
type UserServiceV1Server interface {
	Create(context.Context, *CreateIn) (*CreateOut, error)
	Get(context.Context, *GetIn) (*GetOut, error)
	Update(context.Context, *UpdateIn) (*emptypb.Empty, error)
	Delete(context.Context, *DeleteIn) (*emptypb.Empty, error)
	mustEmbedUnimplementedUserServiceV1Server()
}

// UnimplementedUserServiceV1Server must be embedded to have forward compatible implementations.
type UnimplementedUserServiceV1Server struct {
}

func (UnimplementedUserServiceV1Server) Create(context.Context, *CreateIn) (*CreateOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedUserServiceV1Server) Get(context.Context, *GetIn) (*GetOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedUserServiceV1Server) Update(context.Context, *UpdateIn) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedUserServiceV1Server) Delete(context.Context, *DeleteIn) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedUserServiceV1Server) mustEmbedUnimplementedUserServiceV1Server() {}

// UnsafeUserServiceV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceV1Server will
// result in compilation errors.
type UnsafeUserServiceV1Server interface {
	mustEmbedUnimplementedUserServiceV1Server()
}

func RegisterUserServiceV1Server(s grpc.ServiceRegistrar, srv UserServiceV1Server) {
	s.RegisterService(&UserServiceV1_ServiceDesc, srv)
}

func _UserServiceV1_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceV1Server).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_v1.UserServiceV1/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceV1Server).Create(ctx, req.(*CreateIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServiceV1_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceV1Server).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_v1.UserServiceV1/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceV1Server).Get(ctx, req.(*GetIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServiceV1_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceV1Server).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_v1.UserServiceV1/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceV1Server).Update(ctx, req.(*UpdateIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServiceV1_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceV1Server).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_v1.UserServiceV1/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceV1Server).Delete(ctx, req.(*DeleteIn))
	}
	return interceptor(ctx, in, info, handler)
}

// UserServiceV1_ServiceDesc is the grpc.ServiceDesc for UserServiceV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserServiceV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user_v1.UserServiceV1",
	HandlerType: (*UserServiceV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _UserServiceV1_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _UserServiceV1_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _UserServiceV1_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _UserServiceV1_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
