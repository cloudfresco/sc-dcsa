// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: internal/proto/ebl/v1/issuerequestresponse.proto

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
	IssueRequestResponseService_CreateIssuanceRequestResponse_FullMethodName = "/internal.proto.ebl.v1.IssueRequestResponseService/CreateIssuanceRequestResponse"
	IssueRequestResponseService_UpdateIssuanceRequestResponse_FullMethodName = "/internal.proto.ebl.v1.IssueRequestResponseService/UpdateIssuanceRequestResponse"
)

// IssueRequestResponseServiceClient is the client API for IssueRequestResponseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IssueRequestResponseServiceClient interface {
	CreateIssuanceRequestResponse(ctx context.Context, in *CreateIssuanceRequestResponseRequest, opts ...grpc.CallOption) (*CreateIssuanceRequestResponseResponse, error)
	UpdateIssuanceRequestResponse(ctx context.Context, in *UpdateIssuanceRequestResponseRequest, opts ...grpc.CallOption) (*UpdateIssuanceRequestResponseResponse, error)
}

type issueRequestResponseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIssueRequestResponseServiceClient(cc grpc.ClientConnInterface) IssueRequestResponseServiceClient {
	return &issueRequestResponseServiceClient{cc}
}

func (c *issueRequestResponseServiceClient) CreateIssuanceRequestResponse(ctx context.Context, in *CreateIssuanceRequestResponseRequest, opts ...grpc.CallOption) (*CreateIssuanceRequestResponseResponse, error) {
	out := new(CreateIssuanceRequestResponseResponse)
	err := c.cc.Invoke(ctx, IssueRequestResponseService_CreateIssuanceRequestResponse_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *issueRequestResponseServiceClient) UpdateIssuanceRequestResponse(ctx context.Context, in *UpdateIssuanceRequestResponseRequest, opts ...grpc.CallOption) (*UpdateIssuanceRequestResponseResponse, error) {
	out := new(UpdateIssuanceRequestResponseResponse)
	err := c.cc.Invoke(ctx, IssueRequestResponseService_UpdateIssuanceRequestResponse_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IssueRequestResponseServiceServer is the server API for IssueRequestResponseService service.
// All implementations must embed UnimplementedIssueRequestResponseServiceServer
// for forward compatibility
type IssueRequestResponseServiceServer interface {
	CreateIssuanceRequestResponse(context.Context, *CreateIssuanceRequestResponseRequest) (*CreateIssuanceRequestResponseResponse, error)
	UpdateIssuanceRequestResponse(context.Context, *UpdateIssuanceRequestResponseRequest) (*UpdateIssuanceRequestResponseResponse, error)
	mustEmbedUnimplementedIssueRequestResponseServiceServer()
}

// UnimplementedIssueRequestResponseServiceServer must be embedded to have forward compatible implementations.
type UnimplementedIssueRequestResponseServiceServer struct {
}

func (UnimplementedIssueRequestResponseServiceServer) CreateIssuanceRequestResponse(context.Context, *CreateIssuanceRequestResponseRequest) (*CreateIssuanceRequestResponseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateIssuanceRequestResponse not implemented")
}
func (UnimplementedIssueRequestResponseServiceServer) UpdateIssuanceRequestResponse(context.Context, *UpdateIssuanceRequestResponseRequest) (*UpdateIssuanceRequestResponseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateIssuanceRequestResponse not implemented")
}
func (UnimplementedIssueRequestResponseServiceServer) mustEmbedUnimplementedIssueRequestResponseServiceServer() {
}

// UnsafeIssueRequestResponseServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IssueRequestResponseServiceServer will
// result in compilation errors.
type UnsafeIssueRequestResponseServiceServer interface {
	mustEmbedUnimplementedIssueRequestResponseServiceServer()
}

func RegisterIssueRequestResponseServiceServer(s grpc.ServiceRegistrar, srv IssueRequestResponseServiceServer) {
	s.RegisterService(&IssueRequestResponseService_ServiceDesc, srv)
}

func _IssueRequestResponseService_CreateIssuanceRequestResponse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateIssuanceRequestResponseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IssueRequestResponseServiceServer).CreateIssuanceRequestResponse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IssueRequestResponseService_CreateIssuanceRequestResponse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IssueRequestResponseServiceServer).CreateIssuanceRequestResponse(ctx, req.(*CreateIssuanceRequestResponseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IssueRequestResponseService_UpdateIssuanceRequestResponse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateIssuanceRequestResponseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IssueRequestResponseServiceServer).UpdateIssuanceRequestResponse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IssueRequestResponseService_UpdateIssuanceRequestResponse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IssueRequestResponseServiceServer).UpdateIssuanceRequestResponse(ctx, req.(*UpdateIssuanceRequestResponseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// IssueRequestResponseService_ServiceDesc is the grpc.ServiceDesc for IssueRequestResponseService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IssueRequestResponseService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "internal.proto.ebl.v1.IssueRequestResponseService",
	HandlerType: (*IssueRequestResponseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateIssuanceRequestResponse",
			Handler:    _IssueRequestResponseService_CreateIssuanceRequestResponse_Handler,
		},
		{
			MethodName: "UpdateIssuanceRequestResponse",
			Handler:    _IssueRequestResponseService_UpdateIssuanceRequestResponse_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/proto/ebl/v1/issuerequestresponse.proto",
}