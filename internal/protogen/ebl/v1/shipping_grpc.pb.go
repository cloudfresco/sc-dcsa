// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: ebl/v1/shipping.proto

package v1

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
	ShippingService_CreateShipment_FullMethodName                                     = "/ebl.v1.ShippingService/CreateShipment"
	ShippingService_GetShipments_FullMethodName                                       = "/ebl.v1.ShippingService/GetShipments"
	ShippingService_GetShipment_FullMethodName                                        = "/ebl.v1.ShippingService/GetShipment"
	ShippingService_GetShipmentByPk_FullMethodName                                    = "/ebl.v1.ShippingService/GetShipmentByPk"
	ShippingService_CreateTransport_FullMethodName                                    = "/ebl.v1.ShippingService/CreateTransport"
	ShippingService_GetTransports_FullMethodName                                      = "/ebl.v1.ShippingService/GetTransports"
	ShippingService_GetTransport_FullMethodName                                       = "/ebl.v1.ShippingService/GetTransport"
	ShippingService_GetTransportByPk_FullMethodName                                   = "/ebl.v1.ShippingService/GetTransportByPk"
	ShippingService_FindCarrierBookingReferenceByShippingInstructionId_FullMethodName = "/ebl.v1.ShippingService/FindCarrierBookingReferenceByShippingInstructionId"
)

// ShippingServiceClient is the client API for ShippingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// The ShippingService service definition.
type ShippingServiceClient interface {
	CreateShipment(ctx context.Context, in *CreateShipmentRequest, opts ...grpc.CallOption) (*CreateShipmentResponse, error)
	GetShipments(ctx context.Context, in *GetShipmentsRequest, opts ...grpc.CallOption) (*GetShipmentsResponse, error)
	GetShipment(ctx context.Context, in *GetShipmentRequest, opts ...grpc.CallOption) (*GetShipmentResponse, error)
	GetShipmentByPk(ctx context.Context, in *GetShipmentByPkRequest, opts ...grpc.CallOption) (*GetShipmentByPkResponse, error)
	CreateTransport(ctx context.Context, in *CreateTransportRequest, opts ...grpc.CallOption) (*CreateTransportResponse, error)
	GetTransports(ctx context.Context, in *GetTransportsRequest, opts ...grpc.CallOption) (*GetTransportsResponse, error)
	GetTransport(ctx context.Context, in *GetTransportRequest, opts ...grpc.CallOption) (*GetTransportResponse, error)
	GetTransportByPk(ctx context.Context, in *GetTransportByPkRequest, opts ...grpc.CallOption) (*GetTransportByPkResponse, error)
	FindCarrierBookingReferenceByShippingInstructionId(ctx context.Context, in *FindCarrierBookingReferenceByShippingInstructionIdRequest, opts ...grpc.CallOption) (*FindCarrierBookingReferenceByShippingInstructionIdResponse, error)
}

type shippingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewShippingServiceClient(cc grpc.ClientConnInterface) ShippingServiceClient {
	return &shippingServiceClient{cc}
}

