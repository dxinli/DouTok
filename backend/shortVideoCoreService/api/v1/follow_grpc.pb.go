// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: v1/follow.proto

package v1

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

// FollowServiceClient is the client API for FollowService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FollowServiceClient interface {
	AddFollow(ctx context.Context, in *AddFollowRequest, opts ...grpc.CallOption) (*AddFollowResponse, error)
	RemoveFollow(ctx context.Context, in *RemoveFollowRequest, opts ...grpc.CallOption) (*RemoveFollowResponse, error)
	ListFollowing(ctx context.Context, in *ListFollowingRequest, opts ...grpc.CallOption) (*ListFollowingResponse, error)
	IsFollowing(ctx context.Context, in *IsFollowingRequest, opts ...grpc.CallOption) (*IsFollowingResponse, error)
	CountFollow(ctx context.Context, in *CountFollowRequest, opts ...grpc.CallOption) (*CountFollowResponse, error)
}

type followServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFollowServiceClient(cc grpc.ClientConnInterface) FollowServiceClient {
	return &followServiceClient{cc}
}

func (c *followServiceClient) AddFollow(ctx context.Context, in *AddFollowRequest, opts ...grpc.CallOption) (*AddFollowResponse, error) {
	out := new(AddFollowResponse)
	err := c.cc.Invoke(ctx, "/shortVideoCoreService.api.v1.FollowService/AddFollow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followServiceClient) RemoveFollow(ctx context.Context, in *RemoveFollowRequest, opts ...grpc.CallOption) (*RemoveFollowResponse, error) {
	out := new(RemoveFollowResponse)
	err := c.cc.Invoke(ctx, "/shortVideoCoreService.api.v1.FollowService/RemoveFollow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followServiceClient) ListFollowing(ctx context.Context, in *ListFollowingRequest, opts ...grpc.CallOption) (*ListFollowingResponse, error) {
	out := new(ListFollowingResponse)
	err := c.cc.Invoke(ctx, "/shortVideoCoreService.api.v1.FollowService/ListFollowing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followServiceClient) IsFollowing(ctx context.Context, in *IsFollowingRequest, opts ...grpc.CallOption) (*IsFollowingResponse, error) {
	out := new(IsFollowingResponse)
	err := c.cc.Invoke(ctx, "/shortVideoCoreService.api.v1.FollowService/IsFollowing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followServiceClient) CountFollow(ctx context.Context, in *CountFollowRequest, opts ...grpc.CallOption) (*CountFollowResponse, error) {
	out := new(CountFollowResponse)
	err := c.cc.Invoke(ctx, "/shortVideoCoreService.api.v1.FollowService/CountFollow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FollowServiceServer is the server API for FollowService service.
// All implementations must embed UnimplementedFollowServiceServer
// for forward compatibility
type FollowServiceServer interface {
	AddFollow(context.Context, *AddFollowRequest) (*AddFollowResponse, error)
	RemoveFollow(context.Context, *RemoveFollowRequest) (*RemoveFollowResponse, error)
	ListFollowing(context.Context, *ListFollowingRequest) (*ListFollowingResponse, error)
	IsFollowing(context.Context, *IsFollowingRequest) (*IsFollowingResponse, error)
	CountFollow(context.Context, *CountFollowRequest) (*CountFollowResponse, error)
	mustEmbedUnimplementedFollowServiceServer()
}

// UnimplementedFollowServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFollowServiceServer struct {
}

func (UnimplementedFollowServiceServer) AddFollow(context.Context, *AddFollowRequest) (*AddFollowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFollow not implemented")
}
func (UnimplementedFollowServiceServer) RemoveFollow(context.Context, *RemoveFollowRequest) (*RemoveFollowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveFollow not implemented")
}
func (UnimplementedFollowServiceServer) ListFollowing(context.Context, *ListFollowingRequest) (*ListFollowingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFollowing not implemented")
}
func (UnimplementedFollowServiceServer) IsFollowing(context.Context, *IsFollowingRequest) (*IsFollowingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsFollowing not implemented")
}
func (UnimplementedFollowServiceServer) CountFollow(context.Context, *CountFollowRequest) (*CountFollowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountFollow not implemented")
}
func (UnimplementedFollowServiceServer) mustEmbedUnimplementedFollowServiceServer() {}

// UnsafeFollowServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FollowServiceServer will
// result in compilation errors.
type UnsafeFollowServiceServer interface {
	mustEmbedUnimplementedFollowServiceServer()
}

func RegisterFollowServiceServer(s grpc.ServiceRegistrar, srv FollowServiceServer) {
	s.RegisterService(&FollowService_ServiceDesc, srv)
}

func _FollowService_AddFollow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFollowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowServiceServer).AddFollow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shortVideoCoreService.api.v1.FollowService/AddFollow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowServiceServer).AddFollow(ctx, req.(*AddFollowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FollowService_RemoveFollow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveFollowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowServiceServer).RemoveFollow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shortVideoCoreService.api.v1.FollowService/RemoveFollow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowServiceServer).RemoveFollow(ctx, req.(*RemoveFollowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FollowService_ListFollowing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFollowingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowServiceServer).ListFollowing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shortVideoCoreService.api.v1.FollowService/ListFollowing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowServiceServer).ListFollowing(ctx, req.(*ListFollowingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FollowService_IsFollowing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsFollowingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowServiceServer).IsFollowing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shortVideoCoreService.api.v1.FollowService/IsFollowing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowServiceServer).IsFollowing(ctx, req.(*IsFollowingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FollowService_CountFollow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountFollowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowServiceServer).CountFollow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shortVideoCoreService.api.v1.FollowService/CountFollow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowServiceServer).CountFollow(ctx, req.(*CountFollowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FollowService_ServiceDesc is the grpc.ServiceDesc for FollowService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FollowService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "shortVideoCoreService.api.v1.FollowService",
	HandlerType: (*FollowServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddFollow",
			Handler:    _FollowService_AddFollow_Handler,
		},
		{
			MethodName: "RemoveFollow",
			Handler:    _FollowService_RemoveFollow_Handler,
		},
		{
			MethodName: "ListFollowing",
			Handler:    _FollowService_ListFollowing_Handler,
		},
		{
			MethodName: "IsFollowing",
			Handler:    _FollowService_IsFollowing_Handler,
		},
		{
			MethodName: "CountFollow",
			Handler:    _FollowService_CountFollow_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/follow.proto",
}