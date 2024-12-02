package v2

import (
	"time"

	bkgproto "github.com/cloudfresco/sc-dcsa/internal/protogen/bkg/v2"
	commonproto "github.com/cloudfresco/sc-dcsa/internal/protogen/common/v1"
	commonstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/common/v1"
)

// Booking - struct booking
type Booking struct {
	*bkgproto.BookingD
	*BookingT
	*commonproto.CrUpdUser
	*commonstruct.CrUpdTime
}

// BookingT - struct bookingT
type BookingT struct {
	SubmissionDateTime                        time.Time `protobuf:"bytes,1,opt,name=submission_date_time,json=submissionDateTime,proto3" json:"submission_date_time,omitempty"`
	ExpectedDepartureDate                     time.Time `protobuf:"bytes,2,opt,name=expected_departure_date,json=expectedDepartureDate,proto3" json:"expected_departure_date,omitempty"`
	ExpectedArrivalAtPlaceOfDeliveryStartDate time.Time `protobuf:"bytes,3,opt,name=expected_arrival_at_place_of_delivery_start_date,json=expectedArrivalAtPlaceOfDeliveryStartDate,proto3" json:"expected_arrival_at_place_of_delivery_start_date,omitempty"`
	ExpectedArrivalAtPlaceOfDeliveryEndDate   time.Time `protobuf:"bytes,4,opt,name=expected_arrival_at_place_of_delivery_end_date,json=expectedArrivalAtPlaceOfDeliveryEndDate,proto3" json:"expected_arrival_at_place_of_delivery_end_date,omitempty"`
}

// Commodity - struct Commodity
type Commodity struct {
	*bkgproto.CommodityD
	*CommodityT
	*commonproto.CrUpdUser
	*commonstruct.CrUpdTime
}

// CommodityT - struct CommodityT
type CommodityT struct {
	ExportLicenseIssueDate  time.Time `protobuf:"bytes,1,opt,name=export_license_issue_date,json=exportLicenseIssueDate,proto3" json:"export_license_issue_date,omitempty"`
	ExportLicenseExpiryDate time.Time `protobuf:"bytes,2,opt,name=export_license_expiry_date,json=exportLicenseExpiryDate,proto3" json:"export_license_expiry_date,omitempty"`
}

// RequestedEquipment - struct RequestedEquipment
type RequestedEquipment struct {
	*bkgproto.RequestedEquipmentD
	*commonproto.CrUpdUser
	*commonstruct.CrUpdTime
}

// ShipmentLocation - struct ShipmentLocation
type ShipmentLocation struct {
	*bkgproto.ShipmentLocationD
	*ShipmentLocationT
}

// ShipmentLocationT - struct ShipmentLocationT
type ShipmentLocationT struct {
	EventDateTime time.Time `protobuf:"bytes,1,opt,name=event_date_time,json=eventDateTime,proto3" json:"event_date_time,omitempty"`
}
