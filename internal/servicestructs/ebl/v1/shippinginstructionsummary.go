package v1

import (
	"time"

	commonproto "github.com/cloudfresco/sc-dcsa/internal/protogen/common/v1"
	eblproto "github.com/cloudfresco/sc-dcsa/internal/protogen/ebl/v1"
	commonstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/common/v1"
)

// ShippingInstructionSummary - struct ShippingInstructionSummary
type ShippingInstructionSummary struct {
	*eblproto.ShippingInstructionSummaryD
	*ShippingInstructionSummaryT
	*commonproto.CrUpdUser
	*commonstruct.CrUpdTime
}

// ShippingInstructionSummaryT - struct ShippingInstructionSummaryT
type ShippingInstructionSummaryT struct {
	ShippingInstructionCreatedDateTime time.Time `protobuf:"bytes,1,opt,name=shipping_instruction_created_date_time,json=shippingInstructionCreatedDateTime,proto3" json:"shipping_instruction_created_date_time,omitempty"`
	ShippingInstructionUpdatedDateTime time.Time `protobuf:"bytes,2,opt,name=shipping_instruction_updated_date_time,json=shippingInstructionUpdatedDateTime,proto3" json:"shipping_instruction_updated_date_time,omitempty"`
}
