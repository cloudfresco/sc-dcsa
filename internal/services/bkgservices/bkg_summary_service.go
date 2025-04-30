package bkgservices

// https://github.com/dcsaorg/DCSA-BKG/blob/master/src/main/java/org/dcsa/bkg/service/impl/BookingSummaryServiceImpl.java
import (
	"context"
	"time"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	bkgproto "github.com/cloudfresco/sc-dcsa/internal/protogen/bkg/v2"
	commonproto "github.com/cloudfresco/sc-dcsa/internal/protogen/common/v1"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"
	partyservice "github.com/cloudfresco/sc-dcsa/internal/services/partyservices"
	bkgstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/bkg/v2"
	commonstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/common/v1"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// BkgSummaryService - For accessing BkgSummary services
type BkgSummaryService struct {
	log               *zap.Logger
	DBService         *common.DBService
	RedisService      *common.RedisService
	UserServiceClient partyproto.UserServiceClient
	bkgproto.UnimplementedBkgSummaryServiceServer
}

// NewBkgSummaryService - Create BkgSummary Service
func NewBkgSummaryService(log *zap.Logger, dbOpt *common.DBService, redisOpt *common.RedisService, userServiceClient partyproto.UserServiceClient) *BkgSummaryService {
	return &BkgSummaryService{
		log:               log,
		DBService:         dbOpt,
		RedisService:      redisOpt,
		UserServiceClient: userServiceClient,
	}
}

// insertBookingSummarySQL - insert BookingSummarySQL query
const insertBookingSummarySQL = `insert into booking_summaries
	  ( 
  uuid4,
  carrier_booking_request_reference,
  document_status,
  receipt_type_at_origin,
  delivery_type_at_destination,
  cargo_movement_type_at_origin,
  cargo_movement_type_at_destination,
  service_contract_reference,
  vessel_name,
  carrier_export_voyage_number,
  universal_export_voyage_reference,
  declared_value,
  delivery_value_currency,
  payment_term_code,
  is_partial_load_allowed,
  is_export_declaration_required,
  export_declaration_reference,
  is_import_license_required,
  import_license_reference,
  is_ams_aci_filing_required,
  is_destination_filing_required,
  contract_quotation_reference,
  transport_document_type_code,
  transport_document_reference,
  booking_channel_reference,
  inco_terms,
  communication_channel_code,
  is_equipment_substitution_allowed,
  vessel_imo_number,
  pre_carriage_mode_of_transport_code, 
  booking_request_created_date_time,
  booking_request_updated_date_time,
  submission_date_time,
  expected_departure_date,
  expected_arrival_at_place_of_delivery_start_date,
  expected_arrival_at_place_of_delivery_end_date,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at)
  values (:uuid4,
  :carrier_booking_request_reference,
  :document_status,
  :receipt_type_at_origin,
  :delivery_type_at_destination,
  :cargo_movement_type_at_origin,
  :cargo_movement_type_at_destination,
  :service_contract_reference,
  :vessel_name,
  :carrier_export_voyage_number,
  :universal_export_voyage_reference,
  :declared_value,
  :delivery_value_currency,
  :payment_term_code,
  :is_partial_load_allowed,
  :is_export_declaration_required,
  :export_declaration_reference,
  :is_import_license_required,
  :import_license_reference,
  :is_ams_aci_filing_required,
  :is_destination_filing_required,
  :contract_quotation_reference,
  :transport_document_type_code,
  :transport_document_reference,
  :booking_channel_reference,
  :inco_terms,
  :communication_channel_code,
  :is_equipment_substitution_allowed,
  :vessel_imo_number,
  :pre_carriage_mode_of_transport_code,
  :booking_request_created_date_time,
  :booking_request_updated_date_time,
  :submission_date_time,
  :expected_departure_date,
  :expected_arrival_at_place_of_delivery_start_date,
  :expected_arrival_at_place_of_delivery_end_date,
  :status_code,
  :created_by_user_id,
  :updated_by_user_id,
  :created_at,
  :updated_at);`

// selectBookingSummariesSQL - select BookingSummariesSQL query
const selectBookingSummariesSQL = `select 
  id,
  uuid4,
  carrier_booking_request_reference,
  document_status,
  receipt_type_at_origin,
  delivery_type_at_destination,
  cargo_movement_type_at_origin,
  cargo_movement_type_at_destination,
  service_contract_reference,
  vessel_name,
  carrier_export_voyage_number,
  universal_export_voyage_reference,
  declared_value,
  delivery_value_currency,
  payment_term_code,
  is_partial_load_allowed,
  is_export_declaration_required,
  export_declaration_reference,
  is_import_license_required,
  import_license_reference,
  is_ams_aci_filing_required,
  is_destination_filing_required,
  contract_quotation_reference,
  transport_document_type_code,
  transport_document_reference,
  booking_channel_reference,
  inco_terms,
  communication_channel_code,
  is_equipment_substitution_allowed,
  vessel_imo_number,
  pre_carriage_mode_of_transport_code, 
  booking_request_created_date_time,
  booking_request_updated_date_time,
  submission_date_time,
  expected_departure_date,
  expected_arrival_at_place_of_delivery_start_date,
  expected_arrival_at_place_of_delivery_end_date,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at from booking_summaries`

// CreateBookingSummary - Create BookingSummary
func (bs *BkgSummaryService) CreateBookingSummary(ctx context.Context, in *bkgproto.CreateBookingSummaryRequest) (*bkgproto.CreateBookingSummaryResponse, error) {
	user, err := partyservice.GetUserWithNewContext(ctx, in.UserId, in.UserEmail, in.RequestId, bs.UserServiceClient)
	if err != nil {
		bs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	ttime := common.GetTimeDetails()
	tn := common.TimeToTimestamp(ttime)

	submissionDateTime, err := time.Parse(common.Layout, in.SubmissionDateTime)
	if err != nil {
		bs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	bookingRequestCreatedDateTime, err := time.Parse(common.Layout, in.BookingRequestCreatedDateTime)
	if err != nil {
		bs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	bookingRequestUpdatedDateTime, err := time.Parse(common.Layout, in.BookingRequestUpdatedDateTime)
	if err != nil {
		bs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	expectedDepartureDate, err := time.Parse(common.Layout, in.ExpectedDepartureDate)
	if err != nil {
		bs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	expectedArrivalAtPlaceOfDeliveryStartDate, err := time.Parse(common.Layout, in.ExpectedArrivalAtPlaceOfDeliveryStartDate)
	if err != nil {
		bs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	expectedArrivalAtPlaceOfDeliveryEndDate, err := time.Parse(common.Layout, in.ExpectedArrivalAtPlaceOfDeliveryEndDate)
	if err != nil {
		bs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	bkgSummaryD := bkgproto.BookingSummaryD{}
	bkgSummaryD.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		bs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	bkgSummaryD.CarrierBookingRequestReference = in.CarrierBookingRequestReference
	bkgSummaryD.DocumentStatus = in.DocumentStatus
	bkgSummaryD.ReceiptTypeAtOrigin = in.ReceiptTypeAtOrigin
	bkgSummaryD.DeliveryTypeAtDestination = in.DeliveryTypeAtDestination
	bkgSummaryD.CargoMovementTypeAtOrigin = in.CargoMovementTypeAtOrigin
	bkgSummaryD.CargoMovementTypeAtDestination = in.CargoMovementTypeAtDestination
	bkgSummaryD.ServiceContractReference = in.ServiceContractReference
	bkgSummaryD.VesselName = in.VesselName
	bkgSummaryD.CarrierExportVoyageNumber = in.CarrierExportVoyageNumber
	bkgSummaryD.UniversalExportVoyageReference = in.UniversalExportVoyageReference
	bkgSummaryD.DeclaredValue = in.DeclaredValue
	bkgSummaryD.DeliveryValueCurrency = in.DeliveryValueCurrency
	bkgSummaryD.PaymentTermCode = in.PaymentTermCode
	bkgSummaryD.IsPartialLoadAllowed = in.IsPartialLoadAllowed
	bkgSummaryD.IsExportDeclarationRequired = in.IsExportDeclarationRequired
	bkgSummaryD.ExportDeclarationReference = in.ExportDeclarationReference
	bkgSummaryD.IsImportLicenseRequired = in.IsImportLicenseRequired
	bkgSummaryD.ImportLicenseReference = in.ImportLicenseReference
	bkgSummaryD.IsAmsAciFilingRequired = in.IsAmsAciFilingRequired
	bkgSummaryD.IsDestinationFilingRequired = in.IsDestinationFilingRequired
	bkgSummaryD.ContractQuotationReference = in.ContractQuotationReference
	bkgSummaryD.TransportDocumentTypeCode = in.TransportDocumentTypeCode
	bkgSummaryD.TransportDocumentReference = in.TransportDocumentReference
	bkgSummaryD.BookingChannelReference = in.BookingChannelReference
	bkgSummaryD.IncoTerms = in.IncoTerms
	bkgSummaryD.CommunicationChannelCode = in.CommunicationChannelCode
	bkgSummaryD.IsEquipmentSubstitutionAllowed = in.IsEquipmentSubstitutionAllowed
	bkgSummaryD.VesselImoNumber = in.VesselImoNumber
	bkgSummaryD.PreCarriageModeOfTransportCode = in.PreCarriageModeOfTransportCode

	bkgSummaryT := bkgproto.BookingSummaryT{}
	bkgSummaryT.BookingRequestCreatedDateTime = common.TimeToTimestamp(bookingRequestCreatedDateTime.UTC().Truncate(time.Second))
	bkgSummaryT.BookingRequestUpdatedDateTime = common.TimeToTimestamp(bookingRequestUpdatedDateTime.UTC().Truncate(time.Second))
	bkgSummaryT.SubmissionDateTime = common.TimeToTimestamp(submissionDateTime.UTC().Truncate(time.Second))
	bkgSummaryT.ExpectedDepartureDate = common.TimeToTimestamp(expectedDepartureDate.UTC().Truncate(time.Second))
	bkgSummaryT.ExpectedArrivalAtPlaceOfDeliveryStartDate = common.TimeToTimestamp(expectedArrivalAtPlaceOfDeliveryStartDate.UTC().Truncate(time.Second))
	bkgSummaryT.ExpectedArrivalAtPlaceOfDeliveryEndDate = common.TimeToTimestamp(expectedArrivalAtPlaceOfDeliveryEndDate.UTC().Truncate(time.Second))

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	bkgSummary := bkgproto.BookingSummary{BookingSummaryD: &bkgSummaryD, BookingSummaryT: &bkgSummaryT, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	err = bs.insertBookingSummary(ctx, insertBookingSummarySQL, &bkgSummary, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		bs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	bkgSummaryResponse := bkgproto.CreateBookingSummaryResponse{}
	bkgSummaryResponse.BookingSummary = &bkgSummary
	return &bkgSummaryResponse, nil
}

// insertBookingSummary - Insert Booking Summary
func (bs *BkgSummaryService) insertBookingSummary(ctx context.Context, insertBookingSummarySQL string, bkgSummary *bkgproto.BookingSummary, userEmail string, requestID string) error {
	bkgSummaryTmp, err := bs.crBookingSummaryStruct(ctx, bkgSummary, userEmail, requestID)
	if err != nil {
		bs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	err = bs.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertBookingSummarySQL, bkgSummaryTmp)
		if err != nil {
			bs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			bs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		bkgSummary.BookingSummaryD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(bkgSummary.BookingSummaryD.Uuid4)
		if err != nil {
			bs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		bkgSummary.BookingSummaryD.IdS = uuid4Str
		return nil
	})

	if err != nil {
		bs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// crBookingSummaryStruct - process BookingSummary details
func (bs *BkgSummaryService) crBookingSummaryStruct(ctx context.Context, bkgSummary *bkgproto.BookingSummary, userEmail string, requestID string) (*bkgstruct.BookingSummary, error) {
	bkgSummaryT := new(bkgstruct.BookingSummaryT)
	bkgSummaryT.BookingRequestCreatedDateTime = common.TimestampToTime(bkgSummary.BookingSummaryT.BookingRequestCreatedDateTime)
	bkgSummaryT.BookingRequestUpdatedDateTime = common.TimestampToTime(bkgSummary.BookingSummaryT.BookingRequestUpdatedDateTime)
	bkgSummaryT.SubmissionDateTime = common.TimestampToTime(bkgSummary.BookingSummaryT.SubmissionDateTime)
	bkgSummaryT.ExpectedDepartureDate = common.TimestampToTime(bkgSummary.BookingSummaryT.ExpectedDepartureDate)
	bkgSummaryT.ExpectedArrivalAtPlaceOfDeliveryStartDate = common.TimestampToTime(bkgSummary.BookingSummaryT.ExpectedArrivalAtPlaceOfDeliveryStartDate)
	bkgSummaryT.ExpectedArrivalAtPlaceOfDeliveryEndDate = common.TimestampToTime(bkgSummary.BookingSummaryT.ExpectedArrivalAtPlaceOfDeliveryEndDate)

	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(bkgSummary.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(bkgSummary.CrUpdTime.UpdatedAt)
	bkgSummaryTmp := bkgstruct.BookingSummary{BookingSummaryD: bkgSummary.BookingSummaryD, BookingSummaryT: bkgSummaryT, CrUpdUser: bkgSummary.CrUpdUser, CrUpdTime: crUpdTime}

	return &bkgSummaryTmp, nil
}

// GetBookingSummaryByCarrierBookingRequestReference - Get BookingSummaryByCarrierBookingReference
func (bs *BkgSummaryService) GetBookingSummaryByCarrierBookingRequestReference(ctx context.Context, in *bkgproto.GetBookingSummaryByCarrierBookingRequestReferenceRequest) (*bkgproto.GetBookingSummaryByCarrierBookingRequestReferenceResponse, error) {
	nselectBookingSummariesSQL := selectBookingSummariesSQL + ` where carrier_booking_request_reference = ? and status_code = ?;`
	row := bs.DBService.DB.QueryRowxContext(ctx, nselectBookingSummariesSQL, in.CarrierBookingRequestReference, "active")
	bkgSummaryTmp := bkgstruct.BookingSummary{}
	err := row.StructScan(&bkgSummaryTmp)
	if err != nil {
		bs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	getRequest := commonproto.GetRequest{}
	getRequest.UserEmail = in.UserEmail
	getRequest.RequestId = in.RequestId

	bkgSummary, err := bs.getBookingSummaryStruct(ctx, &getRequest, bkgSummaryTmp)
	if err != nil {
		bs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	bookingSummaryResponse := bkgproto.GetBookingSummaryByCarrierBookingRequestReferenceResponse{}
	bookingSummaryResponse.BookingSummary = bkgSummary

	return &bookingSummaryResponse, nil
}

// getBookingSummaryStruct - Get booking Summary header
func (bs *BkgSummaryService) getBookingSummaryStruct(ctx context.Context, in *commonproto.GetRequest, bkgSummaryTmp bkgstruct.BookingSummary) (*bkgproto.BookingSummary, error) {
	bkgSummaryT := new(bkgproto.BookingSummaryT)
	bkgSummaryT.BookingRequestCreatedDateTime = common.TimeToTimestamp(bkgSummaryTmp.BookingSummaryT.BookingRequestCreatedDateTime)
	bkgSummaryT.BookingRequestUpdatedDateTime = common.TimeToTimestamp(bkgSummaryTmp.BookingSummaryT.BookingRequestUpdatedDateTime)
	bkgSummaryT.SubmissionDateTime = common.TimeToTimestamp(bkgSummaryTmp.BookingSummaryT.SubmissionDateTime)
	bkgSummaryT.ExpectedDepartureDate = common.TimeToTimestamp(bkgSummaryTmp.BookingSummaryT.ExpectedDepartureDate)
	bkgSummaryT.ExpectedArrivalAtPlaceOfDeliveryStartDate = common.TimeToTimestamp(bkgSummaryTmp.BookingSummaryT.ExpectedArrivalAtPlaceOfDeliveryStartDate)
	bkgSummaryT.ExpectedArrivalAtPlaceOfDeliveryEndDate = common.TimeToTimestamp(bkgSummaryTmp.BookingSummaryT.ExpectedArrivalAtPlaceOfDeliveryEndDate)

	crUpdTime := new(commonproto.CrUpdTime)
	crUpdTime.CreatedAt = common.TimeToTimestamp(bkgSummaryTmp.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimeToTimestamp(bkgSummaryTmp.CrUpdTime.UpdatedAt)

	uuid4Str, err := common.UUIDBytesToStr(bkgSummaryTmp.BookingSummaryD.Uuid4)
	if err != nil {
		bs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	bkgSummaryTmp.BookingSummaryD.IdS = uuid4Str

	bkgSummary := bkgproto.BookingSummary{BookingSummaryD: bkgSummaryTmp.BookingSummaryD, BookingSummaryT: bkgSummaryT, CrUpdUser: bkgSummaryTmp.CrUpdUser, CrUpdTime: crUpdTime}

	return &bkgSummary, nil
}

// GetBookingSummaries - Get BookingSummaries
func (bs *BkgSummaryService) GetBookingSummaries(ctx context.Context, in *bkgproto.GetBookingSummariesRequest) (*bkgproto.GetBookingSummariesResponse, error) {
	limit := in.GetLimit()
	nextCursor := in.GetNextCursor()
	if limit == "" {
		limit = bs.DBService.LimitSQLRows
	}
	query := "status_code = ?"
	if nextCursor == "" {
		query = query + " order by id desc " + " limit " + limit + ";"
	} else {
		nextCursor = common.DecodeCursor(nextCursor)
		query = query + " " + "and" + " " + "id <= " + nextCursor + " order by id desc " + " limit " + limit + ";"
	}

	bkgSummaries := []*bkgproto.BookingSummary{}

	nselectBookingSummariesSQL := selectBookingSummariesSQL + ` where ` + query

	rows, err := bs.DBService.DB.QueryxContext(ctx, nselectBookingSummariesSQL, "active")
	if err != nil {
		bs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	for rows.Next() {

		bkgSummaryTmp := bkgstruct.BookingSummary{}
		err = rows.StructScan(&bkgSummaryTmp)
		if err != nil {
			bs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		getRequest := commonproto.GetRequest{}
		getRequest.UserEmail = in.UserEmail
		getRequest.RequestId = in.RequestId
		bkgSummary, err := bs.getBookingSummaryStruct(ctx, &getRequest, bkgSummaryTmp)
		if err != nil {
			bs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		bkgSummaries = append(bkgSummaries, bkgSummary)

	}

	bkgSummaryResponse := bkgproto.GetBookingSummariesResponse{}
	if len(bkgSummaries) != 0 {
		next := bkgSummaries[len(bkgSummaries)-1].BookingSummaryD.Id
		next--
		nextc := common.EncodeCursor(next)
		bkgSummaryResponse = bkgproto.GetBookingSummariesResponse{BookingSummaries: bkgSummaries, NextCursor: nextc}
	} else {
		bkgSummaryResponse = bkgproto.GetBookingSummariesResponse{BookingSummaries: bkgSummaries, NextCursor: "0"}
	}
	return &bkgSummaryResponse, nil
}