func (c *shippingServiceClient) CreateShipment(ctx context.Context, in *CreateShipmentRequest, opts ...grpc.CallOption) (*CreateShipmentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateShipmentResponse)
	err := c.cc.Invoke(ctx, ShippingService_CreateShipment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shippingServiceClient) GetShipments(ctx context.Context, in *GetShipmentsRequest, opts ...grpc.CallOption) (*GetShipmentsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetShipmentsResponse)
	err := c.cc.Invoke(ctx, ShippingService_GetShipments_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shippingServiceClient) GetShipment(ctx context.Context, in *GetShipmentRequest, opts ...grpc.CallOption) (*GetShipmentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetShipmentResponse)
	err := c.cc.Invoke(ctx, ShippingService_GetShipment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shippingServiceClient) GetShipmentByPk(ctx context.Context, in *GetShipmentByPkRequest, opts ...grpc.CallOption) (*GetShipmentByPkResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetShipmentByPkResponse)
	err := c.cc.Invoke(ctx, ShippingService_GetShipmentByPk_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shippingServiceClient) CreateTransport(ctx context.Context, in *CreateTransportRequest, opts ...grpc.CallOption) (*CreateTransportResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateTransportResponse)
	err := c.cc.Invoke(ctx, ShippingService_CreateTransport_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shippingServiceClient) GetTransports(ctx context.Context, in *GetTransportsRequest, opts ...grpc.CallOption) (*GetTransportsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTransportsResponse)
	err := c.cc.Invoke(ctx, ShippingService_GetTransports_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shippingServiceClient) GetTransport(ctx context.Context, in *GetTransportRequest, opts ...grpc.CallOption) (*GetTransportResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTransportResponse)
	err := c.cc.Invoke(ctx, ShippingService_GetTransport_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shippingServiceClient) GetTransportByPk(ctx context.Context, in *GetTransportByPkRequest, opts ...grpc.CallOption) (*GetTransportByPkResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTransportByPkResponse)
	err := c.cc.Invoke(ctx, ShippingService_GetTransportByPk_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shippingServiceClient) FindCarrierBookingReferenceByShippingInstructionId(ctx context.Context, in *FindCarrierBookingReferenceByShippingInstructionIdRequest, opts ...grpc.CallOption) (*FindCarrierBookingReferenceByShippingInstructionIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FindCarrierBookingReferenceByShippingInstructionIdResponse)
	err := c.cc.Invoke(ctx, ShippingService_FindCarrierBookingReferenceByShippingInstructionId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShippingServiceServer is the server API for ShippingService service.
// All implementations must embed UnimplementedShippingServiceServer
// for forward compatibility.
//
// The ShippingService service definition.
type ShippingServiceServer interface {
	CreateShipment(context.Context, *CreateShipmentRequest) (*CreateShipmentResponse, error)
	GetShipments(context.Context, *GetShipmentsRequest) (*GetShipmentsResponse, error)
	GetShipment(context.Context, *GetShipmentRequest) (*GetShipmentResponse, error)
	GetShipmentByPk(context.Context, *GetShipmentByPkRequest) (*GetShipmentByPkResponse, error)
	CreateTransport(context.Context, *CreateTransportRequest) (*CreateTransportResponse, error)
	GetTransports(context.Context, *GetTransportsRequest) (*GetTransportsResponse, error)
	GetTransport(context.Context, *GetTransportRequest) (*GetTransportResponse, error)
	GetTransportByPk(context.Context, *GetTransportByPkRequest) (*GetTransportByPkResponse, error)
	FindCarrierBookingReferenceByShippingInstructionId(context.Context, *FindCarrierBookingReferenceByShippingInstructionIdRequest) (*FindCarrierBookingReferenceByShippingInstructionIdResponse, error)
	mustEmbedUnimplementedShippingServiceServer()
}

// UnimplementedShippingServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedShippingServiceServer struct{}

func (UnimplementedShippingServiceServer) CreateShipment(context.Context, *CreateShipmentRequest) (*CreateShipmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateShipment not implemented")
}
func (UnimplementedShippingServiceServer) GetShipments(context.Context, *GetShipmentsRequest) (*GetShipmentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShipments not implemented")
}
func (UnimplementedShippingServiceServer) GetShipment(context.Context, *GetShipmentRequest) (*GetShipmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShipment not implemented")
}
func (UnimplementedShippingServiceServer) GetShipmentByPk(context.Context, *GetShipmentByPkRequest) (*GetShipmentByPkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShipmentByPk not implemented")
}
func (UnimplementedShippingServiceServer) CreateTransport(context.Context, *CreateTransportRequest) (*CreateTransportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTransport not implemented")
}
func (UnimplementedShippingServiceServer) GetTransports(context.Context, *GetTransportsRequest) (*GetTransportsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransports not implemented")
}
func (UnimplementedShippingServiceServer) GetTransport(context.Context, *GetTransportRequest) (*GetTransportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransport not implemented")
}
func (UnimplementedShippingServiceServer) GetTransportByPk(context.Context, *GetTransportByPkRequest) (*GetTransportByPkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransportByPk not implemented")
}
func (UnimplementedShippingServiceServer) FindCarrierBookingReferenceByShippingInstructionId(context.Context, *FindCarrierBookingReferenceByShippingInstructionIdRequest) (*FindCarrierBookingReferenceByShippingInstructionIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindCarrierBookingReferenceByShippingInstructionId not implemented")
}
func (UnimplementedShippingServiceServer) mustEmbedUnimplementedShippingServiceServer() {}
func (UnimplementedShippingServiceServer) testEmbeddedByValue()                         {}

// UnsafeShippingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShippingServiceServer will
// result in compilation errors.
type UnsafeShippingServiceServer interface {
	mustEmbedUnimplementedShippingServiceServer()
}

func RegisterShippingServiceServer(s grpc.ServiceRegistrar, srv ShippingServiceServer) {
	// If the following call pancis, it indicates UnimplementedShippingServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ShippingService_ServiceDesc, srv)
}

func _ShippingService_CreateShipment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateShipmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShippingServiceServer).CreateShipment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShippingService_CreateShipment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShippingServiceServer).CreateShipment(ctx, req.(*CreateShipmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShippingService_GetShipments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShipmentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShippingServiceServer).GetShipments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShippingService_GetShipments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShippingServiceServer).GetShipments(ctx, req.(*GetShipmentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShippingService_GetShipment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShipmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShippingServiceServer).GetShipment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShippingService_GetShipment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShippingServiceServer).GetShipment(ctx, req.(*GetShipmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShippingService_GetShipmentByPk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShipmentByPkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShippingServiceServer).GetShipmentByPk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShippingService_GetShipmentByPk_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShippingServiceServer).GetShipmentByPk(ctx, req.(*GetShipmentByPkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShippingService_CreateTransport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTransportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShippingServiceServer).CreateTransport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShippingService_CreateTransport_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShippingServiceServer).CreateTransport(ctx, req.(*CreateTransportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShippingService_GetTransports_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransportsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShippingServiceServer).GetTransports(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShippingService_GetTransports_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShippingServiceServer).GetTransports(ctx, req.(*GetTransportsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShippingService_GetTransport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShippingServiceServer).GetTransport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShippingService_GetTransport_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShippingServiceServer).GetTransport(ctx, req.(*GetTransportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShippingService_GetTransportByPk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransportByPkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShippingServiceServer).GetTransportByPk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShippingService_GetTransportByPk_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShippingServiceServer).GetTransportByPk(ctx, req.(*GetTransportByPkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShippingService_FindCarrierBookingReferenceByShippingInstructionId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindCarrierBookingReferenceByShippingInstructionIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShippingServiceServer).FindCarrierBookingReferenceByShippingInstructionId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShippingService_FindCarrierBookingReferenceByShippingInstructionId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShippingServiceServer).FindCarrierBookingReferenceByShippingInstructionId(ctx, req.(*FindCarrierBookingReferenceByShippingInstructionIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ShippingService_ServiceDesc is the grpc.ServiceDesc for ShippingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ShippingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ebl.v1.ShippingService",
	HandlerType: (*ShippingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateShipment",
			Handler:    _ShippingService_CreateShipment_Handler,
		},
		{
			MethodName: "GetShipments",
			Handler:    _ShippingService_GetShipments_Handler,
		},
		{
			MethodName: "GetShipment",
			Handler:    _ShippingService_GetShipment_Handler,
		},
		{
			MethodName: "GetShipmentByPk",
			Handler:    _ShippingService_GetShipmentByPk_Handler,
		},
		{
			MethodName: "CreateTransport",
			Handler:    _ShippingService_CreateTransport_Handler,
		},
		{
			MethodName: "GetTransports",
			Handler:    _ShippingService_GetTransports_Handler,
		},
		{
			MethodName: "GetTransport",
			Handler:    _ShippingService_GetTransport_Handler,
		},
		{
			MethodName: "GetTransportByPk",
			Handler:    _ShippingService_GetTransportByPk_Handler,
		},
		{
			MethodName: "FindCarrierBookingReferenceByShippingInstructionId",
			Handler:    _ShippingService_FindCarrierBookingReferenceByShippingInstructionId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ebl/v1/shipping.proto",
}