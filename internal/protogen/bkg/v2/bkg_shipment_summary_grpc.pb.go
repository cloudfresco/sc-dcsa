// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: bkg/v2/bkg_shipment_summary.proto

package v2

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	BkgShipmentSummaryService_CreateBkgShipmentSummary_FullMethodName                       = "/bkg.v2.BkgShipmentSummaryService/CreateBkgShipmentSummary"
	BkgShipmentSummaryService_GetBkgShipmentSummaries_FullMethodName                        = "/bkg.v2.BkgShipmentSummaryService/GetBkgShipmentSummaries"
	BkgShipmentSummaryService_GetBkgShipmentSummaryByCarrierBookingReference_FullMethodName = "/bkg.v2.BkgShipmentSummaryService/GetBkgShipmentSummaryByCarrierBookingReference"
)

// BkgShipmentSummaryServiceClient is the client API for BkgShipmentSummaryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// The BkgShipmentSummaryService service definition.
type BkgShipmentSummaryServiceClient interface {
	CreateBkgShipmentSummary(ctx context.Context, in *CreateBkgShipmentSummaryRequest, opts ...grpc.CallOption) (*CreateBkgShipmentSummaryResponse, error)
	GetBkgShipmentSummaries(ctx context.Context, in *GetBkgShipmentSummariesRequest, opts ...grpc.CallOption) (*GetBkgShipmentSummariesResponse, error)
	GetBkgShipmentSummaryByCarrierBookingReference(ctx context.Context, in *GetBkgShipmentSummaryByCarrierBookingReferenceRequest, opts ...grpc.CallOption) (*GetBkgShipmentSummaryByCarrierBookingReferenceResponse, error)
}

type bkgShipmentSummaryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBkgShipmentSummaryServiceClient(cc grpc.ClientConnInterface) BkgShipmentSummaryServiceClient {
	return &bkgShipmentSummaryServiceClient{cc}
}

