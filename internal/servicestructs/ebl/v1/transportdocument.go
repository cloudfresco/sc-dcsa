package v1

import (
	"time"

	commonproto "github.com/cloudfresco/sc-dcsa/internal/protogen/common/v1"
	eblproto "github.com/cloudfresco/sc-dcsa/internal/protogen/ebl/v1"
	commonstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/common/v1"
)

// TransportDocument - struct TransportDocument
type TransportDocument struct {
	*eblproto.TransportDocumentD
	*TransportDocumentT
}

// TransportDocumentT - struct TransportDocumentT
type TransportDocumentT struct {
	IssueDate               time.Time `protobuf:"bytes,1,opt,name=issue_date,json=issueDate,proto3" json:"issue_date,omitempty"`
	ShippedOnboardDate      time.Time `protobuf:"bytes,2,opt,name=shipped_onboard_date,json=shippedOnboardDate,proto3" json:"shipped_onboard_date,omitempty"`
	ReceivedForShipmentDate time.Time `protobuf:"bytes,3,opt,name=received_for_shipment_date,json=receivedForShipmentDate,proto3" json:"received_for_shipment_date,omitempty"`
	CreatedDateTime         time.Time `protobuf:"bytes,4,opt,name=created_date_time,json=createdDateTime,proto3" json:"created_date_time,omitempty"`
	UpdatedDateTime         time.Time `protobuf:"bytes,5,opt,name=updated_date_time,json=updatedDateTime,proto3" json:"updated_date_time,omitempty"`
}

// TransportDocumentSummary - struct TransportDocumentSummary
type TransportDocumentSummary struct {
	*eblproto.TransportDocumentSummaryD
	*TransportDocumentSummaryT
	*commonproto.CrUpdUser
	*commonstruct.CrUpdTime
}

// TransportDocumentSummaryT - struct TransportDocumentSummaryT
type TransportDocumentSummaryT struct {
	TransportDocumentCreatedDateTime time.Time `protobuf:"bytes,1,opt,name=transport_document_created_date_time,json=transportDocumentCreatedDateTime,proto3" json:"transport_document_created_date_time,omitempty"`
	TransportDocumentUpdatedDateTime time.Time `protobuf:"bytes,2,opt,name=transport_document_updated_date_time,json=transportDocumentUpdatedDateTime,proto3" json:"transport_document_updated_date_time,omitempty"`
	IssueDate                        time.Time `protobuf:"bytes,3,opt,name=issue_date,json=issueDate,proto3" json:"issue_date,omitempty"`
	ShippedOnboardDate               time.Time `protobuf:"bytes,4,opt,name=shipped_onboard_date,json=shippedOnboardDate,proto3" json:"shipped_onboard_date,omitempty"`
	ReceivedForShipmentDate          time.Time `protobuf:"bytes,5,opt,name=received_for_shipment_date,json=receivedForShipmentDate,proto3" json:"received_for_shipment_date,omitempty"`
}
