// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: tnt/v3/eventsubscription.proto

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
	EventSubscriptionService_CreateEventSubscription_FullMethodName     = "/tnt.v3.EventSubscriptionService/CreateEventSubscription"
	EventSubscriptionService_GetEventSubscriptions_FullMethodName       = "/tnt.v3.EventSubscriptionService/GetEventSubscriptions"
	EventSubscriptionService_FindEventSubscriptionByID_FullMethodName   = "/tnt.v3.EventSubscriptionService/FindEventSubscriptionByID"
	EventSubscriptionService_DeleteEventSubscriptionByID_FullMethodName = "/tnt.v3.EventSubscriptionService/DeleteEventSubscriptionByID"
	EventSubscriptionService_UpdateEventSubscription_FullMethodName     = "/tnt.v3.EventSubscriptionService/UpdateEventSubscription"
)

// EventSubscriptionServiceClient is the client API for EventSubscriptionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// The EventSubscriptionService service definition.
type EventSubscriptionServiceClient interface {
	CreateEventSubscription(ctx context.Context, in *CreateEventSubscriptionRequest, opts ...grpc.CallOption) (*CreateEventSubscriptionResponse, error)
	GetEventSubscriptions(ctx context.Context, in *GetEventSubscriptionsRequest, opts ...grpc.CallOption) (*GetEventSubscriptionsResponse, error)
	FindEventSubscriptionByID(ctx context.Context, in *FindEventSubscriptionByIDRequest, opts ...grpc.CallOption) (*FindEventSubscriptionByIDResponse, error)
	DeleteEventSubscriptionByID(ctx context.Context, in *DeleteEventSubscriptionByIDRequest, opts ...grpc.CallOption) (*DeleteEventSubscriptionByIDResponse, error)
	UpdateEventSubscription(ctx context.Context, in *UpdateEventSubscriptionRequest, opts ...grpc.CallOption) (*UpdateEventSubscriptionResponse, error)
}

type eventSubscriptionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEventSubscriptionServiceClient(cc grpc.ClientConnInterface) EventSubscriptionServiceClient {
	return &eventSubscriptionServiceClient{cc}
}

