// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: party/v1/party.proto

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
	PartyService_CreateParty_FullMethodName                = "/party.v1.PartyService/CreateParty"
	PartyService_GetParties_FullMethodName                 = "/party.v1.PartyService/GetParties"
	PartyService_GetParty_FullMethodName                   = "/party.v1.PartyService/GetParty"
	PartyService_GetPartyByPk_FullMethodName               = "/party.v1.PartyService/GetPartyByPk"
	PartyService_UpdateParty_FullMethodName                = "/party.v1.PartyService/UpdateParty"
	PartyService_DeleteParty_FullMethodName                = "/party.v1.PartyService/DeleteParty"
	PartyService_CreatePartyContactDetail_FullMethodName   = "/party.v1.PartyService/CreatePartyContactDetail"
	PartyService_GetPartyContactDetail_FullMethodName      = "/party.v1.PartyService/GetPartyContactDetail"
	PartyService_UpdatePartyContactDetail_FullMethodName   = "/party.v1.PartyService/UpdatePartyContactDetail"
	PartyService_DeletePartyContactDetail_FullMethodName   = "/party.v1.PartyService/DeletePartyContactDetail"
	PartyService_CreateDisplayedAddress_FullMethodName     = "/party.v1.PartyService/CreateDisplayedAddress"
	PartyService_CreateFacility_FullMethodName             = "/party.v1.PartyService/CreateFacility"
	PartyService_CreatePartyFunction_FullMethodName        = "/party.v1.PartyService/CreatePartyFunction"
	PartyService_CreatePartyIdentifyingCode_FullMethodName = "/party.v1.PartyService/CreatePartyIdentifyingCode"
)

// PartyServiceClient is the client API for PartyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// The PartyService service definition.
type PartyServiceClient interface {
	CreateParty(ctx context.Context, in *CreatePartyRequest, opts ...grpc.CallOption) (*CreatePartyResponse, error)
	GetParties(ctx context.Context, in *GetPartiesRequest, opts ...grpc.CallOption) (*GetPartiesResponse, error)
	GetParty(ctx context.Context, in *GetPartyRequest, opts ...grpc.CallOption) (*GetPartyResponse, error)
	GetPartyByPk(ctx context.Context, in *GetPartyByPkRequest, opts ...grpc.CallOption) (*GetPartyByPkResponse, error)
	UpdateParty(ctx context.Context, in *UpdatePartyRequest, opts ...grpc.CallOption) (*UpdatePartyResponse, error)
	DeleteParty(ctx context.Context, in *DeletePartyRequest, opts ...grpc.CallOption) (*DeletePartyResponse, error)
	CreatePartyContactDetail(ctx context.Context, in *CreatePartyContactDetailRequest, opts ...grpc.CallOption) (*CreatePartyContactDetailResponse, error)
	GetPartyContactDetail(ctx context.Context, in *GetPartyContactDetailRequest, opts ...grpc.CallOption) (*GetPartyContactDetailResponse, error)
	UpdatePartyContactDetail(ctx context.Context, in *UpdatePartyContactDetailRequest, opts ...grpc.CallOption) (*UpdatePartyContactDetailResponse, error)
	DeletePartyContactDetail(ctx context.Context, in *DeletePartyContactDetailRequest, opts ...grpc.CallOption) (*DeletePartyContactDetailResponse, error)
	CreateDisplayedAddress(ctx context.Context, in *CreateDisplayedAddressRequest, opts ...grpc.CallOption) (*CreateDisplayedAddressResponse, error)
	CreateFacility(ctx context.Context, in *CreateFacilityRequest, opts ...grpc.CallOption) (*CreateFacilityResponse, error)
	CreatePartyFunction(ctx context.Context, in *CreatePartyFunctionRequest, opts ...grpc.CallOption) (*CreatePartyFunctionResponse, error)
	CreatePartyIdentifyingCode(ctx context.Context, in *CreatePartyIdentifyingCodeRequest, opts ...grpc.CallOption) (*CreatePartyIdentifyingCodeResponse, error)
}

type partyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPartyServiceClient(cc grpc.ClientConnInterface) PartyServiceClient {
	return &partyServiceClient{cc}
}