func (c *bkgShipmentSummaryServiceClient) CreateBkgShipmentSummary(ctx context.Context, in *CreateBkgShipmentSummaryRequest, opts ...grpc.CallOption) (*CreateBkgShipmentSummaryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateBkgShipmentSummaryResponse)
	err := c.cc.Invoke(ctx, BkgShipmentSummaryService_CreateBkgShipmentSummary_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bkgShipmentSummaryServiceClient) GetBkgShipmentSummaries(ctx context.Context, in *GetBkgShipmentSummariesRequest, opts ...grpc.CallOption) (*GetBkgShipmentSummariesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBkgShipmentSummariesResponse)
	err := c.cc.Invoke(ctx, BkgShipmentSummaryService_GetBkgShipmentSummaries_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bkgShipmentSummaryServiceClient) GetBkgShipmentSummaryByCarrierBookingReference(ctx context.Context, in *GetBkgShipmentSummaryByCarrierBookingReferenceRequest, opts ...grpc.CallOption) (*GetBkgShipmentSummaryByCarrierBookingReferenceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBkgShipmentSummaryByCarrierBookingReferenceResponse)
	err := c.cc.Invoke(ctx, BkgShipmentSummaryService_GetBkgShipmentSummaryByCarrierBookingReference_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BkgShipmentSummaryServiceServer is the server API for BkgShipmentSummaryService service.
// All implementations must embed UnimplementedBkgShipmentSummaryServiceServer
// for forward compatibility.
//
// The BkgShipmentSummaryService service definition.
type BkgShipmentSummaryServiceServer interface {
	CreateBkgShipmentSummary(context.Context, *CreateBkgShipmentSummaryRequest) (*CreateBkgShipmentSummaryResponse, error)
	GetBkgShipmentSummaries(context.Context, *GetBkgShipmentSummariesRequest) (*GetBkgShipmentSummariesResponse, error)
	GetBkgShipmentSummaryByCarrierBookingReference(context.Context, *GetBkgShipmentSummaryByCarrierBookingReferenceRequest) (*GetBkgShipmentSummaryByCarrierBookingReferenceResponse, error)
	mustEmbedUnimplementedBkgShipmentSummaryServiceServer()
}

// UnimplementedBkgShipmentSummaryServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBkgShipmentSummaryServiceServer struct{}

func (UnimplementedBkgShipmentSummaryServiceServer) CreateBkgShipmentSummary(context.Context, *CreateBkgShipmentSummaryRequest) (*CreateBkgShipmentSummaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBkgShipmentSummary not implemented")
}
func (UnimplementedBkgShipmentSummaryServiceServer) GetBkgShipmentSummaries(context.Context, *GetBkgShipmentSummariesRequest) (*GetBkgShipmentSummariesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBkgShipmentSummaries not implemented")
}
func (UnimplementedBkgShipmentSummaryServiceServer) GetBkgShipmentSummaryByCarrierBookingReference(context.Context, *GetBkgShipmentSummaryByCarrierBookingReferenceRequest) (*GetBkgShipmentSummaryByCarrierBookingReferenceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBkgShipmentSummaryByCarrierBookingReference not implemented")
}
func (UnimplementedBkgShipmentSummaryServiceServer) mustEmbedUnimplementedBkgShipmentSummaryServiceServer() {
}
func (UnimplementedBkgShipmentSummaryServiceServer) testEmbeddedByValue() {}

// UnsafeBkgShipmentSummaryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BkgShipmentSummaryServiceServer will
// result in compilation errors.
type UnsafeBkgShipmentSummaryServiceServer interface {
	mustEmbedUnimplementedBkgShipmentSummaryServiceServer()
}

func RegisterBkgShipmentSummaryServiceServer(s grpc.ServiceRegistrar, srv BkgShipmentSummaryServiceServer) {
	// If the following call pancis, it indicates UnimplementedBkgShipmentSummaryServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BkgShipmentSummaryService_ServiceDesc, srv)
}

func _BkgShipmentSummaryService_CreateBkgShipmentSummary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBkgShipmentSummaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BkgShipmentSummaryServiceServer).CreateBkgShipmentSummary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BkgShipmentSummaryService_CreateBkgShipmentSummary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BkgShipmentSummaryServiceServer).CreateBkgShipmentSummary(ctx, req.(*CreateBkgShipmentSummaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BkgShipmentSummaryService_GetBkgShipmentSummaries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBkgShipmentSummariesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BkgShipmentSummaryServiceServer).GetBkgShipmentSummaries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BkgShipmentSummaryService_GetBkgShipmentSummaries_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BkgShipmentSummaryServiceServer).GetBkgShipmentSummaries(ctx, req.(*GetBkgShipmentSummariesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BkgShipmentSummaryService_GetBkgShipmentSummaryByCarrierBookingReference_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBkgShipmentSummaryByCarrierBookingReferenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BkgShipmentSummaryServiceServer).GetBkgShipmentSummaryByCarrierBookingReference(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BkgShipmentSummaryService_GetBkgShipmentSummaryByCarrierBookingReference_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BkgShipmentSummaryServiceServer).GetBkgShipmentSummaryByCarrierBookingReference(ctx, req.(*GetBkgShipmentSummaryByCarrierBookingReferenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BkgShipmentSummaryService_ServiceDesc is the grpc.ServiceDesc for BkgShipmentSummaryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BkgShipmentSummaryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bkg.v2.BkgShipmentSummaryService",
	HandlerType: (*BkgShipmentSummaryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBkgShipmentSummary",
			Handler:    _BkgShipmentSummaryService_CreateBkgShipmentSummary_Handler,
		},
		{
			MethodName: "GetBkgShipmentSummaries",
			Handler:    _BkgShipmentSummaryService_GetBkgShipmentSummaries_Handler,
		},
		{
			MethodName: "GetBkgShipmentSummaryByCarrierBookingReference",
			Handler:    _BkgShipmentSummaryService_GetBkgShipmentSummaryByCarrierBookingReference_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bkg/v2/bkg_shipment_summary.proto",
}