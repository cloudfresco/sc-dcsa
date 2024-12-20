// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: party/v1/location.proto

package v1

import (
	v1 "github.com/cloudfresco/sc-dcsa/internal/protogen/common/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateLocationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LocationName     string `protobuf:"bytes,1,opt,name=location_name,json=locationName,proto3" json:"location_name,omitempty"`
	Latitude         string `protobuf:"bytes,2,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude        string `protobuf:"bytes,3,opt,name=longitude,proto3" json:"longitude,omitempty"`
	FacilitySmdgCode string `protobuf:"bytes,4,opt,name=facility_smdg_code,json=facilitySmdgCode,proto3" json:"facility_smdg_code,omitempty"`
	UnLocationCode   string `protobuf:"bytes,5,opt,name=un_location_code,json=unLocationCode,proto3" json:"un_location_code,omitempty"`
	AddressId        uint32 `protobuf:"varint,6,opt,name=address_id,json=addressId,proto3" json:"address_id,omitempty"`
	FacilityId       uint32 `protobuf:"varint,7,opt,name=facility_id,json=facilityId,proto3" json:"facility_id,omitempty"`
	UserId           string `protobuf:"bytes,8,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserEmail        string `protobuf:"bytes,9,opt,name=user_email,json=userEmail,proto3" json:"user_email,omitempty"`
	RequestId        string `protobuf:"bytes,10,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
}

func (x *CreateLocationRequest) Reset() {
	*x = CreateLocationRequest{}
	mi := &file_party_v1_location_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateLocationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLocationRequest) ProtoMessage() {}

func (x *CreateLocationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_party_v1_location_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLocationRequest.ProtoReflect.Descriptor instead.
func (*CreateLocationRequest) Descriptor() ([]byte, []int) {
	return file_party_v1_location_proto_rawDescGZIP(), []int{0}
}

func (x *CreateLocationRequest) GetLocationName() string {
	if x != nil {
		return x.LocationName
	}
	return ""
}

func (x *CreateLocationRequest) GetLatitude() string {
	if x != nil {
		return x.Latitude
	}
	return ""
}

func (x *CreateLocationRequest) GetLongitude() string {
	if x != nil {
		return x.Longitude
	}
	return ""
}

func (x *CreateLocationRequest) GetFacilitySmdgCode() string {
	if x != nil {
		return x.FacilitySmdgCode
	}
	return ""
}

func (x *CreateLocationRequest) GetUnLocationCode() string {
	if x != nil {
		return x.UnLocationCode
	}
	return ""
}

func (x *CreateLocationRequest) GetAddressId() uint32 {
	if x != nil {
		return x.AddressId
	}
	return 0
}

func (x *CreateLocationRequest) GetFacilityId() uint32 {
	if x != nil {
		return x.FacilityId
	}
	return 0
}

func (x *CreateLocationRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateLocationRequest) GetUserEmail() string {
	if x != nil {
		return x.UserEmail
	}
	return ""
}

func (x *CreateLocationRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

type CreateLocationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Location *Location `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *CreateLocationResponse) Reset() {
	*x = CreateLocationResponse{}
	mi := &file_party_v1_location_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateLocationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLocationResponse) ProtoMessage() {}

func (x *CreateLocationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_party_v1_location_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLocationResponse.ProtoReflect.Descriptor instead.
func (*CreateLocationResponse) Descriptor() ([]byte, []int) {
	return file_party_v1_location_proto_rawDescGZIP(), []int{1}
}

func (x *CreateLocationResponse) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

type LoadLocationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Locations  []*Location `protobuf:"bytes,1,rep,name=locations,proto3" json:"locations,omitempty"`
	NextCursor string      `protobuf:"bytes,2,opt,name=next_cursor,json=nextCursor,proto3" json:"next_cursor,omitempty"`
}

func (x *LoadLocationsResponse) Reset() {
	*x = LoadLocationsResponse{}
	mi := &file_party_v1_location_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoadLocationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadLocationsResponse) ProtoMessage() {}

