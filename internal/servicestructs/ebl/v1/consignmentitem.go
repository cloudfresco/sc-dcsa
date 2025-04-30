package v1

import (
	commonproto "github.com/cloudfresco/sc-dcsa/internal/protogen/common/v1"
	eblproto "github.com/cloudfresco/sc-dcsa/internal/protogen/ebl/v1"
	commonstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/common/v1"
)

// ConsignmentItem - struct ConsignmentItem
type ConsignmentItem struct {
	*eblproto.ConsignmentItemD
	*commonproto.CrUpdUser
	*commonstruct.CrUpdTime
}

type CargoItem struct {
	*eblproto.CargoItemD
	*commonproto.CrUpdUser
	*commonstruct.CrUpdTime
	CargoLineItems []*CargoLineItem `protobuf:"bytes,4,rep,name=cargo_line_items,json=cargoLineItems,proto3" json:"cargo_line_items,omitempty"`
}

type CargoLineItem struct {
	*eblproto.CargoLineItemD
	*commonproto.CrUpdUser
	*commonstruct.CrUpdTime
}
