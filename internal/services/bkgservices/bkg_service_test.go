// https://github.com/dcsaorg/DCSA-Edocumentation/blob/master/edocumentation-service/src/test/java/org/dcsa/edocumentation/service/BookingServiceTest.java
package bkgservices

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	bkgproto "github.com/cloudfresco/sc-dcsa/internal/protogen/bkg/v2"
	commonproto "github.com/cloudfresco/sc-dcsa/internal/protogen/common/v1"
	"github.com/cloudfresco/sc-dcsa/test"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestBkgService_GetBookings(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	bkgService := NewBkgService(log, dbService, redisService, userServiceClient, currencyService)

	bkg, err := GetBooking(uint32(16), []byte{43, 149, 187, 185, 210, 212, 68, 22, 172, 254, 98, 7, 225, 129, 181, 244}, "2b95bbb9-d2d4-4416-acfe-6207e181b5f4", "ABC123123123", "RECE", "CY", "CFS", "FCL", "LCL", "SERVICE_CONTRACT_REFERENCE_03", "PRE", true, true, "EXPORT_DECLARATION_REFERENCE_03", false, "IMPORT_LICENSE_REFERENCE_03", "2020-03-10T00:00:00Z", false, true, "", "2020-03-10T00:00:00Z", "2020-03-12T00:00:00Z", "2020-03-13T00:00:00Z", "SWB", "TRANSPORT_DOC_REF_03", "BOOKING_CHA_REF_03", "FCA", "EI", false, "", "", "CARRIER_VOYAGE_NUMBER_03", "", uint32(1), "EUR", int64(1212), "12.12", uint32(0), uint32(0), "", "2020-03-10T00:00:00Z", "2021-12-16T00:00:00Z", "auth0|66fd06d0bfea78a82bb42459", "auth0|66fd06d0bfea78a82bb42459")
	if err != nil {
		t.Error(err)
		return
	}
	bkg1, err := GetBooking(uint32(15), []byte{11, 187, 52, 122, 129, 59, 68, 141, 146, 110, 220, 104, 180, 134, 54, 147}, "0bbb347a-813b-448d-926e-dc68b4863693", "BR1239719872", "PENU", "CY", "CFS", "FCL", "LCL", "SERVICE_CONTRACT_REFERENCE_02", "PRE", true, true, "EXPORT_DECLARATION_REFERENCE_02", false, "IMPORT_LICENSE_REFERENCE_02", "2020-04-15T00:00:00Z", false, true, "", "2020-04-15T00:00:00Z", "2020-04-16T00:00:00Z", "2020-04-17T00:00:00Z", "SWB", "TRANSPORT_DOC_REF_02", "BOOKING_CHA_REF_02", "FCA", "EI", false, "", "", "CARRIER_VOYAGE_NUMBER_02", "", uint32(1), "EUR", int64(1212), "12.12", uint32(0), uint32(0), "", "2020-04-15T00:00:00Z", "2021-01-10T00:00:00Z", "auth0|66fd06d0bfea78a82bb42459", "auth0|66fd06d0bfea78a82bb42459")
	if err != nil {
		t.Error(err)
		return
	}
	bkgs := []*bkgproto.Booking{}
	bkgs = append(bkgs, bkg, bkg1)

	form := bkgproto.GetBookingsRequest{}
	form.Limit = "2"
	form.NextCursor = ""
	form.UserEmail = "sprov300@gmail.com"
	form.RequestId = "bks1m1g91jau4nkks2f0"

	nextc := "MTQ="
	bookingResponse := bkgproto.GetBookingsResponse{Bookings: bkgs, NextCursor: nextc}

	type args struct {
		ctx context.Context
		in  *bkgproto.GetBookingsRequest
	}
	tests := []struct {
		bs      *BkgService
		args    args
		want    *bkgproto.GetBookingsResponse
		wantErr bool
	}{
		{
			bs: bkgService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &bookingResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		bookingResult, err := tt.bs.GetBookings(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("BkgService.GetBookings() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(bookingResult, tt.want) {
			t.Errorf("BkgService.GetBookings() = %v, want %v", bookingResult, tt.want)
		}
		assert.NotNil(t, bookingResult)
		booking := bookingResult.Bookings[0]
		assert.Equal(t, booking.BookingD.CarrierBookingRequestReference, "ABC123123123", "they should be equal")
		assert.Equal(t, booking.BookingD.DocumentStatus, "RECE", "they should be equal")
		assert.Equal(t, booking.BookingD.TransportDocumentReference, "TRANSPORT_DOC_REF_03", "they should be equal")
		assert.True(t, booking.BookingD.IsPartialLoadAllowed, "Its true")
	}
}

func TestBkgService_GetBooking(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}
	ctx := LoginUser()

	bkgService := NewBkgService(log, dbService, redisService, userServiceClient, currencyService)

	bkg, err := GetBooking(uint32(16), []byte{43, 149, 187, 185, 210, 212, 68, 22, 172, 254, 98, 7, 225, 129, 181, 244}, "2b95bbb9-d2d4-4416-acfe-6207e181b5f4", "ABC123123123", "RECE", "CY", "CFS", "FCL", "LCL", "SERVICE_CONTRACT_REFERENCE_03", "PRE", true, true, "EXPORT_DECLARATION_REFERENCE_03", false, "IMPORT_LICENSE_REFERENCE_03", "2020-03-10T00:00:00Z", false, true, "", "2020-03-10T00:00:00Z", "2020-03-12T00:00:00Z", "2020-03-13T00:00:00Z", "SWB", "TRANSPORT_DOC_REF_03", "BOOKING_CHA_REF_03", "FCA", "EI", false, "", "", "CARRIER_VOYAGE_NUMBER_03", "", uint32(1), "EUR", int64(1212), "12.12", uint32(0), uint32(0), "", "2020-03-10T00:00:00Z", "2021-12-16T00:00:00Z", "auth0|66fd06d0bfea78a82bb42459", "auth0|66fd06d0bfea78a82bb42459")
	if err != nil {
		t.Error(err)
		return
	}
	bookingResponse := bkgproto.GetBookingResponse{}
	bookingResponse.Booking = bkg

	gform := commonproto.GetRequest{}
	gform.Id = "2b95bbb9-d2d4-4416-acfe-6207e181b5f4"
	gform.UserEmail = "sprov300@gmail.com"
	gform.RequestId = "bks1m1g91jau4nkks2f0"

	form := bkgproto.GetBookingRequest{}
	form.GetRequest = &gform

	type args struct {
		ctx context.Context
		in  *bkgproto.GetBookingRequest
	}

	tests := []struct {
		bs      *BkgService
		args    args
		want    *bkgproto.GetBookingResponse
		wantErr bool
	}{
		{
			bs: bkgService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &bookingResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {

		bkgResponse, err := tt.bs.GetBooking(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("BkgService.GetBooking() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(bkgResponse, tt.want) {
			t.Errorf("BkgService.GetBooking() = %v, want %v", bkgResponse, tt.want)
		}
		bookingResult := bkgResponse.Booking
		assert.NotNil(t, bookingResult)
		assert.Equal(t, bookingResult.BookingD.CarrierBookingRequestReference, "ABC123123123", "they should be equal")
		assert.Equal(t, bookingResult.BookingD.ExportVoyageNumber, "CARRIER_VOYAGE_NUMBER_03", "they should be equal")
		assert.Equal(t, bookingResult.BookingD.DocumentStatus, "RECE", "they should be equal")
	}
}

func TestBkgService_GetBookingByPk(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}
	ctx := LoginUser()

	bkgService := NewBkgService(log, dbService, redisService, userServiceClient, currencyService)

	bkg, err := GetBooking(uint32(16), []byte{43, 149, 187, 185, 210, 212, 68, 22, 172, 254, 98, 7, 225, 129, 181, 244}, "2b95bbb9-d2d4-4416-acfe-6207e181b5f4", "ABC123123123", "RECE", "CY", "CFS", "FCL", "LCL", "SERVICE_CONTRACT_REFERENCE_03", "PRE", true, true, "EXPORT_DECLARATION_REFERENCE_03", false, "IMPORT_LICENSE_REFERENCE_03", "2020-03-10T00:00:00Z", false, true, "", "2020-03-10T00:00:00Z", "2020-03-12T00:00:00Z", "2020-03-13T00:00:00Z", "SWB", "TRANSPORT_DOC_REF_03", "BOOKING_CHA_REF_03", "FCA", "EI", false, "", "", "CARRIER_VOYAGE_NUMBER_03", "", uint32(1), "EUR", int64(1212), "12.12", uint32(0), uint32(0), "", "2020-03-10T00:00:00Z", "2021-12-16T00:00:00Z", "auth0|66fd06d0bfea78a82bb42459", "auth0|66fd06d0bfea78a82bb42459")
	if err != nil {
		t.Error(err)
		return
	}

	bookingResponse := bkgproto.GetBookingByPkResponse{}
	bookingResponse.Booking = bkg

	gform := commonproto.GetByIdRequest{}
	gform.Id = uint32(16)
	gform.UserEmail = "sprov300@gmail.com"
	gform.RequestId = "bks1m1g91jau4nkks2f0"

	form := bkgproto.GetBookingByPkRequest{}
	form.GetByIdRequest = &gform

	type args struct {
		ctx context.Context
		in  *bkgproto.GetBookingByPkRequest
	}
	tests := []struct {
		bs      *BkgService
		args    args
		want    *bkgproto.GetBookingByPkResponse
		wantErr bool
	}{
		{
			bs: bkgService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &bookingResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {

		bkgResponse, err := tt.bs.GetBookingByPk(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("BkgService.GetBookingByPk() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(bkgResponse, tt.want) {
			t.Errorf("BkgService.GetBookingByPk() = %v, want %v", bkgResponse, tt.want)
		}
		bookingResult := bkgResponse.Booking
		assert.NotNil(t, bookingResult)
		assert.Equal(t, bookingResult.BookingD.CarrierBookingRequestReference, "ABC123123123", "they should be equal")
		assert.Equal(t, bookingResult.BookingD.ExportVoyageNumber, "CARRIER_VOYAGE_NUMBER_03", "they should be equal")
		assert.Equal(t, bookingResult.BookingD.DocumentStatus, "RECE", "they should be equal")
	}
}

func TestBkgService_UpdateBookingByReferenceCarrierBookingRequestReference(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}
	ctx := LoginUser()

	bkgService := NewBkgService(log, dbService, redisService, userServiceClient, currencyService)

	form := bkgproto.UpdateBookingByReferenceCarrierBookingRequestReferenceRequest{}
	form.DocumentStatus = "PENU"
	form.ReceiptTypeAtOrigin = "CY"
	form.DeliveryTypeAtDestination = "CFS"
	form.CargoMovementTypeAtOrigin = "FCL"
	form.CargoMovementTypeAtDestination = "LCL"
	form.ServiceContractReference = "SERVICE_CONTRACT_REFERENCE_01"
	form.PaymentTermCode = "PRE"

	updateResponse := bkgproto.UpdateBookingByReferenceCarrierBookingRequestReferenceResponse{}

	type args struct {
		ctx context.Context
		in  *bkgproto.UpdateBookingByReferenceCarrierBookingRequestReferenceRequest
	}
	tests := []struct {
		bs      *BkgService
		args    args
		want    *bkgproto.UpdateBookingByReferenceCarrierBookingRequestReferenceResponse
		wantErr bool
	}{
		{
			bs: bkgService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &updateResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {

		got, err := tt.bs.UpdateBookingByReferenceCarrierBookingRequestReference(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("BkgService.UpdateBkg() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("BkgService.UpdateBkg() = %v, want %v", got, tt.want)
		}

	}
}

func TestBkgService_GetBookingByCarrierBookingRequestReference(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}
	ctx := LoginUser()

	bkgService := NewBkgService(log, dbService, redisService, userServiceClient, currencyService)

	bkg, err := GetBooking(uint32(16), []byte{43, 149, 187, 185, 210, 212, 68, 22, 172, 254, 98, 7, 225, 129, 181, 244}, "2b95bbb9-d2d4-4416-acfe-6207e181b5f4", "ABC123123123", "RECE", "CY", "CFS", "FCL", "LCL", "SERVICE_CONTRACT_REFERENCE_03", "PRE", true, true, "EXPORT_DECLARATION_REFERENCE_03", false, "IMPORT_LICENSE_REFERENCE_03", "2020-03-10T00:00:00Z", false, true, "", "2020-03-10T00:00:00Z", "2020-03-12T00:00:00Z", "2020-03-13T00:00:00Z", "SWB", "TRANSPORT_DOC_REF_03", "BOOKING_CHA_REF_03", "FCA", "EI", false, "", "", "CARRIER_VOYAGE_NUMBER_03", "", uint32(1), "EUR", int64(1212), "12.12", uint32(0), uint32(0), "", "2020-03-10T00:00:00Z", "2021-12-16T00:00:00Z", "auth0|66fd06d0bfea78a82bb42459", "auth0|66fd06d0bfea78a82bb42459")
	if err != nil {
		t.Error(err)
		return
	}
	bookingResponse := bkgproto.GetBookingByCarrierBookingRequestReferenceResponse{}
	bookingResponse.Booking = bkg

	form := bkgproto.GetBookingByCarrierBookingRequestReferenceRequest{}
	form.CarrierBookingRequestReference = "ABC123123123"
	form.UserEmail = "sprov300@gmail.com"
	form.RequestId = "bks1m1g91jau4nkks2f0"

	type args struct {
		ctx context.Context
		in  *bkgproto.GetBookingByCarrierBookingRequestReferenceRequest
	}
	tests := []struct {
		bs      *BkgService
		args    args
		want    *bkgproto.GetBookingByCarrierBookingRequestReferenceResponse
		wantErr bool
	}{
		{
			bs: bkgService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &bookingResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {

		bkgResponse, err := tt.bs.GetBookingByCarrierBookingRequestReference(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("BkgService.GetBookingByCarrierBookingRequestReference() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(bkgResponse, tt.want) {
			t.Errorf("BkgService.GetBookingByCarrierBookingRequestReference() = %v, want %v", bkgResponse, tt.want)
		}
		bookingResult := bkgResponse.Booking
		assert.NotNil(t, bookingResult)
		assert.Equal(t, bookingResult.BookingD.CarrierBookingRequestReference, "ABC123123123", "they should be equal")
		assert.Equal(t, bookingResult.BookingD.ServiceContractReference, "SERVICE_CONTRACT_REFERENCE_03", "they should be equal")
		assert.Equal(t, bookingResult.BookingD.ExportDeclarationReference, "EXPORT_DECLARATION_REFERENCE_03", "they should be equal")

	}
}

func GetBooking(id uint32, uuid4 []byte, idS string, carrierBookingRequestReference string, documentStatus string, receiptTypeAtOrigin string, deliveryTypeAtDestination string, cargoMovementTypeAtOrigin string, cargoMovementTypeAtDestination string, serviceContractReference string, paymentTermCode string, isPartialLoadAllowed bool, isExportDeclarationRequired bool, exportDeclarationReference string, isImportLicenseRequired bool, importLicenseReference string, submissionDateTime string, isAmsAciFilingRequired bool, isDestinationFilingRequired bool, contractQuotationReference string, expectedDepartureDate string, expectedArrivalAtPlaceOfDeliveryStartDate string, expectedArrivalAtPlaceOfDeliveryEndDate string, transportDocumentTypeCode string, transportDocumentReference string, bookingChannelReference string, incoTerms string, communicationChannelCode string, isEquipmentSubstitutionAllowed bool, vesselName string, vesselImoNumber string, exportVoyageNumber string, preCarriageModeOfTransportCode string, vesselId uint32, declaredValueCurrency string, declaredValue int64, declaredValueString string, voyageId uint32, locationId uint32, invoicePayableAt string, createdAt string, updatedAt string, createdByUserId string, updatedByUserId string) (*bkgproto.Booking, error) {
	createdAt1, err := common.ConvertTimeToTimestamp(Layout, createdAt)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	updatedAt1, err := common.ConvertTimeToTimestamp(Layout, updatedAt)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	submissionDateTime1, err := common.ConvertTimeToTimestamp(Layout, submissionDateTime)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	expectedDepartureDate1, err := common.ConvertTimeToTimestamp(Layout, expectedDepartureDate)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	expectedArrivalAtPlaceOfDeliveryStartDate1, err := common.ConvertTimeToTimestamp(Layout, expectedArrivalAtPlaceOfDeliveryStartDate)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	expectedArrivalAtPlaceOfDeliveryEndDate1, err := common.ConvertTimeToTimestamp(Layout, expectedArrivalAtPlaceOfDeliveryEndDate)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	bookingD := new(bkgproto.BookingD)
	bookingD.Id = id
	bookingD.Uuid4 = uuid4
	bookingD.IdS = idS
	bookingD.CarrierBookingRequestReference = carrierBookingRequestReference
	bookingD.DocumentStatus = documentStatus
	bookingD.ReceiptTypeAtOrigin = receiptTypeAtOrigin
	bookingD.DeliveryTypeAtDestination = deliveryTypeAtDestination
	bookingD.CargoMovementTypeAtOrigin = cargoMovementTypeAtOrigin
	bookingD.CargoMovementTypeAtDestination = cargoMovementTypeAtDestination
	bookingD.ServiceContractReference = serviceContractReference
	bookingD.PaymentTermCode = paymentTermCode
	bookingD.IsPartialLoadAllowed = isPartialLoadAllowed
	bookingD.IsExportDeclarationRequired = isExportDeclarationRequired
	bookingD.ExportDeclarationReference = exportDeclarationReference
	bookingD.IsImportLicenseRequired = isImportLicenseRequired
	bookingD.ImportLicenseReference = importLicenseReference
	bookingD.IsAmsAciFilingRequired = isAmsAciFilingRequired
	bookingD.IsDestinationFilingRequired = isDestinationFilingRequired
	bookingD.ContractQuotationReference = contractQuotationReference
	bookingD.TransportDocumentTypeCode = transportDocumentTypeCode
	bookingD.TransportDocumentReference = transportDocumentReference
	bookingD.BookingChannelReference = bookingChannelReference
	bookingD.IncoTerms = incoTerms
	bookingD.CommunicationChannelCode = communicationChannelCode
	bookingD.IsEquipmentSubstitutionAllowed = isEquipmentSubstitutionAllowed
	bookingD.VesselName = vesselName
	bookingD.VesselImoNumber = vesselImoNumber
	bookingD.ExportVoyageNumber = exportVoyageNumber
	bookingD.PreCarriageModeOfTransportCode = preCarriageModeOfTransportCode
	bookingD.VesselId = vesselId
	bookingD.DeclaredValueCurrency = declaredValueCurrency
	bookingD.DeclaredValue = declaredValue
	bookingD.DeclaredValueString = declaredValueString
	bookingD.VoyageId = voyageId
	bookingD.LocationId = locationId
	bookingD.InvoicePayableAt = invoicePayableAt

	bookingT := new(bkgproto.BookingT)
	bookingT.SubmissionDateTime = submissionDateTime1
	bookingT.ExpectedDepartureDate = expectedDepartureDate1
	bookingT.ExpectedArrivalAtPlaceOfDeliveryStartDate = expectedArrivalAtPlaceOfDeliveryStartDate1
	bookingT.ExpectedArrivalAtPlaceOfDeliveryEndDate = expectedArrivalAtPlaceOfDeliveryEndDate1

	crUpdTime := new(commonproto.CrUpdTime)
	crUpdTime.CreatedAt = createdAt1
	crUpdTime.UpdatedAt = updatedAt1

	crUpdUser := new(commonproto.CrUpdUser)
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = createdByUserId
	crUpdUser.UpdatedByUserId = updatedByUserId

	booking := bkgproto.Booking{BookingD: bookingD, BookingT: bookingT, CrUpdUser: crUpdUser, CrUpdTime: crUpdTime}

	return &booking, nil
}

func TestBkgService_CreateBooking(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}
	ctx := LoginUser()

	bkgService := NewBkgService(log, dbService, redisService, userServiceClient, currencyService)

	bkg := bkgproto.CreateBookingRequest{}
	bkg.CarrierBookingRequestReference = "ef223019-ff16-4870-be69-9dbaaaae9b11"
	bkg.DocumentStatus = "Received"
	bkg.ReceiptTypeAtOrigin = "CY"
	bkg.DeliveryTypeAtDestination = "CY"
	bkg.CargoMovementTypeAtOrigin = "FCL"
	bkg.CargoMovementTypeAtDestination = "FCL"
	bkg.ServiceContractReference = ""
	bkg.PaymentTermCode = "PRE"
	bkg.IsPartialLoadAllowed = false
	bkg.IsExportDeclarationRequired = true
	bkg.ExportDeclarationReference = "export_declaration_reference"
	bkg.IsImportLicenseRequired = true
	bkg.ImportLicenseReference = "import_license_reference"
	bkg.SubmissionDateTime = "02/16/2023"
	bkg.IsAmsAciFilingRequired = false
	bkg.IsDestinationFilingRequired = false
	bkg.ContractQuotationReference = "contractRef"
	bkg.ExpectedDepartureDate = "12/10/2023"
	bkg.ExpectedArrivalAtPlaceOfDeliveryStartDate = "11/09/2023"
	bkg.ExpectedArrivalAtPlaceOfDeliveryEndDate = "11/15/2023"
	bkg.TransportDocumentTypeCode = "BOL"
	bkg.TransportDocumentReference = ""
	bkg.BookingChannelReference = ""
	bkg.IncoTerms = ""
	bkg.CommunicationChannelCode = ""
	bkg.IsEquipmentSubstitutionAllowed = false
	bkg.VesselName = "Rum Runner"
	bkg.VesselImoNumber = "9321483"
	bkg.ExportVoyageNumber = "export-voyage-number"
	bkg.PreCarriageModeOfTransportCode = ""
	bkg.DeclaredValueCurrency = "EUR"
	bkg.DeclaredValue = "100.00"
	bkg.InvoicePayableAt = ""
	bkg.UserId = "auth0|66fd06d0bfea78a82bb42459"
	bkg.UserEmail = "sprov300@gmail.com"
	bkg.RequestId = "bks1m1g91jau4nkks2f0"

	commodity := bkgproto.CreateCommodityRequest{}
	commodity.CommodityType = "Mobile phones"
	commodity.HsCode = "720711"
	commodity.CargoGrossWeight = 12000.00
	commodity.CargoGrossWeightUnit = "KGM"
	commodity.CargoGrossVolume = float64(0)
	commodity.CargoGrossVolumeUnit = ""
	commodity.NumberOfPackages = uint32(0)
	commodity.ExportLicenseIssueDate = "02/22/2023"
	commodity.ExportLicenseExpiryDate = "02/22/2023"

	commodities := []*bkgproto.CreateCommodityRequest{}
	commodities = append(commodities, &commodity)
	bkg.Commodities = commodities

	reference := bkgproto.CreateReferenceRequest{}
	reference.ReferenceTypeCode = "FF"
	reference.ReferenceValue = "test"

	references := []*bkgproto.CreateReferenceRequest{}
	references = append(references, &reference)
	bkg.References = references

	valueAddedServiceRequest := bkgproto.CreateValueAddedServiceRequest{}
	valueAddedServiceRequest.ValueAddedServiceCode = "CDECL"

	valueAddedServiceRequests := []*bkgproto.CreateValueAddedServiceRequest{}
	valueAddedServiceRequests = append(valueAddedServiceRequests, &valueAddedServiceRequest)
	bkg.ValueAddedServiceRequests = valueAddedServiceRequests

	requestedEquipment := bkgproto.CreateRequestedEquipmentRequest{}
	requestedEquipment.RequestedEquipmentSizetype = "22GP"
	requestedEquipment.RequestedEquipmentUnits = 3

	requestedEquipments := []*bkgproto.CreateRequestedEquipmentRequest{}
	requestedEquipments = append(requestedEquipments, &requestedEquipment)
	bkg.RequestedEquipments = requestedEquipments

	shipmentLocation := bkgproto.CreateShipmentLocationRequest{}
	shipmentLocation.ShipmentLocationTypeCode = "FCD"
	shipmentLocation.DisplayedName = "Singapore"
	shipmentLocation.EventDateTime = "02/22/2023"

	shipmentLocations := []*bkgproto.CreateShipmentLocationRequest{}
	shipmentLocations = append(shipmentLocations, &shipmentLocation)
	bkg.ShipmentLocations = shipmentLocations

	type args struct {
		ctx context.Context
		in  *bkgproto.CreateBookingRequest
	}
	tests := []struct {
		bs      *BkgService
		args    args
		wantErr bool
	}{
		{
			bs: bkgService,
			args: args{
				ctx: ctx,
				in:  &bkg,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		bookingResponse, err := tt.bs.CreateBooking(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("BkgService.CreateBooking() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		bookingResult := bookingResponse.Booking
		fmt.Println("bookingResult", bookingResult)
		assert.NotNil(t, bookingResult)
		assert.Equal(t, bookingResult.BookingD.CarrierBookingRequestReference, "ef223019-ff16-4870-be69-9dbaaaae9b11", "they should be equal")
		assert.Equal(t, bookingResult.BookingD.DocumentStatus, "Received", "they should be equal")
		assert.NotNil(t, bookingResult.CrUpdTime.CreatedAt)
		assert.NotNil(t, bookingResult.CrUpdTime.UpdatedAt)
		assert.NotNil(t, bookingResult.Commodities)
		assert.NotNil(t, bookingResult.References)
		assert.NotNil(t, bookingResult.ValueAddedServiceRequests)
		assert.NotNil(t, bookingResult.RequestedEquipments)
		assert.NotNil(t, bookingResult.ShipmentLocations)
	}
}
