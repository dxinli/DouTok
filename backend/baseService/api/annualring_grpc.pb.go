// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.1
// source: annualring.proto

package api

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

// AnnualRingServiceClient is the client API for AnnualRingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AnnualRingServiceClient interface {
	// 获取全局唯一Guid号段
	GetGuidSegment(ctx context.Context, in *GetGuidSegmentRequest, opts ...grpc.CallOption) (*GetGuidSegmentResponse, error)
	// 上报号段使用情况
	UpdateGuidSegmentUsage(ctx context.Context, in *UpdateGuidSegmentUsageRequest, opts ...grpc.CallOption) (*UpdateGuidSegmentUsageResponse, error)
}

type annualRingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAnnualRingServiceClient(cc grpc.ClientConnInterface) AnnualRingServiceClient {
	return &annualRingServiceClient{cc}
}

func (c *annualRingServiceClient) GetGuidSegment(ctx context.Context, in *GetGuidSegmentRequest, opts ...grpc.CallOption) (*GetGuidSegmentResponse, error) {
	out := new(GetGuidSegmentResponse)
	err := c.cc.Invoke(ctx, "/api.AnnualRingService/GetGuidSegment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *annualRingServiceClient) UpdateGuidSegmentUsage(ctx context.Context, in *UpdateGuidSegmentUsageRequest, opts ...grpc.CallOption) (*UpdateGuidSegmentUsageResponse, error) {
	out := new(UpdateGuidSegmentUsageResponse)
	err := c.cc.Invoke(ctx, "/api.AnnualRingService/UpdateGuidSegmentUsage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AnnualRingServiceServer is the server API for AnnualRingService service.
// All implementations must embed UnimplementedAnnualRingServiceServer
// for forward compatibility
type AnnualRingServiceServer interface {
	// 获取全局唯一Guid号段
	GetGuidSegment(context.Context, *GetGuidSegmentRequest) (*GetGuidSegmentResponse, error)
	// 上报号段使用情况
	UpdateGuidSegmentUsage(context.Context, *UpdateGuidSegmentUsageRequest) (*UpdateGuidSegmentUsageResponse, error)
	mustEmbedUnimplementedAnnualRingServiceServer()
}

// UnimplementedAnnualRingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAnnualRingServiceServer struct {
}

func (UnimplementedAnnualRingServiceServer) GetGuidSegment(context.Context, *GetGuidSegmentRequest) (*GetGuidSegmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGuidSegment not implemented")
}
func (UnimplementedAnnualRingServiceServer) UpdateGuidSegmentUsage(context.Context, *UpdateGuidSegmentUsageRequest) (*UpdateGuidSegmentUsageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGuidSegmentUsage not implemented")
}
func (UnimplementedAnnualRingServiceServer) mustEmbedUnimplementedAnnualRingServiceServer() {}

// UnsafeAnnualRingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AnnualRingServiceServer will
// result in compilation errors.
type UnsafeAnnualRingServiceServer interface {
	mustEmbedUnimplementedAnnualRingServiceServer()
}

func RegisterAnnualRingServiceServer(s grpc.ServiceRegistrar, srv AnnualRingServiceServer) {
	s.RegisterService(&AnnualRingService_ServiceDesc, srv)
}

func _AnnualRingService_GetGuidSegment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGuidSegmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnnualRingServiceServer).GetGuidSegment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.AnnualRingService/GetGuidSegment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnnualRingServiceServer).GetGuidSegment(ctx, req.(*GetGuidSegmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnnualRingService_UpdateGuidSegmentUsage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGuidSegmentUsageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnnualRingServiceServer).UpdateGuidSegmentUsage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.AnnualRingService/UpdateGuidSegmentUsage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnnualRingServiceServer).UpdateGuidSegmentUsage(ctx, req.(*UpdateGuidSegmentUsageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AnnualRingService_ServiceDesc is the grpc.ServiceDesc for AnnualRingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AnnualRingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.AnnualRingService",
	HandlerType: (*AnnualRingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGuidSegment",
			Handler:    _AnnualRingService_GetGuidSegment_Handler,
		},
		{
			MethodName: "UpdateGuidSegmentUsage",
			Handler:    _AnnualRingService_UpdateGuidSegmentUsage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "annualring.proto",
}
