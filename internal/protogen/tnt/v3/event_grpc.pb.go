// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: tnt/v3/event.proto

package v3

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
	EventService_CreateEquipmentEvent_FullMethodName                       = "/tnt.v3.EventService/CreateEquipmentEvent"
	EventService_LoadEquipmentRelatedEntities_FullMethodName               = "/tnt.v3.EventService/LoadEquipmentRelatedEntities"
	EventService_CreateOperationsEvent_FullMethodName                      = "/tnt.v3.EventService/CreateOperationsEvent"
	EventService_LoadOperationsRelatedEntities_FullMethodName              = "/tnt.v3.EventService/LoadOperationsRelatedEntities"
	EventService_CreateTransportEvent_FullMethodName                       = "/tnt.v3.EventService/CreateTransportEvent"
	EventService_LoadTransportRelatedEntities_FullMethodName               = "/tnt.v3.EventService/LoadTransportRelatedEntities"
	EventService_CreateShipmentEvent_FullMethodName                        = "/tnt.v3.EventService/CreateShipmentEvent"
	EventService_CreateShipmentEventFromBooking_FullMethodName             = "/tnt.v3.EventService/CreateShipmentEventFromBooking"
	EventService_CreateShipmentEventFromShippingInstruction_FullMethodName = "/tnt.v3.EventService/CreateShipmentEventFromShippingInstruction"
	EventService_LoadShipmentRelatedEntities_FullMethodName                = "/tnt.v3.EventService/LoadShipmentRelatedEntities"
)

// EventServiceClient is the client API for EventService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// The EventService service definition.
type EventServiceClient interface {
	CreateEquipmentEvent(ctx context.Context, in *CreateEquipmentEventRequest, opts ...grpc.CallOption) (*CreateEquipmentEventResponse, error)
	LoadEquipmentRelatedEntities(ctx context.Context, in *LoadEquipmentRelatedEntitiesRequest, opts ...grpc.CallOption) (*LoadEquipmentRelatedEntitiesResponse, error)
	CreateOperationsEvent(ctx context.Context, in *CreateOperationsEventRequest, opts ...grpc.CallOption) (*CreateOperationsEventResponse, error)
	LoadOperationsRelatedEntities(ctx context.Context, in *LoadOperationsRelatedEntitiesRequest, opts ...grpc.CallOption) (*LoadOperationsRelatedEntitiesResponse, error)
	CreateTransportEvent(ctx context.Context, in *CreateTransportEventRequest, opts ...grpc.CallOption) (*CreateTransportEventResponse, error)
	LoadTransportRelatedEntities(ctx context.Context, in *LoadTransportRelatedEntitiesRequest, opts ...grpc.CallOption) (*LoadTransportRelatedEntitiesResponse, error)
	CreateShipmentEvent(ctx context.Context, in *CreateShipmentEventRequest, opts ...grpc.CallOption) (*CreateShipmentEventResponse, error)
	CreateShipmentEventFromBooking(ctx context.Context, in *CreateShipmentEventFromBookingRequest, opts ...grpc.CallOption) (*CreateShipmentEventFromBookingResponse, error)
	CreateShipmentEventFromShippingInstruction(ctx context.Context, in *CreateShipmentEventFromShippingInstructionRequest, opts ...grpc.CallOption) (*CreateShipmentEventFromShippingInstructionResponse, error)
	LoadShipmentRelatedEntities(ctx context.Context, in *LoadShipmentRelatedEntitiesRequest, opts ...grpc.CallOption) (*LoadShipmentRelatedEntitiesResponse, error)
}

type eventServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEventServiceClient(cc grpc.ClientConnInterface) EventServiceClient {
	return &eventServiceClient{cc}
}