func (x *LoadLocationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_party_v1_location_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadLocationsResponse.ProtoReflect.Descriptor instead.
func (*LoadLocationsResponse) Descriptor() ([]byte, []int) {
	return file_party_v1_location_proto_rawDescGZIP(), []int{2}
}

func (x *LoadLocationsResponse) GetLocations() []*Location {
	if x != nil {
		return x.Locations
	}
	return nil
}

func (x *LoadLocationsResponse) GetNextCursor() string {
	if x != nil {
		return x.NextCursor
	}
	return ""
}

type LoadLocationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit      string `protobuf:"bytes,1,opt,name=limit,proto3" json:"limit,omitempty"`
	NextCursor string `protobuf:"bytes,2,opt,name=next_cursor,json=nextCursor,proto3" json:"next_cursor,omitempty"`
	UserEmail  string `protobuf:"bytes,3,opt,name=user_email,json=userEmail,proto3" json:"user_email,omitempty"`
	RequestId  string `protobuf:"bytes,4,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
}

func (x *LoadLocationsRequest) Reset() {
	*x = LoadLocationsRequest{}
	mi := &file_party_v1_location_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoadLocationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadLocationsRequest) ProtoMessage() {}

func (x *LoadLocationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_party_v1_location_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadLocationsRequest.ProtoReflect.Descriptor instead.
func (*LoadLocationsRequest) Descriptor() ([]byte, []int) {
	return file_party_v1_location_proto_rawDescGZIP(), []int{3}
}

func (x *LoadLocationsRequest) GetLimit() string {
	if x != nil {
		return x.Limit
	}
	return ""
}

func (x *LoadLocationsRequest) GetNextCursor() string {
	if x != nil {
		return x.NextCursor
	}
	return ""
}

func (x *LoadLocationsRequest) GetUserEmail() string {
	if x != nil {
		return x.UserEmail
	}
	return ""
}

func (x *LoadLocationsRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

type FetchLocationByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GetByIdRequest *v1.GetByIdRequest `protobuf:"bytes,1,opt,name=get_by_id_request,json=getByIdRequest,proto3" json:"get_by_id_request,omitempty"`
}

func (x *FetchLocationByIDRequest) Reset() {
	*x = FetchLocationByIDRequest{}
	mi := &file_party_v1_location_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FetchLocationByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchLocationByIDRequest) ProtoMessage() {}

func (x *FetchLocationByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_party_v1_location_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchLocationByIDRequest.ProtoReflect.Descriptor instead.
func (*FetchLocationByIDRequest) Descriptor() ([]byte, []int) {
	return file_party_v1_location_proto_rawDescGZIP(), []int{4}
}

func (x *FetchLocationByIDRequest) GetGetByIdRequest() *v1.GetByIdRequest {
	if x != nil {
		return x.GetByIdRequest
	}
	return nil
}

type FetchLocationByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Location *Location `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *FetchLocationByIDResponse) Reset() {
	*x = FetchLocationByIDResponse{}
	mi := &file_party_v1_location_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FetchLocationByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchLocationByIDResponse) ProtoMessage() {}

func (x *FetchLocationByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_party_v1_location_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchLocationByIDResponse.ProtoReflect.Descriptor instead.
func (*FetchLocationByIDResponse) Descriptor() ([]byte, []int) {
	return file_party_v1_location_proto_rawDescGZIP(), []int{5}
}

func (x *FetchLocationByIDResponse) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uuid4            []byte `protobuf:"bytes,2,opt,name=uuid4,proto3" json:"uuid4,omitempty"`
	IdS              string `protobuf:"bytes,3,opt,name=id_s,json=idS,proto3" json:"id_s,omitempty"`
	LocationName     string `protobuf:"bytes,4,opt,name=location_name,json=locationName,proto3" json:"location_name,omitempty"`
	Latitude         string `protobuf:"bytes,5,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude        string `protobuf:"bytes,6,opt,name=longitude,proto3" json:"longitude,omitempty"`
	FacilitySmdgCode string `protobuf:"bytes,7,opt,name=facility_smdg_code,json=facilitySmdgCode,proto3" json:"facility_smdg_code,omitempty"`
	UnLocationCode   string `protobuf:"bytes,8,opt,name=un_location_code,json=unLocationCode,proto3" json:"un_location_code,omitempty"`
	AddressId        uint32 `protobuf:"varint,9,opt,name=address_id,json=addressId,proto3" json:"address_id,omitempty"`
	FacilityId       uint32 `protobuf:"varint,10,opt,name=facility_id,json=facilityId,proto3" json:"facility_id,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	mi := &file_party_v1_location_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_party_v1_location_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_party_v1_location_proto_rawDescGZIP(), []int{6}
}

