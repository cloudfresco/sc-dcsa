// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: internal/proto/jit/v1/timestamp.proto

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

const (
	TimestampService_CreateTimestamp_FullMethodName = "/internal.proto.jit.v1.TimestampService/CreateTimestamp"
	TimestampService_GetTimestamps_FullMethodName   = "/internal.proto.jit.v1.TimestampService/GetTimestamps"
)

// TimestampServiceClient is the client API for TimestampService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TimestampServiceClient interface {
	CreateTimestamp(ctx context.Context, in *CreateTimestampRequest, opts ...grpc.CallOption) (*CreateTimestampResponse, error)
	GetTimestamps(ctx context.Context, in *GetTimestampsRequest, opts ...grpc.CallOption) (*GetTimestampsResponse, error)
}

type timestampServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTimestampServiceClient(cc grpc.ClientConnInterface) TimestampServiceClient {
	return &timestampServiceClient{cc}
}

func (c *timestampServiceClient) CreateTimestamp(ctx context.Context, in *CreateTimestampRequest, opts ...grpc.CallOption) (*CreateTimestampResponse, error) {
	out := new(CreateTimestampResponse)
	err := c.cc.Invoke(ctx, TimestampService_CreateTimestamp_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timestampServiceClient) GetTimestamps(ctx context.Context, in *GetTimestampsRequest, opts ...grpc.CallOption) (*GetTimestampsResponse, error) {
	out := new(GetTimestampsResponse)
	err := c.cc.Invoke(ctx, TimestampService_GetTimestamps_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TimestampServiceServer is the server API for TimestampService service.
// All implementations must embed UnimplementedTimestampServiceServer
// for forward compatibility
type TimestampServiceServer interface {
	CreateTimestamp(context.Context, *CreateTimestampRequest) (*CreateTimestampResponse, error)
	GetTimestamps(context.Context, *GetTimestampsRequest) (*GetTimestampsResponse, error)
	mustEmbedUnimplementedTimestampServiceServer()
}

// UnimplementedTimestampServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTimestampServiceServer struct {
}

func (UnimplementedTimestampServiceServer) CreateTimestamp(context.Context, *CreateTimestampRequest) (*CreateTimestampResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTimestamp not implemented")
}
func (UnimplementedTimestampServiceServer) GetTimestamps(context.Context, *GetTimestampsRequest) (*GetTimestampsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTimestamps not implemented")
}
func (UnimplementedTimestampServiceServer) mustEmbedUnimplementedTimestampServiceServer() {}

// UnsafeTimestampServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TimestampServiceServer will
// result in compilation errors.
type UnsafeTimestampServiceServer interface {
	mustEmbedUnimplementedTimestampServiceServer()
}

func RegisterTimestampServiceServer(s grpc.ServiceRegistrar, srv TimestampServiceServer) {
	s.RegisterService(&TimestampService_ServiceDesc, srv)
}

func _TimestampService_CreateTimestamp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTimestampRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimestampServiceServer).CreateTimestamp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TimestampService_CreateTimestamp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimestampServiceServer).CreateTimestamp(ctx, req.(*CreateTimestampRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TimestampService_GetTimestamps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTimestampsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimestampServiceServer).GetTimestamps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TimestampService_GetTimestamps_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimestampServiceServer).GetTimestamps(ctx, req.(*GetTimestampsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TimestampService_ServiceDesc is the grpc.ServiceDesc for TimestampService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TimestampService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "internal.proto.jit.v1.TimestampService",
	HandlerType: (*TimestampServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTimestamp",
			Handler:    _TimestampService_CreateTimestamp_Handler,
		},
		{
			MethodName: "GetTimestamps",
			Handler:    _TimestampService_GetTimestamps_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/proto/jit/v1/timestamp.proto",
}