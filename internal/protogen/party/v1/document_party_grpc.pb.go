// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: party/v1/document_party.proto

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
	DocumentPartyService_CreateDocumentParty_FullMethodName                            = "/party.v1.DocumentPartyService/CreateDocumentParty"
	DocumentPartyService_CreateDocumentPartiesByBookingID_FullMethodName               = "/party.v1.DocumentPartyService/CreateDocumentPartiesByBookingID"
	DocumentPartyService_CreateDocumentPartiesByShippingInstructionID_FullMethodName   = "/party.v1.DocumentPartyService/CreateDocumentPartiesByShippingInstructionID"
	DocumentPartyService_FetchDocumentPartiesByBookingID_FullMethodName                = "/party.v1.DocumentPartyService/FetchDocumentPartiesByBookingID"
	DocumentPartyService_FetchDocumentPartiesByByShippingInstructionID_FullMethodName  = "/party.v1.DocumentPartyService/FetchDocumentPartiesByByShippingInstructionID"
	DocumentPartyService_ResolveDocumentPartiesForShippingInstructionID_FullMethodName = "/party.v1.DocumentPartyService/ResolveDocumentPartiesForShippingInstructionID"
)

// DocumentPartyServiceClient is the client API for DocumentPartyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// The DocumentPartyService service definition.
type DocumentPartyServiceClient interface {
	CreateDocumentParty(ctx context.Context, in *CreateDocumentPartyRequest, opts ...grpc.CallOption) (*CreateDocumentPartyResponse, error)
	CreateDocumentPartiesByBookingID(ctx context.Context, in *CreateDocumentPartiesByBookingIDRequest, opts ...grpc.CallOption) (*CreateDocumentPartiesByBookingIDResponse, error)
	CreateDocumentPartiesByShippingInstructionID(ctx context.Context, in *CreateDocumentPartiesByShippingInstructionIDRequest, opts ...grpc.CallOption) (*CreateDocumentPartiesByShippingInstructionIDResponse, error)
	FetchDocumentPartiesByBookingID(ctx context.Context, in *FetchDocumentPartiesByBookingIDRequest, opts ...grpc.CallOption) (*FetchDocumentPartiesByBookingIDResponse, error)
	FetchDocumentPartiesByByShippingInstructionID(ctx context.Context, in *FetchDocumentPartiesByByShippingInstructionIDRequest, opts ...grpc.CallOption) (*FetchDocumentPartiesByByShippingInstructionIDResponse, error)
	ResolveDocumentPartiesForShippingInstructionID(ctx context.Context, in *ResolveDocumentPartiesForShippingInstructionIDRequest, opts ...grpc.CallOption) (*ResolveDocumentPartiesForShippingInstructionIDResponse, error)
}

type documentPartyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDocumentPartyServiceClient(cc grpc.ClientConnInterface) DocumentPartyServiceClient {
	return &documentPartyServiceClient{cc}
}