func (x *Location) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Location) GetUuid4() []byte {
	if x != nil {
		return x.Uuid4
	}
	return nil
}

func (x *Location) GetIdS() string {
	if x != nil {
		return x.IdS
	}
	return ""
}

func (x *Location) GetLocationName() string {
	if x != nil {
		return x.LocationName
	}
	return ""
}

func (x *Location) GetLatitude() string {
	if x != nil {
		return x.Latitude
	}
	return ""
}

func (x *Location) GetLongitude() string {
	if x != nil {
		return x.Longitude
	}
	return ""
}

func (x *Location) GetFacilitySmdgCode() string {
	if x != nil {
		return x.FacilitySmdgCode
	}
	return ""
}

func (x *Location) GetUnLocationCode() string {
	if x != nil {
		return x.UnLocationCode
	}
	return ""
}

func (x *Location) GetAddressId() uint32 {
	if x != nil {
		return x.AddressId
	}
	return 0
}

func (x *Location) GetFacilityId() uint32 {
	if x != nil {
		return x.FacilityId
	}
	return 0
}

var File_party_v1_location_proto protoreflect.FileDescriptor

var file_party_v1_location_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x61, 0x72, 0x74, 0x79,
	0x2e, 0x76, 0x31, 0x1a, 0x16, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe5, 0x02, 0x0a, 0x15,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61,
	0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61,
	0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x66, 0x61, 0x63, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x5f, 0x73, 0x6d, 0x64, 0x67, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x10, 0x66, 0x61, 0x63, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x6d, 0x64, 0x67, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x75, 0x6e, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x75, 0x6e,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x66,
	0x61, 0x63, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0a, 0x66, 0x61, 0x63, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x49, 0x64, 0x22, 0x48, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a,
	0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x6a, 0x0a,
	0x15, 0x4c, 0x6f, 0x61, 0x64, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x61, 0x72, 0x74,
	0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x65, 0x78, 0x74,
	0x5f, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6e,
	0x65, 0x78, 0x74, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x22, 0x8b, 0x01, 0x0a, 0x14, 0x4c, 0x6f,
	0x61, 0x64, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x65, 0x78, 0x74,
	0x5f, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6e,
	0x65, 0x78, 0x74, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75,
	0x73, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x22, 0x60, 0x0a, 0x18, 0x46, 0x65, 0x74, 0x63, 0x68,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x44, 0x0a, 0x11, 0x67, 0x65, 0x74, 0x5f, 0x62, 0x79, 0x5f, 0x69, 0x64,
	0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79,
	0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0e, 0x67, 0x65, 0x74, 0x42, 0x79,
	0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4b, 0x0a, 0x19, 0x46, 0x65, 0x74,
	0x63, 0x68, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xba, 0x02, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x75, 0x69, 0x64, 0x34, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x05, 0x75, 0x75, 0x69, 0x64, 0x34, 0x12, 0x11, 0x0a, 0x04, 0x69, 0x64, 0x5f,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x53, 0x12, 0x23, 0x0a, 0x0d,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x66,
	0x61, 0x63, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x73, 0x6d, 0x64, 0x67, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x66, 0x61, 0x63, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x53, 0x6d, 0x64, 0x67, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x75, 0x6e, 0x5f,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x75, 0x6e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x69,
	0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x61, 0x63, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x69,
	0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x66, 0x61, 0x63, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x49, 0x64, 0x32, 0x96, 0x02, 0x0a, 0x0f, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x53, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x2e, 0x70, 0x61, 0x72, 0x74,
	0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x70, 0x61, 0x72,
	0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x0d,
	0x4c, 0x6f, 0x61, 0x64, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1e, 0x2e,
	0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x61, 0x64, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e,
	0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x61, 0x64, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5c,
	0x0a, 0x11, 0x46, 0x65, 0x74, 0x63, 0x68, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x79, 0x49, 0x44, 0x12, 0x22, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x46,
	0x65, 0x74, 0x63, 0x68, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x49, 0x44,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3b, 0x5a, 0x39,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x66, 0x72, 0x65, 0x73, 0x63, 0x6f, 0x2f, 0x73, 0x63, 0x2d, 0x64, 0x63, 0x73, 0x61, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e,
	0x2f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_party_v1_location_proto_rawDescOnce sync.Once
	file_party_v1_location_proto_rawDescData = file_party_v1_location_proto_rawDesc
)

