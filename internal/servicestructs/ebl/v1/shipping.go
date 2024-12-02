package v1

import (
	"time"

	commonproto "github.com/cloudfresco/sc-dcsa/internal/protogen/common/v1"
	eblproto "github.com/cloudfresco/sc-dcsa/internal/protogen/ebl/v1"
	commonstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/common/v1"
)

// Shipment - struct Shipment
type Shipment struct {
	*eblproto.ShipmentD
	*ShipmentT
	*commonproto.CrUpdUser
	*commonstruct.CrUpdTime
}

// ShipmentT - struct ShipmentT
type ShipmentT struct {
	ConfirmationDatetime time.Time `protobuf:"bytes,1,opt,name=confirmation_datetime,json=confirmationDatetime,proto3" json:"confirmation_datetime,omitempty"`
	UpdatedDateTime      time.Time `protobuf:"bytes,2,opt,name=updated_date_time,json=updatedDateTime,proto3" json:"updated_date_time,omitempty"`
}

// ShipmentCutoffTime - struct ShipmentCutoffTime
type ShipmentCutoffTime struct {
	*eblproto.ShipmentCutoffTimeD
	*ShipmentCutoffTimeT
}

// ShipmentCutoffTimeT - struct ShipmentCutoffTimeT
type ShipmentCutoffTimeT struct {
	CutOffTime time.Time `protobuf:"bytes,1,opt,name=cut_off_time,json=cutOffTime,proto3" json:"cut_off_time,omitempty"`
}

// Transport - struct Transport
type Transport struct {
	*eblproto.TransportD
	*commonproto.CrUpdUser
	*commonstruct.CrUpdTime
}
