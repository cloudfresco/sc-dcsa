package v2

import (
	"time"

	bkgproto "github.com/cloudfresco/sc-dcsa/internal/protogen/bkg/v2"
	commonproto "github.com/cloudfresco/sc-dcsa/internal/protogen/common/v1"
	commonstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/common/v1"
)

// BkgShipmentSummary - struct BkgShipmentSummary
type BkgShipmentSummary struct {
	*bkgproto.BkgShipmentSummaryD
	*BkgShipmentSummaryT
	*commonproto.CrUpdUser
	*commonstruct.CrUpdTime
}

// BkgShipmentSummaryT - struct BkgShipmentSummaryT
type BkgShipmentSummaryT struct {
	ShipmentCreatedDateTime time.Time `protobuf:"bytes,1,opt,name=shipment_created_date_time,json=shipmentCreatedDateTime,proto3" json:"shipment_created_date_time,omitempty"`
	ShipmentUpdatedDateTime time.Time `protobuf:"bytes,2,opt,name=shipment_updated_date_time,json=shipmentUpdatedDateTime,proto3" json:"shipment_updated_date_time,omitempty"`
}