func (c *eventSubscriptionServiceClient) CreateEventSubscription(ctx context.Context, in *CreateEventSubscriptionRequest, opts ...grpc.CallOption) (*CreateEventSubscriptionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateEventSubscriptionResponse)
	err := c.cc.Invoke(ctx, EventSubscriptionService_CreateEventSubscription_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventSubscriptionServiceClient) GetEventSubscriptions(ctx context.Context, in *GetEventSubscriptionsRequest, opts ...grpc.CallOption) (*GetEventSubscriptionsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetEventSubscriptionsResponse)
	err := c.cc.Invoke(ctx, EventSubscriptionService_GetEventSubscriptions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventSubscriptionServiceClient) FindEventSubscriptionByID(ctx context.Context, in *FindEventSubscriptionByIDRequest, opts ...grpc.CallOption) (*FindEventSubscriptionByIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FindEventSubscriptionByIDResponse)
	err := c.cc.Invoke(ctx, EventSubscriptionService_FindEventSubscriptionByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventSubscriptionServiceClient) DeleteEventSubscriptionByID(ctx context.Context, in *DeleteEventSubscriptionByIDRequest, opts ...grpc.CallOption) (*DeleteEventSubscriptionByIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteEventSubscriptionByIDResponse)
	err := c.cc.Invoke(ctx, EventSubscriptionService_DeleteEventSubscriptionByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventSubscriptionServiceClient) UpdateEventSubscription(ctx context.Context, in *UpdateEventSubscriptionRequest, opts ...grpc.CallOption) (*UpdateEventSubscriptionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateEventSubscriptionResponse)
	err := c.cc.Invoke(ctx, EventSubscriptionService_UpdateEventSubscription_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventSubscriptionServiceServer is the server API for EventSubscriptionService service.
// All implementations must embed UnimplementedEventSubscriptionServiceServer
// for forward compatibility.
//
// The EventSubscriptionService service definition.
type EventSubscriptionServiceServer interface {
	CreateEventSubscription(context.Context, *CreateEventSubscriptionRequest) (*CreateEventSubscriptionResponse, error)
	GetEventSubscriptions(context.Context, *GetEventSubscriptionsRequest) (*GetEventSubscriptionsResponse, error)
	FindEventSubscriptionByID(context.Context, *FindEventSubscriptionByIDRequest) (*FindEventSubscriptionByIDResponse, error)
	DeleteEventSubscriptionByID(context.Context, *DeleteEventSubscriptionByIDRequest) (*DeleteEventSubscriptionByIDResponse, error)
	UpdateEventSubscription(context.Context, *UpdateEventSubscriptionRequest) (*UpdateEventSubscriptionResponse, error)
	mustEmbedUnimplementedEventSubscriptionServiceServer()
}

// UnimplementedEventSubscriptionServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedEventSubscriptionServiceServer struct{}

func (UnimplementedEventSubscriptionServiceServer) CreateEventSubscription(context.Context, *CreateEventSubscriptionRequest) (*CreateEventSubscriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEventSubscription not implemented")
}
func (UnimplementedEventSubscriptionServiceServer) GetEventSubscriptions(context.Context, *GetEventSubscriptionsRequest) (*GetEventSubscriptionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventSubscriptions not implemented")
}
func (UnimplementedEventSubscriptionServiceServer) FindEventSubscriptionByID(context.Context, *FindEventSubscriptionByIDRequest) (*FindEventSubscriptionByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindEventSubscriptionByID not implemented")
}
func (UnimplementedEventSubscriptionServiceServer) DeleteEventSubscriptionByID(context.Context, *DeleteEventSubscriptionByIDRequest) (*DeleteEventSubscriptionByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEventSubscriptionByID not implemented")
}
func (UnimplementedEventSubscriptionServiceServer) UpdateEventSubscription(context.Context, *UpdateEventSubscriptionRequest) (*UpdateEventSubscriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEventSubscription not implemented")
}
func (UnimplementedEventSubscriptionServiceServer) mustEmbedUnimplementedEventSubscriptionServiceServer() {
}
func (UnimplementedEventSubscriptionServiceServer) testEmbeddedByValue() {}

// UnsafeEventSubscriptionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventSubscriptionServiceServer will
// result in compilation errors.
type UnsafeEventSubscriptionServiceServer interface {
	mustEmbedUnimplementedEventSubscriptionServiceServer()
}

func RegisterEventSubscriptionServiceServer(s grpc.ServiceRegistrar, srv EventSubscriptionServiceServer) {
	// If the following call pancis, it indicates UnimplementedEventSubscriptionServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&EventSubscriptionService_ServiceDesc, srv)
}

func _EventSubscriptionService_CreateEventSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEventSubscriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventSubscriptionServiceServer).CreateEventSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventSubscriptionService_CreateEventSubscription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventSubscriptionServiceServer).CreateEventSubscription(ctx, req.(*CreateEventSubscriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventSubscriptionService_GetEventSubscriptions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventSubscriptionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventSubscriptionServiceServer).GetEventSubscriptions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventSubscriptionService_GetEventSubscriptions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventSubscriptionServiceServer).GetEventSubscriptions(ctx, req.(*GetEventSubscriptionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventSubscriptionService_FindEventSubscriptionByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindEventSubscriptionByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventSubscriptionServiceServer).FindEventSubscriptionByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventSubscriptionService_FindEventSubscriptionByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventSubscriptionServiceServer).FindEventSubscriptionByID(ctx, req.(*FindEventSubscriptionByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventSubscriptionService_DeleteEventSubscriptionByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEventSubscriptionByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventSubscriptionServiceServer).DeleteEventSubscriptionByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventSubscriptionService_DeleteEventSubscriptionByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventSubscriptionServiceServer).DeleteEventSubscriptionByID(ctx, req.(*DeleteEventSubscriptionByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventSubscriptionService_UpdateEventSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEventSubscriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventSubscriptionServiceServer).UpdateEventSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventSubscriptionService_UpdateEventSubscription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventSubscriptionServiceServer).UpdateEventSubscription(ctx, req.(*UpdateEventSubscriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EventSubscriptionService_ServiceDesc is the grpc.ServiceDesc for EventSubscriptionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EventSubscriptionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tnt.v3.EventSubscriptionService",
	HandlerType: (*EventSubscriptionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEventSubscription",
			Handler:    _EventSubscriptionService_CreateEventSubscription_Handler,
		},
		{
			MethodName: "GetEventSubscriptions",
			Handler:    _EventSubscriptionService_GetEventSubscriptions_Handler,
		},
		{
			MethodName: "FindEventSubscriptionByID",
			Handler:    _EventSubscriptionService_FindEventSubscriptionByID_Handler,
		},
		{
			MethodName: "DeleteEventSubscriptionByID",
			Handler:    _EventSubscriptionService_DeleteEventSubscriptionByID_Handler,
		},
		{
			MethodName: "UpdateEventSubscription",
			Handler:    _EventSubscriptionService_UpdateEventSubscription_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tnt/v3/eventsubscription.proto",
}
