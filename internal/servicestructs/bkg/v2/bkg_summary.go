package v2

import (
	"time"

	bkgproto "github.com/cloudfresco/sc-dcsa/internal/protogen/bkg/v2"
	commonproto "github.com/cloudfresco/sc-dcsa/internal/protogen/common/v1"
	commonstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/common/v1"
)

// BookingSummary - struct BookingSummary
type BookingSummary struct {
	*bkgproto.BookingSummaryD
	*BookingSummaryT
	*commonproto.CrUpdUser
	*commonstruct.CrUpdTime
}

// BookingSummaryT - struct BookingSummaryT
type BookingSummaryT struct {
	BookingRequestCreatedDateTime             time.Time `protobuf:"bytes,1,opt,name=booking_request_created_date_time,json=bookingRequestCreatedDateTime,proto3" json:"booking_request_created_date_time,omitempty"`
	BookingRequestUpdatedDateTime             time.Time `protobuf:"bytes,2,opt,name=booking_request_updated_date_time,json=bookingRequestUpdatedDateTime,proto3" json:"booking_request_updated_date_time,omitempty"`
	SubmissionDateTime                        time.Time `protobuf:"bytes,3,opt,name=submission_date_time,json=submissionDateTime,proto3" json:"submission_date_time,omitempty"`
	ExpectedDepartureDate                     time.Time `protobuf:"bytes,4,opt,name=expected_departure_date,json=expectedDepartureDate,proto3" json:"expected_departure_date,omitempty"`
	ExpectedArrivalAtPlaceOfDeliveryStartDate time.Time `protobuf:"bytes,5,opt,name=expected_arrival_at_place_of_delivery_start_date,json=expectedArrivalAtPlaceOfDeliveryStartDate,proto3" json:"expected_arrival_at_place_of_delivery_start_date,omitempty"`
	ExpectedArrivalAtPlaceOfDeliveryEndDate   time.Time `protobuf:"bytes,6,opt,name=expected_arrival_at_place_of_delivery_end_date,json=expectedArrivalAtPlaceOfDeliveryEndDate,proto3" json:"expected_arrival_at_place_of_delivery_end_date,omitempty"`
}