func file_party_v1_location_proto_rawDescGZIP() []byte {
	file_party_v1_location_proto_rawDescOnce.Do(func() {
		file_party_v1_location_proto_rawDescData = protoimpl.X.CompressGZIP(file_party_v1_location_proto_rawDescData)
	})
	return file_party_v1_location_proto_rawDescData
}

var file_party_v1_location_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_party_v1_location_proto_goTypes = []any{
	(*CreateLocationRequest)(nil),     // 0: party.v1.CreateLocationRequest
	(*CreateLocationResponse)(nil),    // 1: party.v1.CreateLocationResponse
	(*LoadLocationsResponse)(nil),     // 2: party.v1.LoadLocationsResponse
	(*LoadLocationsRequest)(nil),      // 3: party.v1.LoadLocationsRequest
	(*FetchLocationByIDRequest)(nil),  // 4: party.v1.FetchLocationByIDRequest
	(*FetchLocationByIDResponse)(nil), // 5: party.v1.FetchLocationByIDResponse
	(*Location)(nil),                  // 6: party.v1.Location
	(*v1.GetByIdRequest)(nil),         // 7: common.v1.GetByIdRequest
}
var file_party_v1_location_proto_depIdxs = []int32{
	6, // 0: party.v1.CreateLocationResponse.location:type_name -> party.v1.Location
	6, // 1: party.v1.LoadLocationsResponse.locations:type_name -> party.v1.Location
	7, // 2: party.v1.FetchLocationByIDRequest.get_by_id_request:type_name -> common.v1.GetByIdRequest
	6, // 3: party.v1.FetchLocationByIDResponse.location:type_name -> party.v1.Location
	0, // 4: party.v1.LocationService.CreateLocation:input_type -> party.v1.CreateLocationRequest
	3, // 5: party.v1.LocationService.LoadLocations:input_type -> party.v1.LoadLocationsRequest
	4, // 6: party.v1.LocationService.FetchLocationByID:input_type -> party.v1.FetchLocationByIDRequest
	1, // 7: party.v1.LocationService.CreateLocation:output_type -> party.v1.CreateLocationResponse
	2, // 8: party.v1.LocationService.LoadLocations:output_type -> party.v1.LoadLocationsResponse
	5, // 9: party.v1.LocationService.FetchLocationByID:output_type -> party.v1.FetchLocationByIDResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_party_v1_location_proto_init() }
func file_party_v1_location_proto_init() {
	if File_party_v1_location_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_party_v1_location_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_party_v1_location_proto_goTypes,
		DependencyIndexes: file_party_v1_location_proto_depIdxs,
		MessageInfos:      file_party_v1_location_proto_msgTypes,
	}.Build()
	File_party_v1_location_proto = out.File
	file_party_v1_location_proto_rawDesc = nil
	file_party_v1_location_proto_goTypes = nil
	file_party_v1_location_proto_depIdxs = nil
}