func (c *documentPartyServiceClient) CreateDocumentParty(ctx context.Context, in *CreateDocumentPartyRequest, opts ...grpc.CallOption) (*CreateDocumentPartyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateDocumentPartyResponse)
	err := c.cc.Invoke(ctx, DocumentPartyService_CreateDocumentParty_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *documentPartyServiceClient) CreateDocumentPartiesByBookingID(ctx context.Context, in *CreateDocumentPartiesByBookingIDRequest, opts ...grpc.CallOption) (*CreateDocumentPartiesByBookingIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateDocumentPartiesByBookingIDResponse)
	err := c.cc.Invoke(ctx, DocumentPartyService_CreateDocumentPartiesByBookingID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *documentPartyServiceClient) CreateDocumentPartiesByShippingInstructionID(ctx context.Context, in *CreateDocumentPartiesByShippingInstructionIDRequest, opts ...grpc.CallOption) (*CreateDocumentPartiesByShippingInstructionIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateDocumentPartiesByShippingInstructionIDResponse)
	err := c.cc.Invoke(ctx, DocumentPartyService_CreateDocumentPartiesByShippingInstructionID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *documentPartyServiceClient) FetchDocumentPartiesByBookingID(ctx context.Context, in *FetchDocumentPartiesByBookingIDRequest, opts ...grpc.CallOption) (*FetchDocumentPartiesByBookingIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FetchDocumentPartiesByBookingIDResponse)
	err := c.cc.Invoke(ctx, DocumentPartyService_FetchDocumentPartiesByBookingID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *documentPartyServiceClient) FetchDocumentPartiesByByShippingInstructionID(ctx context.Context, in *FetchDocumentPartiesByByShippingInstructionIDRequest, opts ...grpc.CallOption) (*FetchDocumentPartiesByByShippingInstructionIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FetchDocumentPartiesByByShippingInstructionIDResponse)
	err := c.cc.Invoke(ctx, DocumentPartyService_FetchDocumentPartiesByByShippingInstructionID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *documentPartyServiceClient) ResolveDocumentPartiesForShippingInstructionID(ctx context.Context, in *ResolveDocumentPartiesForShippingInstructionIDRequest, opts ...grpc.CallOption) (*ResolveDocumentPartiesForShippingInstructionIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResolveDocumentPartiesForShippingInstructionIDResponse)
	err := c.cc.Invoke(ctx, DocumentPartyService_ResolveDocumentPartiesForShippingInstructionID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DocumentPartyServiceServer is the server API for DocumentPartyService service.
// All implementations must embed UnimplementedDocumentPartyServiceServer
// for forward compatibility.
//
// The DocumentPartyService service definition.
type DocumentPartyServiceServer interface {
	CreateDocumentParty(context.Context, *CreateDocumentPartyRequest) (*CreateDocumentPartyResponse, error)
	CreateDocumentPartiesByBookingID(context.Context, *CreateDocumentPartiesByBookingIDRequest) (*CreateDocumentPartiesByBookingIDResponse, error)
	CreateDocumentPartiesByShippingInstructionID(context.Context, *CreateDocumentPartiesByShippingInstructionIDRequest) (*CreateDocumentPartiesByShippingInstructionIDResponse, error)
	FetchDocumentPartiesByBookingID(context.Context, *FetchDocumentPartiesByBookingIDRequest) (*FetchDocumentPartiesByBookingIDResponse, error)
	FetchDocumentPartiesByByShippingInstructionID(context.Context, *FetchDocumentPartiesByByShippingInstructionIDRequest) (*FetchDocumentPartiesByByShippingInstructionIDResponse, error)
	ResolveDocumentPartiesForShippingInstructionID(context.Context, *ResolveDocumentPartiesForShippingInstructionIDRequest) (*ResolveDocumentPartiesForShippingInstructionIDResponse, error)
	mustEmbedUnimplementedDocumentPartyServiceServer()
}

// UnimplementedDocumentPartyServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDocumentPartyServiceServer struct{}

func (UnimplementedDocumentPartyServiceServer) CreateDocumentParty(context.Context, *CreateDocumentPartyRequest) (*CreateDocumentPartyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDocumentParty not implemented")
}
func (UnimplementedDocumentPartyServiceServer) CreateDocumentPartiesByBookingID(context.Context, *CreateDocumentPartiesByBookingIDRequest) (*CreateDocumentPartiesByBookingIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDocumentPartiesByBookingID not implemented")
}
func (UnimplementedDocumentPartyServiceServer) CreateDocumentPartiesByShippingInstructionID(context.Context, *CreateDocumentPartiesByShippingInstructionIDRequest) (*CreateDocumentPartiesByShippingInstructionIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDocumentPartiesByShippingInstructionID not implemented")
}
func (UnimplementedDocumentPartyServiceServer) FetchDocumentPartiesByBookingID(context.Context, *FetchDocumentPartiesByBookingIDRequest) (*FetchDocumentPartiesByBookingIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchDocumentPartiesByBookingID not implemented")
}
func (UnimplementedDocumentPartyServiceServer) FetchDocumentPartiesByByShippingInstructionID(context.Context, *FetchDocumentPartiesByByShippingInstructionIDRequest) (*FetchDocumentPartiesByByShippingInstructionIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchDocumentPartiesByByShippingInstructionID not implemented")
}
func (UnimplementedDocumentPartyServiceServer) ResolveDocumentPartiesForShippingInstructionID(context.Context, *ResolveDocumentPartiesForShippingInstructionIDRequest) (*ResolveDocumentPartiesForShippingInstructionIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResolveDocumentPartiesForShippingInstructionID not implemented")
}
func (UnimplementedDocumentPartyServiceServer) mustEmbedUnimplementedDocumentPartyServiceServer() {}
func (UnimplementedDocumentPartyServiceServer) testEmbeddedByValue()                              {}

// UnsafeDocumentPartyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DocumentPartyServiceServer will
// result in compilation errors.
type UnsafeDocumentPartyServiceServer interface {
	mustEmbedUnimplementedDocumentPartyServiceServer()
}

func RegisterDocumentPartyServiceServer(s grpc.ServiceRegistrar, srv DocumentPartyServiceServer) {
	// If the following call pancis, it indicates UnimplementedDocumentPartyServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DocumentPartyService_ServiceDesc, srv)
}

func _DocumentPartyService_CreateDocumentParty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDocumentPartyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentPartyServiceServer).CreateDocumentParty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DocumentPartyService_CreateDocumentParty_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentPartyServiceServer).CreateDocumentParty(ctx, req.(*CreateDocumentPartyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DocumentPartyService_CreateDocumentPartiesByBookingID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDocumentPartiesByBookingIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentPartyServiceServer).CreateDocumentPartiesByBookingID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DocumentPartyService_CreateDocumentPartiesByBookingID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentPartyServiceServer).CreateDocumentPartiesByBookingID(ctx, req.(*CreateDocumentPartiesByBookingIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DocumentPartyService_CreateDocumentPartiesByShippingInstructionID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDocumentPartiesByShippingInstructionIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentPartyServiceServer).CreateDocumentPartiesByShippingInstructionID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DocumentPartyService_CreateDocumentPartiesByShippingInstructionID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentPartyServiceServer).CreateDocumentPartiesByShippingInstructionID(ctx, req.(*CreateDocumentPartiesByShippingInstructionIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DocumentPartyService_FetchDocumentPartiesByBookingID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchDocumentPartiesByBookingIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentPartyServiceServer).FetchDocumentPartiesByBookingID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DocumentPartyService_FetchDocumentPartiesByBookingID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentPartyServiceServer).FetchDocumentPartiesByBookingID(ctx, req.(*FetchDocumentPartiesByBookingIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DocumentPartyService_FetchDocumentPartiesByByShippingInstructionID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchDocumentPartiesByByShippingInstructionIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentPartyServiceServer).FetchDocumentPartiesByByShippingInstructionID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DocumentPartyService_FetchDocumentPartiesByByShippingInstructionID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentPartyServiceServer).FetchDocumentPartiesByByShippingInstructionID(ctx, req.(*FetchDocumentPartiesByByShippingInstructionIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DocumentPartyService_ResolveDocumentPartiesForShippingInstructionID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResolveDocumentPartiesForShippingInstructionIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentPartyServiceServer).ResolveDocumentPartiesForShippingInstructionID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DocumentPartyService_ResolveDocumentPartiesForShippingInstructionID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentPartyServiceServer).ResolveDocumentPartiesForShippingInstructionID(ctx, req.(*ResolveDocumentPartiesForShippingInstructionIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DocumentPartyService_ServiceDesc is the grpc.ServiceDesc for DocumentPartyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DocumentPartyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "party.v1.DocumentPartyService",
	HandlerType: (*DocumentPartyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDocumentParty",
			Handler:    _DocumentPartyService_CreateDocumentParty_Handler,
		},
		{
			MethodName: "CreateDocumentPartiesByBookingID",
			Handler:    _DocumentPartyService_CreateDocumentPartiesByBookingID_Handler,
		},
		{
			MethodName: "CreateDocumentPartiesByShippingInstructionID",
			Handler:    _DocumentPartyService_CreateDocumentPartiesByShippingInstructionID_Handler,
		},
		{
			MethodName: "FetchDocumentPartiesByBookingID",
			Handler:    _DocumentPartyService_FetchDocumentPartiesByBookingID_Handler,
		},
		{
			MethodName: "FetchDocumentPartiesByByShippingInstructionID",
			Handler:    _DocumentPartyService_FetchDocumentPartiesByByShippingInstructionID_Handler,
		},
		{
			MethodName: "ResolveDocumentPartiesForShippingInstructionID",
			Handler:    _DocumentPartyService_ResolveDocumentPartiesForShippingInstructionID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "party/v1/document_party.proto",
}