func (c *partyServiceClient) CreateParty(ctx context.Context, in *CreatePartyRequest, opts ...grpc.CallOption) (*CreatePartyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreatePartyResponse)
	err := c.cc.Invoke(ctx, PartyService_CreateParty_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partyServiceClient) GetParties(ctx context.Context, in *GetPartiesRequest, opts ...grpc.CallOption) (*GetPartiesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPartiesResponse)
	err := c.cc.Invoke(ctx, PartyService_GetParties_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partyServiceClient) GetParty(ctx context.Context, in *GetPartyRequest, opts ...grpc.CallOption) (*GetPartyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPartyResponse)
	err := c.cc.Invoke(ctx, PartyService_GetParty_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partyServiceClient) GetPartyByPk(ctx context.Context, in *GetPartyByPkRequest, opts ...grpc.CallOption) (*GetPartyByPkResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPartyByPkResponse)
	err := c.cc.Invoke(ctx, PartyService_GetPartyByPk_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partyServiceClient) UpdateParty(ctx context.Context, in *UpdatePartyRequest, opts ...grpc.CallOption) (*UpdatePartyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdatePartyResponse)
	err := c.cc.Invoke(ctx, PartyService_UpdateParty_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partyServiceClient) DeleteParty(ctx context.Context, in *DeletePartyRequest, opts ...grpc.CallOption) (*DeletePartyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeletePartyResponse)
	err := c.cc.Invoke(ctx, PartyService_DeleteParty_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partyServiceClient) CreatePartyContactDetail(ctx context.Context, in *CreatePartyContactDetailRequest, opts ...grpc.CallOption) (*CreatePartyContactDetailResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreatePartyContactDetailResponse)
	err := c.cc.Invoke(ctx, PartyService_CreatePartyContactDetail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partyServiceClient) GetPartyContactDetail(ctx context.Context, in *GetPartyContactDetailRequest, opts ...grpc.CallOption) (*GetPartyContactDetailResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPartyContactDetailResponse)
	err := c.cc.Invoke(ctx, PartyService_GetPartyContactDetail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partyServiceClient) UpdatePartyContactDetail(ctx context.Context, in *UpdatePartyContactDetailRequest, opts ...grpc.CallOption) (*UpdatePartyContactDetailResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdatePartyContactDetailResponse)
	err := c.cc.Invoke(ctx, PartyService_UpdatePartyContactDetail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partyServiceClient) DeletePartyContactDetail(ctx context.Context, in *DeletePartyContactDetailRequest, opts ...grpc.CallOption) (*DeletePartyContactDetailResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeletePartyContactDetailResponse)
	err := c.cc.Invoke(ctx, PartyService_DeletePartyContactDetail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partyServiceClient) CreateDisplayedAddress(ctx context.Context, in *CreateDisplayedAddressRequest, opts ...grpc.CallOption) (*CreateDisplayedAddressResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateDisplayedAddressResponse)
	err := c.cc.Invoke(ctx, PartyService_CreateDisplayedAddress_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partyServiceClient) CreateFacility(ctx context.Context, in *CreateFacilityRequest, opts ...grpc.CallOption) (*CreateFacilityResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateFacilityResponse)
	err := c.cc.Invoke(ctx, PartyService_CreateFacility_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partyServiceClient) CreatePartyFunction(ctx context.Context, in *CreatePartyFunctionRequest, opts ...grpc.CallOption) (*CreatePartyFunctionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreatePartyFunctionResponse)
	err := c.cc.Invoke(ctx, PartyService_CreatePartyFunction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partyServiceClient) CreatePartyIdentifyingCode(ctx context.Context, in *CreatePartyIdentifyingCodeRequest, opts ...grpc.CallOption) (*CreatePartyIdentifyingCodeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreatePartyIdentifyingCodeResponse)
	err := c.cc.Invoke(ctx, PartyService_CreatePartyIdentifyingCode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PartyServiceServer is the server API for PartyService service.
// All implementations must embed UnimplementedPartyServiceServer
// for forward compatibility.
//
// The PartyService service definition.
type PartyServiceServer interface {
	CreateParty(context.Context, *CreatePartyRequest) (*CreatePartyResponse, error)
	GetParties(context.Context, *GetPartiesRequest) (*GetPartiesResponse, error)
	GetParty(context.Context, *GetPartyRequest) (*GetPartyResponse, error)
	GetPartyByPk(context.Context, *GetPartyByPkRequest) (*GetPartyByPkResponse, error)
	UpdateParty(context.Context, *UpdatePartyRequest) (*UpdatePartyResponse, error)
	DeleteParty(context.Context, *DeletePartyRequest) (*DeletePartyResponse, error)
	CreatePartyContactDetail(context.Context, *CreatePartyContactDetailRequest) (*CreatePartyContactDetailResponse, error)
	GetPartyContactDetail(context.Context, *GetPartyContactDetailRequest) (*GetPartyContactDetailResponse, error)
	UpdatePartyContactDetail(context.Context, *UpdatePartyContactDetailRequest) (*UpdatePartyContactDetailResponse, error)
	DeletePartyContactDetail(context.Context, *DeletePartyContactDetailRequest) (*DeletePartyContactDetailResponse, error)
	CreateDisplayedAddress(context.Context, *CreateDisplayedAddressRequest) (*CreateDisplayedAddressResponse, error)
	CreateFacility(context.Context, *CreateFacilityRequest) (*CreateFacilityResponse, error)
	CreatePartyFunction(context.Context, *CreatePartyFunctionRequest) (*CreatePartyFunctionResponse, error)
	CreatePartyIdentifyingCode(context.Context, *CreatePartyIdentifyingCodeRequest) (*CreatePartyIdentifyingCodeResponse, error)
	mustEmbedUnimplementedPartyServiceServer()
}

// UnimplementedPartyServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPartyServiceServer struct{}

func (UnimplementedPartyServiceServer) CreateParty(context.Context, *CreatePartyRequest) (*CreatePartyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateParty not implemented")
}
func (UnimplementedPartyServiceServer) GetParties(context.Context, *GetPartiesRequest) (*GetPartiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetParties not implemented")
}
func (UnimplementedPartyServiceServer) GetParty(context.Context, *GetPartyRequest) (*GetPartyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetParty not implemented")
}
func (UnimplementedPartyServiceServer) GetPartyByPk(context.Context, *GetPartyByPkRequest) (*GetPartyByPkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPartyByPk not implemented")
}
func (UnimplementedPartyServiceServer) UpdateParty(context.Context, *UpdatePartyRequest) (*UpdatePartyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParty not implemented")
}
func (UnimplementedPartyServiceServer) DeleteParty(context.Context, *DeletePartyRequest) (*DeletePartyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteParty not implemented")
}
func (UnimplementedPartyServiceServer) CreatePartyContactDetail(context.Context, *CreatePartyContactDetailRequest) (*CreatePartyContactDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePartyContactDetail not implemented")
}
func (UnimplementedPartyServiceServer) GetPartyContactDetail(context.Context, *GetPartyContactDetailRequest) (*GetPartyContactDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPartyContactDetail not implemented")
}
func (UnimplementedPartyServiceServer) UpdatePartyContactDetail(context.Context, *UpdatePartyContactDetailRequest) (*UpdatePartyContactDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePartyContactDetail not implemented")
}
func (UnimplementedPartyServiceServer) DeletePartyContactDetail(context.Context, *DeletePartyContactDetailRequest) (*DeletePartyContactDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePartyContactDetail not implemented")
}
func (UnimplementedPartyServiceServer) CreateDisplayedAddress(context.Context, *CreateDisplayedAddressRequest) (*CreateDisplayedAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDisplayedAddress not implemented")
}
func (UnimplementedPartyServiceServer) CreateFacility(context.Context, *CreateFacilityRequest) (*CreateFacilityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFacility not implemented")
}
func (UnimplementedPartyServiceServer) CreatePartyFunction(context.Context, *CreatePartyFunctionRequest) (*CreatePartyFunctionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePartyFunction not implemented")
}
func (UnimplementedPartyServiceServer) CreatePartyIdentifyingCode(context.Context, *CreatePartyIdentifyingCodeRequest) (*CreatePartyIdentifyingCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePartyIdentifyingCode not implemented")
}
func (UnimplementedPartyServiceServer) mustEmbedUnimplementedPartyServiceServer() {}
func (UnimplementedPartyServiceServer) testEmbeddedByValue()                      {}

// UnsafePartyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PartyServiceServer will
// result in compilation errors.
type UnsafePartyServiceServer interface {
	mustEmbedUnimplementedPartyServiceServer()
}

func RegisterPartyServiceServer(s grpc.ServiceRegistrar, srv PartyServiceServer) {
	// If the following call pancis, it indicates UnimplementedPartyServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PartyService_ServiceDesc, srv)
}

func _PartyService_CreateParty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePartyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartyServiceServer).CreateParty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PartyService_CreateParty_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartyServiceServer).CreateParty(ctx, req.(*CreatePartyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartyService_GetParties_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPartiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartyServiceServer).GetParties(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PartyService_GetParties_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartyServiceServer).GetParties(ctx, req.(*GetPartiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartyService_GetParty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPartyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartyServiceServer).GetParty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PartyService_GetParty_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartyServiceServer).GetParty(ctx, req.(*GetPartyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartyService_GetPartyByPk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPartyByPkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartyServiceServer).GetPartyByPk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PartyService_GetPartyByPk_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartyServiceServer).GetPartyByPk(ctx, req.(*GetPartyByPkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartyService_UpdateParty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePartyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartyServiceServer).UpdateParty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PartyService_UpdateParty_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartyServiceServer).UpdateParty(ctx, req.(*UpdatePartyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartyService_DeleteParty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePartyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartyServiceServer).DeleteParty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PartyService_DeleteParty_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartyServiceServer).DeleteParty(ctx, req.(*DeletePartyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartyService_CreatePartyContactDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePartyContactDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartyServiceServer).CreatePartyContactDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PartyService_CreatePartyContactDetail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartyServiceServer).CreatePartyContactDetail(ctx, req.(*CreatePartyContactDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartyService_GetPartyContactDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPartyContactDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartyServiceServer).GetPartyContactDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PartyService_GetPartyContactDetail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartyServiceServer).GetPartyContactDetail(ctx, req.(*GetPartyContactDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartyService_UpdatePartyContactDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePartyContactDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartyServiceServer).UpdatePartyContactDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PartyService_UpdatePartyContactDetail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartyServiceServer).UpdatePartyContactDetail(ctx, req.(*UpdatePartyContactDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartyService_DeletePartyContactDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePartyContactDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartyServiceServer).DeletePartyContactDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PartyService_DeletePartyContactDetail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartyServiceServer).DeletePartyContactDetail(ctx, req.(*DeletePartyContactDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartyService_CreateDisplayedAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDisplayedAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartyServiceServer).CreateDisplayedAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PartyService_CreateDisplayedAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartyServiceServer).CreateDisplayedAddress(ctx, req.(*CreateDisplayedAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartyService_CreateFacility_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFacilityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartyServiceServer).CreateFacility(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PartyService_CreateFacility_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartyServiceServer).CreateFacility(ctx, req.(*CreateFacilityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartyService_CreatePartyFunction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePartyFunctionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartyServiceServer).CreatePartyFunction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PartyService_CreatePartyFunction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartyServiceServer).CreatePartyFunction(ctx, req.(*CreatePartyFunctionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartyService_CreatePartyIdentifyingCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePartyIdentifyingCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartyServiceServer).CreatePartyIdentifyingCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PartyService_CreatePartyIdentifyingCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartyServiceServer).CreatePartyIdentifyingCode(ctx, req.(*CreatePartyIdentifyingCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PartyService_ServiceDesc is the grpc.ServiceDesc for PartyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PartyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "party.v1.PartyService",
	HandlerType: (*PartyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateParty",
			Handler:    _PartyService_CreateParty_Handler,
		},
		{
			MethodName: "GetParties",
			Handler:    _PartyService_GetParties_Handler,
		},
		{
			MethodName: "GetParty",
			Handler:    _PartyService_GetParty_Handler,
		},
		{
			MethodName: "GetPartyByPk",
			Handler:    _PartyService_GetPartyByPk_Handler,
		},
		{
			MethodName: "UpdateParty",
			Handler:    _PartyService_UpdateParty_Handler,
		},
		{
			MethodName: "DeleteParty",
			Handler:    _PartyService_DeleteParty_Handler,
		},
		{
			MethodName: "CreatePartyContactDetail",
			Handler:    _PartyService_CreatePartyContactDetail_Handler,
		},
		{
			MethodName: "GetPartyContactDetail",
			Handler:    _PartyService_GetPartyContactDetail_Handler,
		},
		{
			MethodName: "UpdatePartyContactDetail",
			Handler:    _PartyService_UpdatePartyContactDetail_Handler,
		},
		{
			MethodName: "DeletePartyContactDetail",
			Handler:    _PartyService_DeletePartyContactDetail_Handler,
		},
		{
			MethodName: "CreateDisplayedAddress",
			Handler:    _PartyService_CreateDisplayedAddress_Handler,
		},
		{
			MethodName: "CreateFacility",
			Handler:    _PartyService_CreateFacility_Handler,
		},
		{
			MethodName: "CreatePartyFunction",
			Handler:    _PartyService_CreatePartyFunction_Handler,
		},
		{
			MethodName: "CreatePartyIdentifyingCode",
			Handler:    _PartyService_CreatePartyIdentifyingCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "party/v1/party.proto",
}