package v1

import (
	"time"

	eblproto "github.com/cloudfresco/sc-dcsa/internal/protogen/ebl/v1"
)

// ShippingInstruction - struct ShippingInstruction
type ShippingInstruction struct {
	*eblproto.ShippingInstructionD
	*ShippingInstructionT
}

// ShippingInstructionT - struct ShippingInstructionT
type ShippingInstructionT struct {
	CreatedDateTime time.Time `protobuf:"bytes,1,opt,name=created_date_time,json=createdDateTime,proto3" json:"created_date_time,omitempty"`
	UpdatedDateTime time.Time `protobuf:"bytes,2,opt,name=updated_date_time,json=updatedDateTime,proto3" json:"updated_date_time,omitempty"`
}