func (c *eventServiceClient) CreateEquipmentEvent(ctx context.Context, in *CreateEquipmentEventRequest, opts ...grpc.CallOption) (*CreateEquipmentEventResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateEquipmentEventResponse)
	err := c.cc.Invoke(ctx, EventService_CreateEquipmentEvent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) LoadEquipmentRelatedEntities(ctx context.Context, in *LoadEquipmentRelatedEntitiesRequest, opts ...grpc.CallOption) (*LoadEquipmentRelatedEntitiesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LoadEquipmentRelatedEntitiesResponse)
	err := c.cc.Invoke(ctx, EventService_LoadEquipmentRelatedEntities_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) CreateOperationsEvent(ctx context.Context, in *CreateOperationsEventRequest, opts ...grpc.CallOption) (*CreateOperationsEventResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateOperationsEventResponse)
	err := c.cc.Invoke(ctx, EventService_CreateOperationsEvent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) LoadOperationsRelatedEntities(ctx context.Context, in *LoadOperationsRelatedEntitiesRequest, opts ...grpc.CallOption) (*LoadOperationsRelatedEntitiesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LoadOperationsRelatedEntitiesResponse)
	err := c.cc.Invoke(ctx, EventService_LoadOperationsRelatedEntities_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) CreateTransportEvent(ctx context.Context, in *CreateTransportEventRequest, opts ...grpc.CallOption) (*CreateTransportEventResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateTransportEventResponse)
	err := c.cc.Invoke(ctx, EventService_CreateTransportEvent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) LoadTransportRelatedEntities(ctx context.Context, in *LoadTransportRelatedEntitiesRequest, opts ...grpc.CallOption) (*LoadTransportRelatedEntitiesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LoadTransportRelatedEntitiesResponse)
	err := c.cc.Invoke(ctx, EventService_LoadTransportRelatedEntities_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) CreateShipmentEvent(ctx context.Context, in *CreateShipmentEventRequest, opts ...grpc.CallOption) (*CreateShipmentEventResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateShipmentEventResponse)
	err := c.cc.Invoke(ctx, EventService_CreateShipmentEvent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) CreateShipmentEventFromBooking(ctx context.Context, in *CreateShipmentEventFromBookingRequest, opts ...grpc.CallOption) (*CreateShipmentEventFromBookingResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateShipmentEventFromBookingResponse)
	err := c.cc.Invoke(ctx, EventService_CreateShipmentEventFromBooking_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) CreateShipmentEventFromShippingInstruction(ctx context.Context, in *CreateShipmentEventFromShippingInstructionRequest, opts ...grpc.CallOption) (*CreateShipmentEventFromShippingInstructionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateShipmentEventFromShippingInstructionResponse)
	err := c.cc.Invoke(ctx, EventService_CreateShipmentEventFromShippingInstruction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) LoadShipmentRelatedEntities(ctx context.Context, in *LoadShipmentRelatedEntitiesRequest, opts ...grpc.CallOption) (*LoadShipmentRelatedEntitiesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LoadShipmentRelatedEntitiesResponse)
	err := c.cc.Invoke(ctx, EventService_LoadShipmentRelatedEntities_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventServiceServer is the server API for EventService service.
// All implementations must embed UnimplementedEventServiceServer
// for forward compatibility.
//
// The EventService service definition.
type EventServiceServer interface {
	CreateEquipmentEvent(context.Context, *CreateEquipmentEventRequest) (*CreateEquipmentEventResponse, error)
	LoadEquipmentRelatedEntities(context.Context, *LoadEquipmentRelatedEntitiesRequest) (*LoadEquipmentRelatedEntitiesResponse, error)
	CreateOperationsEvent(context.Context, *CreateOperationsEventRequest) (*CreateOperationsEventResponse, error)
	LoadOperationsRelatedEntities(context.Context, *LoadOperationsRelatedEntitiesRequest) (*LoadOperationsRelatedEntitiesResponse, error)
	CreateTransportEvent(context.Context, *CreateTransportEventRequest) (*CreateTransportEventResponse, error)
	LoadTransportRelatedEntities(context.Context, *LoadTransportRelatedEntitiesRequest) (*LoadTransportRelatedEntitiesResponse, error)
	CreateShipmentEvent(context.Context, *CreateShipmentEventRequest) (*CreateShipmentEventResponse, error)
	CreateShipmentEventFromBooking(context.Context, *CreateShipmentEventFromBookingRequest) (*CreateShipmentEventFromBookingResponse, error)
	CreateShipmentEventFromShippingInstruction(context.Context, *CreateShipmentEventFromShippingInstructionRequest) (*CreateShipmentEventFromShippingInstructionResponse, error)
	LoadShipmentRelatedEntities(context.Context, *LoadShipmentRelatedEntitiesRequest) (*LoadShipmentRelatedEntitiesResponse, error)
	mustEmbedUnimplementedEventServiceServer()
}

// UnimplementedEventServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedEventServiceServer struct{}

func (UnimplementedEventServiceServer) CreateEquipmentEvent(context.Context, *CreateEquipmentEventRequest) (*CreateEquipmentEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEquipmentEvent not implemented")
}
func (UnimplementedEventServiceServer) LoadEquipmentRelatedEntities(context.Context, *LoadEquipmentRelatedEntitiesRequest) (*LoadEquipmentRelatedEntitiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadEquipmentRelatedEntities not implemented")
}
func (UnimplementedEventServiceServer) CreateOperationsEvent(context.Context, *CreateOperationsEventRequest) (*CreateOperationsEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOperationsEvent not implemented")
}
func (UnimplementedEventServiceServer) LoadOperationsRelatedEntities(context.Context, *LoadOperationsRelatedEntitiesRequest) (*LoadOperationsRelatedEntitiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadOperationsRelatedEntities not implemented")
}
func (UnimplementedEventServiceServer) CreateTransportEvent(context.Context, *CreateTransportEventRequest) (*CreateTransportEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTransportEvent not implemented")
}
func (UnimplementedEventServiceServer) LoadTransportRelatedEntities(context.Context, *LoadTransportRelatedEntitiesRequest) (*LoadTransportRelatedEntitiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadTransportRelatedEntities not implemented")
}
func (UnimplementedEventServiceServer) CreateShipmentEvent(context.Context, *CreateShipmentEventRequest) (*CreateShipmentEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateShipmentEvent not implemented")
}
func (UnimplementedEventServiceServer) CreateShipmentEventFromBooking(context.Context, *CreateShipmentEventFromBookingRequest) (*CreateShipmentEventFromBookingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateShipmentEventFromBooking not implemented")
}
func (UnimplementedEventServiceServer) CreateShipmentEventFromShippingInstruction(context.Context, *CreateShipmentEventFromShippingInstructionRequest) (*CreateShipmentEventFromShippingInstructionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateShipmentEventFromShippingInstruction not implemented")
}
func (UnimplementedEventServiceServer) LoadShipmentRelatedEntities(context.Context, *LoadShipmentRelatedEntitiesRequest) (*LoadShipmentRelatedEntitiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadShipmentRelatedEntities not implemented")
}
func (UnimplementedEventServiceServer) mustEmbedUnimplementedEventServiceServer() {}
func (UnimplementedEventServiceServer) testEmbeddedByValue()                      {}

// UnsafeEventServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventServiceServer will
// result in compilation errors.
type UnsafeEventServiceServer interface {
	mustEmbedUnimplementedEventServiceServer()
}

func RegisterEventServiceServer(s grpc.ServiceRegistrar, srv EventServiceServer) {
	// If the following call pancis, it indicates UnimplementedEventServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&EventService_ServiceDesc, srv)
}

func _EventService_CreateEquipmentEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEquipmentEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).CreateEquipmentEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventService_CreateEquipmentEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).CreateEquipmentEvent(ctx, req.(*CreateEquipmentEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_LoadEquipmentRelatedEntities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadEquipmentRelatedEntitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).LoadEquipmentRelatedEntities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventService_LoadEquipmentRelatedEntities_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).LoadEquipmentRelatedEntities(ctx, req.(*LoadEquipmentRelatedEntitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_CreateOperationsEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOperationsEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).CreateOperationsEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventService_CreateOperationsEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).CreateOperationsEvent(ctx, req.(*CreateOperationsEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_LoadOperationsRelatedEntities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadOperationsRelatedEntitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).LoadOperationsRelatedEntities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventService_LoadOperationsRelatedEntities_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).LoadOperationsRelatedEntities(ctx, req.(*LoadOperationsRelatedEntitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_CreateTransportEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTransportEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).CreateTransportEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventService_CreateTransportEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).CreateTransportEvent(ctx, req.(*CreateTransportEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_LoadTransportRelatedEntities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadTransportRelatedEntitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).LoadTransportRelatedEntities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventService_LoadTransportRelatedEntities_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).LoadTransportRelatedEntities(ctx, req.(*LoadTransportRelatedEntitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_CreateShipmentEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateShipmentEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).CreateShipmentEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventService_CreateShipmentEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).CreateShipmentEvent(ctx, req.(*CreateShipmentEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_CreateShipmentEventFromBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateShipmentEventFromBookingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).CreateShipmentEventFromBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventService_CreateShipmentEventFromBooking_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).CreateShipmentEventFromBooking(ctx, req.(*CreateShipmentEventFromBookingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_CreateShipmentEventFromShippingInstruction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateShipmentEventFromShippingInstructionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).CreateShipmentEventFromShippingInstruction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventService_CreateShipmentEventFromShippingInstruction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).CreateShipmentEventFromShippingInstruction(ctx, req.(*CreateShipmentEventFromShippingInstructionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_LoadShipmentRelatedEntities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadShipmentRelatedEntitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).LoadShipmentRelatedEntities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventService_LoadShipmentRelatedEntities_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).LoadShipmentRelatedEntities(ctx, req.(*LoadShipmentRelatedEntitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EventService_ServiceDesc is the grpc.ServiceDesc for EventService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EventService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tnt.v3.EventService",
	HandlerType: (*EventServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEquipmentEvent",
			Handler:    _EventService_CreateEquipmentEvent_Handler,
		},
		{
			MethodName: "LoadEquipmentRelatedEntities",
			Handler:    _EventService_LoadEquipmentRelatedEntities_Handler,
		},
		{
			MethodName: "CreateOperationsEvent",
			Handler:    _EventService_CreateOperationsEvent_Handler,
		},
		{
			MethodName: "LoadOperationsRelatedEntities",
			Handler:    _EventService_LoadOperationsRelatedEntities_Handler,
		},
		{
			MethodName: "CreateTransportEvent",
			Handler:    _EventService_CreateTransportEvent_Handler,
		},
		{
			MethodName: "LoadTransportRelatedEntities",
			Handler:    _EventService_LoadTransportRelatedEntities_Handler,
		},
		{
			MethodName: "CreateShipmentEvent",
			Handler:    _EventService_CreateShipmentEvent_Handler,
		},
		{
			MethodName: "CreateShipmentEventFromBooking",
			Handler:    _EventService_CreateShipmentEventFromBooking_Handler,
		},
		{
			MethodName: "CreateShipmentEventFromShippingInstruction",
			Handler:    _EventService_CreateShipmentEventFromShippingInstruction_Handler,
		},
		{
			MethodName: "LoadShipmentRelatedEntities",
			Handler:    _EventService_LoadShipmentRelatedEntities_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tnt/v3/event.proto",
}
