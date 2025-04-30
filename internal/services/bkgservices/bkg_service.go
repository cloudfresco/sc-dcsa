// https://github.com/dcsaorg/DCSA-Edocumentation/blob/master/edocumentation-service/src/main/java/org/dcsa/edocumentation/service/BookingService.java
// https://github.com/dcsaorg/DCSA-BKG/blob/master/src/test/java/org/dcsa/bkg/service/impl/BKGServiceImplTest.java
package bkgservices

import (
	"context"
	"net"
	"os"
	"time"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	"github.com/cloudfresco/sc-dcsa/internal/config"
	bkgproto "github.com/cloudfresco/sc-dcsa/internal/protogen/bkg/v2"
	commonproto "github.com/cloudfresco/sc-dcsa/internal/protogen/common/v1"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"
	partyservice "github.com/cloudfresco/sc-dcsa/internal/services/partyservices"
	bkgstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/bkg/v2"
	commonstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/common/v1"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// BkgService - For accessing Booking services
type BkgService struct {
	log               *zap.Logger
	DBService         *common.DBService
	RedisService      *common.RedisService
	UserServiceClient partyproto.UserServiceClient
	bkgproto.UnimplementedBkgServiceServer
}

// NewBkgService - Create Booking service
func NewBkgService(log *zap.Logger, dbOpt *common.DBService, redisOpt *common.RedisService, userServiceClient partyproto.UserServiceClient) *BkgService {
	return &BkgService{
		log:               log,
		DBService:         dbOpt,
		RedisService:      redisOpt,
		UserServiceClient: userServiceClient,
	}
}

// insertBookingSQL - insert BookingSQL query
const insertBookingSQL = `insert into bookings
	  ( 
  uuid4,
  carrier_booking_request_reference,
  document_status,
  receipt_type_at_origin,
  delivery_type_at_destination,
  cargo_movement_type_at_origin,
  cargo_movement_type_at_destination,
  service_contract_reference,
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
  vessel_name,
  vessel_imo_number,
  export_voyage_number,
  pre_carriage_mode_of_transport_code,
  vessel_id,
  declared_value_currency_code,
  declared_value,
  voyage_id,
  location_id,
  invoice_payable_at,
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
:vessel_name,
:vessel_imo_number,
:export_voyage_number,
:pre_carriage_mode_of_transport_code,
:vessel_id,
:declared_value_currency_code,
:declared_value,
:voyage_id,
:location_id,
:invoice_payable_at,
:submission_date_time,
:expected_departure_date,
:expected_arrival_at_place_of_delivery_start_date,
:expected_arrival_at_place_of_delivery_end_date,
:status_code,
:created_by_user_id,
:updated_by_user_id,
:created_at,
:updated_at);`

// updateBookingsSQL - update BookingsSQL query
const updateBookingSQL = `update bookings set 
		  document_status = ?,
      receipt_type_at_origin = ?,
      delivery_type_at_destination = ?,
      cargo_movement_type_at_origin = ?,
      cargo_movement_type_at_destination = ?,
      service_contract_reference = ?,
      payment_term_code = ?,
			updated_at = ? where carrier_booking_request_reference = ?;`

// selectBookingsSQL - select BookingsSQL query
const selectBookingsSQL = `select 
  id,
  uuid4,
  carrier_booking_request_reference,
  document_status,
  receipt_type_at_origin,
  delivery_type_at_destination,
  cargo_movement_type_at_origin,
  cargo_movement_type_at_destination,
  service_contract_reference,
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
  vessel_name,
  vessel_imo_number,
  export_voyage_number,
  pre_carriage_mode_of_transport_code,
  vessel_id,
  declared_value_currency_code,
  declared_value,
  voyage_id,
  location_id,
  invoice_payable_at,
  submission_date_time,
  expected_departure_date,
  expected_arrival_at_place_of_delivery_start_date,
  expected_arrival_at_place_of_delivery_end_date,
  status_code,
  created_by_user_id,
  updated_by_user_id, 
  created_at,
  updated_at from bookings`

// StartBkgServer - Start Bkg server
func StartBkgServer(log *zap.Logger, isTest bool, pwd string, dbOpt *config.DBOptions, redisOpt *config.RedisOptions, mailerOpt *config.MailerOptions, grpcServerOpt *config.GrpcServerOptions, jwtOpt *config.JWTOptions, oauthOpt *config.OauthOptions, userOpt *config.UserOptions, uptraceOpt *config.UptraceOptions, dbService *common.DBService, redisService *common.RedisService, mailerService common.MailerIntf) {
	common.SetJWTOpt(jwtOpt)

	creds, err := common.GetSrvCred(log, isTest, pwd, grpcServerOpt)
	if err != nil {
		os.Exit(1)
	}

	userCreds, err := common.GetClientCred(log, isTest, pwd, grpcServerOpt)
	if err != nil {
		os.Exit(1)
	}

	var srvOpts []grpc.ServerOption

	userConn, err := grpc.NewClient(grpcServerOpt.GrpcUserServerPort, grpc.WithTransportCredentials(userCreds), grpc.WithStatsHandler(otelgrpc.NewClientHandler()))
	if err != nil {
		log.Error("Error", zap.Error(err))
		os.Exit(1)
	}

	srvOpts = append(srvOpts, grpc.Creds(creds))

	srvOpts = append(srvOpts, grpc.StatsHandler(otelgrpc.NewServerHandler()))

	uc := partyproto.NewUserServiceClient(userConn)
	bkgService := NewBkgService(log, dbService, redisService, uc)
	bkgShipmentSummaryService := NewBkgShipmentSummaryService(log, dbService, redisService, uc)
	bkgSummaryService := NewBkgSummaryService(log, dbService, redisService, uc)
	referenceService := NewReferenceService(log, dbService, redisService, uc)

	lis, err := net.Listen("tcp", grpcServerOpt.GrpcBkgServerPort)
	if err != nil {
		log.Error("Error", zap.Error(err))
		os.Exit(1)
	}

	srv := grpc.NewServer(srvOpts...)
	bkgproto.RegisterBkgServiceServer(srv, bkgService)
	bkgproto.RegisterBkgShipmentSummaryServiceServer(srv, bkgShipmentSummaryService)
	bkgproto.RegisterBkgSummaryServiceServer(srv, bkgSummaryService)
	bkgproto.RegisterReferenceServiceServer(srv, referenceService)

	if err := srv.Serve(lis); err != nil {
		log.Error("Error", zap.Error(err))
		os.Exit(1)
	}
}

// CreateBooking - Create Booking
func (bs *BkgService) CreateBooking(ctx context.Context, in *bkgproto.CreateBookingRequest) (*bkgproto.CreateBookingResponse, error) {
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

	bookingD := bkgproto.BookingD{}
	bookingD.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		bs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	bookingD.CarrierBookingRequestReference = in.CarrierBookingRequestReference
	bookingD.DocumentStatus = in.DocumentStatus
	bookingD.ReceiptTypeAtOrigin = in.ReceiptTypeAtOrigin
	bookingD.DeliveryTypeAtDestination = in.DeliveryTypeAtDestination
	bookingD.CargoMovementTypeAtOrigin = in.CargoMovementTypeAtOrigin
	bookingD.CargoMovementTypeAtDestination = in.CargoMovementTypeAtDestination
	bookingD.ServiceContractReference = in.ServiceContractReference
	bookingD.PaymentTermCode = in.PaymentTermCode
	bookingD.IsPartialLoadAllowed = in.IsPartialLoadAllowed
	bookingD.IsExportDeclarationRequired = in.IsExportDeclarationRequired
	bookingD.ExportDeclarationReference = in.ExportDeclarationReference
	bookingD.IsImportLicenseRequired = in.IsImportLicenseRequired
	bookingD.ImportLicenseReference = in.ImportLicenseReference
	bookingD.IsAmsAciFilingRequired = in.IsAmsAciFilingRequired
	bookingD.IsDestinationFilingRequired = in.IsDestinationFilingRequired
	bookingD.ContractQuotationReference = in.ContractQuotationReference
	bookingD.TransportDocumentTypeCode = in.TransportDocumentTypeCode
	bookingD.TransportDocumentReference = in.TransportDocumentReference
	bookingD.BookingChannelReference = in.BookingChannelReference
	bookingD.IncoTerms = in.IncoTerms
	bookingD.CommunicationChannelCode = in.CommunicationChannelCode
	bookingD.IsEquipmentSubstitutionAllowed = in.IsEquipmentSubstitutionAllowed
	bookingD.VesselName = in.VesselName
	bookingD.VesselImoNumber = in.VesselImoNumber
	bookingD.ExportVoyageNumber = in.ExportVoyageNumber
	bookingD.PreCarriageModeOfTransportCode = in.PreCarriageModeOfTransportCode
	bookingD.VesselId = in.VesselId
	bookingD.DeclaredValueCurrencyCode = in.DeclaredValueCurrencyCode
	bookingD.DeclaredValue = in.DeclaredValue
	bookingD.VoyageId = in.VoyageId
	bookingD.LocationId = in.LocationId
	bookingD.InvoicePayableAt = in.InvoicePayableAt

	bookingT := bkgproto.BookingT{}
	bookingT.SubmissionDateTime = common.TimeToTimestamp(submissionDateTime.UTC().Truncate(time.Second))
	bookingT.ExpectedDepartureDate = common.TimeToTimestamp(expectedDepartureDate.UTC().Truncate(time.Second))
	bookingT.ExpectedArrivalAtPlaceOfDeliveryStartDate = common.TimeToTimestamp(expectedArrivalAtPlaceOfDeliveryStartDate.UTC().Truncate(time.Second))
	bookingT.ExpectedArrivalAtPlaceOfDeliveryEndDate = common.TimeToTimestamp(expectedArrivalAtPlaceOfDeliveryEndDate.UTC().Truncate(time.Second))

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	booking := bkgproto.Booking{BookingD: &bookingD, BookingT: &bookingT, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	commodities := []*bkgproto.Commodity{}
	// we will do for loop on commodities, valueAddedServiceRequests, references, requestedEquipments, shipmentLocations which is comes from client form
	for _, commodity := range in.Commodities {
		commodity.UserId = in.UserId
		commodity.UserEmail = in.UserEmail
		commodity.RequestId = in.RequestId
		commodity, err := bs.ProcessCommodityRequest(ctx, commodity)
		if err != nil {
			bs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		commodities = append(commodities, commodity)
	}

	references := []*bkgproto.Reference1{}
	referenceServ := &ReferenceService{DBService: bs.DBService, RedisService: bs.RedisService, UserServiceClient: bs.UserServiceClient}
	for _, reference := range in.References {
		reference.UserId = in.UserId
		reference.UserEmail = in.UserEmail
		reference.RequestId = in.RequestId
		reference, err := referenceServ.ProcessReferenceRequest(ctx, reference)
		if err != nil {
			bs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		references = append(references, reference)
	}

	valueAddedServiceRequests := []*bkgproto.ValueAddedServiceRequest{}
	for _, vreq := range in.ValueAddedServiceRequests {
		vreq.UserId = in.UserId
		vreq.UserEmail = in.UserEmail
		vreq.RequestId = in.RequestId
		vreq, err := bs.ProcessValueAddedServiceRequest(ctx, vreq)
		if err != nil {
			bs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		valueAddedServiceRequests = append(valueAddedServiceRequests, vreq)
	}

	requestedEquipments := []*bkgproto.RequestedEquipment{}
	for _, reqEquipment := range in.RequestedEquipments {
		reqEquipment.UserId = in.UserId
		reqEquipment.UserEmail = in.UserEmail
		reqEquipment.RequestId = in.RequestId
		reqEquipment, err := bs.ProcessRequestedEquipmentRequest(ctx, reqEquipment)
		if err != nil {
			bs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		requestedEquipments = append(requestedEquipments, reqEquipment)
	}

	shipmentLocations := []*bkgproto.ShipmentLocation{}
	for _, shipmentLocation := range in.ShipmentLocations {
		shipmentLocation.UserId = in.UserId
		shipmentLocation.UserEmail = in.UserEmail
		shipmentLocation.RequestId = in.RequestId
		shipmentLocation, err := bs.ProcessShipmentLocationRequest(ctx, shipmentLocation)
		if err != nil {
			bs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		shipmentLocations = append(shipmentLocations, shipmentLocation)
	}

	err = bs.insertBooking(ctx, insertBookingSQL, &booking, insertCommoditySQL, commodities, InsertReferenceSQL, references, insertValueAddedServiceRequestSQL, valueAddedServiceRequests, insertRequestedEquipmentSQL, requestedEquipments, insertShipmentLocationSQL, shipmentLocations, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		bs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	bookingResponse := bkgproto.CreateBookingResponse{}
	bookingResponse.Booking = &booking
	return &bookingResponse, nil
}

// insertBooking - Insert Booking
func (bs *BkgService) insertBooking(ctx context.Context, insertBookingSQL string, booking *bkgproto.Booking, insertCommoditySQL string, commodities []*bkgproto.Commodity, insertReferenceSQL string, references []*bkgproto.Reference1, insertValueAddedServiceRequestsSQL string, valueAddedServiceRequests []*bkgproto.ValueAddedServiceRequest, insertRequestedEquipmentSQL string, requestedEquipments []*bkgproto.RequestedEquipment, insertShipmentLocationSQL string, shipmentLocations []*bkgproto.ShipmentLocation, userEmail string, requestID string) error {
	bookingTmp, err := bs.crBkgStruct(ctx, booking, userEmail, requestID)
	if err != nil {
		bs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	err = bs.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertBookingSQL, bookingTmp)
		if err != nil {
			bs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			bs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		booking.BookingD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(booking.BookingD.Uuid4)
		if err != nil {
			bs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		booking.BookingD.IdS = uuid4Str
		for _, commodity := range commodities {
			commodity.CommodityD.BookingId = booking.BookingD.Id
			commodityTmp, err := bs.CrCommodityStruct(ctx, commodity, userEmail, requestID)
			if err != nil {
				bs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
				return err
			}
			_, err = tx.NamedExecContext(ctx, insertCommoditySQL, commodityTmp)
			if err != nil {
				bs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
				return err
			}
		}

		for _, reference := range references {
			reference.Reference1D.BookingId = booking.BookingD.Id
			_, err = tx.NamedExecContext(ctx, insertReferenceSQL, reference)
			if err != nil {
				bs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
				return err
			}
		}
		for _, valueAddedServiceRequest := range valueAddedServiceRequests {
			valueAddedServiceRequest.BookingId = booking.BookingD.Id

			_, err = tx.NamedExecContext(ctx, insertValueAddedServiceRequestSQL, valueAddedServiceRequest)
			if err != nil {
				bs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
				return err
			}
		}

		for _, requestedEquipment := range requestedEquipments {
			requestedEquipment.RequestedEquipmentD.BookingId = booking.BookingD.Id
			_, err = tx.NamedExecContext(ctx, insertRequestedEquipmentSQL, requestedEquipment)
			if err != nil {
				bs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
				return err
			}
		}

		for _, shipmentLocation := range shipmentLocations {
			shipmentLocation.ShipmentLocationD.BookingId = booking.BookingD.Id
			shipmentLocationTmp, err := bs.CrShipmentLocationStruct(ctx, shipmentLocation, userEmail, requestID)
			if err != nil {
				bs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
				return err
			}
			_, err = tx.NamedExecContext(ctx, insertShipmentLocationSQL, shipmentLocationTmp)
			if err != nil {
				bs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
				return err
			}
		}

		return nil
	})
	booking.Commodities = commodities
	booking.References = references
	booking.ValueAddedServiceRequests = valueAddedServiceRequests
	booking.RequestedEquipments = requestedEquipments
	booking.ShipmentLocations = shipmentLocations
	if err != nil {
		bs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// crBkgStruct - process Booking details
func (bs *BkgService) crBkgStruct(ctx context.Context, booking *bkgproto.Booking, userEmail string, requestID string) (*bkgstruct.Booking, error) {
	bookingT := new(bkgstruct.BookingT)
	bookingT.SubmissionDateTime = common.TimestampToTime(booking.BookingT.SubmissionDateTime)
	bookingT.ExpectedDepartureDate = common.TimestampToTime(booking.BookingT.ExpectedDepartureDate)
	bookingT.ExpectedArrivalAtPlaceOfDeliveryStartDate = common.TimestampToTime(booking.BookingT.ExpectedArrivalAtPlaceOfDeliveryStartDate)
	bookingT.ExpectedArrivalAtPlaceOfDeliveryEndDate = common.TimestampToTime(booking.BookingT.ExpectedArrivalAtPlaceOfDeliveryEndDate)

	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(booking.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(booking.CrUpdTime.UpdatedAt)
	bookingTmp := bkgstruct.Booking{BookingD: booking.BookingD, BookingT: bookingT, CrUpdUser: booking.CrUpdUser, CrUpdTime: crUpdTime}
	return &bookingTmp, nil
}

// GetBookings - Get Bookings
func (bs *BkgService) GetBookings(ctx context.Context, in *bkgproto.GetBookingsRequest) (*bkgproto.GetBookingsResponse, error) {
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

	bookings := []*bkgproto.Booking{}

	nselectBookingsSQL := selectBookingsSQL + ` where ` + query

	rows, err := bs.DBService.DB.QueryxContext(ctx, nselectBookingsSQL, "active")
	if err != nil {
		bs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	for rows.Next() {

		bookingTmp := bkgstruct.Booking{}
		err = rows.StructScan(&bookingTmp)
		if err != nil {
			bs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		getRequest := commonproto.GetRequest{}
		getRequest.UserEmail = in.UserEmail
		getRequest.RequestId = in.RequestId
		booking, err := bs.getBkgStruct(ctx, &getRequest, bookingTmp)
		if err != nil {
			bs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		bookings = append(bookings, booking)

	}

	bookingsResponse := bkgproto.GetBookingsResponse{}
	if len(bookings) != 0 {
		next := bookings[len(bookings)-1].BookingD.Id
		next--
		nextc := common.EncodeCursor(next)
		bookingsResponse = bkgproto.GetBookingsResponse{Bookings: bookings, NextCursor: nextc}
	} else {
		bookingsResponse = bkgproto.GetBookingsResponse{Bookings: bookings, NextCursor: "0"}
	}
	return &bookingsResponse, nil
}

// GetBooking - Get Booking
func (bs *BkgService) GetBooking(ctx context.Context, getBookingRequest *bkgproto.GetBookingRequest) (*bkgproto.GetBookingResponse, error) {
	in := getBookingRequest.GetRequest
	uuid4byte, err := common.UUIDStrToBytes(in.Id)
	if err != nil {
		bs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	nselectBookingsSQL := selectBookingsSQL + ` where uuid4 = ? and status_code = ?;`
	row := bs.DBService.DB.QueryRowxContext(ctx, nselectBookingsSQL, uuid4byte, "active")
	bookingTmp := bkgstruct.Booking{}
	err = row.StructScan(&bookingTmp)
	if err != nil {
		bs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	booking, err := bs.getBkgStruct(ctx, in, bookingTmp)
	if err != nil {
		bs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	bookingResponse := bkgproto.GetBookingResponse{}
	bookingResponse.Booking = booking
	return &bookingResponse, nil
}

// GetBookingByPk - Get Booking By Primary key(Id)
func (bs *BkgService) GetBookingByPk(ctx context.Context, getBookingByPkRequest *bkgproto.GetBookingByPkRequest) (*bkgproto.GetBookingByPkResponse, error) {
	in := getBookingByPkRequest.GetByIdRequest
	nselectBookingsSQL := selectBookingsSQL + ` where id = ? and status_code = ?;`
	row := bs.DBService.DB.QueryRowxContext(ctx, nselectBookingsSQL, in.Id, "active")
	bookingTmp := bkgstruct.Booking{}
	err := row.StructScan(&bookingTmp)
	if err != nil {
		bs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	getRequest := commonproto.GetRequest{}
	getRequest.UserEmail = in.UserEmail
	getRequest.RequestId = in.RequestId
	booking, err := bs.getBkgStruct(ctx, &getRequest, bookingTmp)
	if err != nil {
		bs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	bookingResponse := bkgproto.GetBookingByPkResponse{}
	bookingResponse.Booking = booking
	return &bookingResponse, nil
}

// GetBookingByCarrierBookingRequestReference - Get BookingByCarrierBookingRequestReference
func (bs *BkgService) GetBookingByCarrierBookingRequestReference(ctx context.Context, in *bkgproto.GetBookingByCarrierBookingRequestReferenceRequest) (*bkgproto.GetBookingByCarrierBookingRequestReferenceResponse, error) {
	nselectBookingsSQL := selectBookingsSQL + ` where carrier_booking_request_reference = ? and status_code = ?;`
	row := bs.DBService.DB.QueryRowxContext(ctx, nselectBookingsSQL, in.CarrierBookingRequestReference, "active")
	bookingTmp := bkgstruct.Booking{}
	err := row.StructScan(&bookingTmp)
	if err != nil {
		bs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	getRequest := commonproto.GetRequest{}
	getRequest.UserEmail = in.UserEmail
	getRequest.RequestId = in.RequestId
	booking, err := bs.getBkgStruct(ctx, &getRequest, bookingTmp)
	if err != nil {
		bs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	bookingResponse := bkgproto.GetBookingByCarrierBookingRequestReferenceResponse{}
	bookingResponse.Booking = booking

	return &bookingResponse, nil
}

// CancelBookingByCarrierBookingReference - CancelBookingByCarrierBookingReference
func (bs *BkgService) CancelBookingByCarrierBookingReference(ctx context.Context, in *bkgproto.CancelBookingByCarrierBookingReferenceRequest) (*bkgproto.CancelBookingByCarrierBookingReferenceResponse, error) {
	return &bkgproto.CancelBookingByCarrierBookingReferenceResponse{}, nil
}

// getBkgStruct - Get getBkgStruct header
func (bs *BkgService) getBkgStruct(ctx context.Context, in *commonproto.GetRequest, bookingTmp bkgstruct.Booking) (*bkgproto.Booking, error) {
	bookingT := new(bkgproto.BookingT)
	bookingT.SubmissionDateTime = common.TimeToTimestamp(bookingTmp.BookingT.SubmissionDateTime)
	bookingT.ExpectedDepartureDate = common.TimeToTimestamp(bookingTmp.BookingT.ExpectedDepartureDate)
	bookingT.ExpectedArrivalAtPlaceOfDeliveryStartDate = common.TimeToTimestamp(bookingTmp.BookingT.ExpectedArrivalAtPlaceOfDeliveryStartDate)
	bookingT.ExpectedArrivalAtPlaceOfDeliveryEndDate = common.TimeToTimestamp(bookingTmp.BookingT.ExpectedArrivalAtPlaceOfDeliveryEndDate)

	crUpdTime := new(commonproto.CrUpdTime)
	crUpdTime.CreatedAt = common.TimeToTimestamp(bookingTmp.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimeToTimestamp(bookingTmp.CrUpdTime.UpdatedAt)

	uuid4Str, err := common.UUIDBytesToStr(bookingTmp.BookingD.Uuid4)
	if err != nil {
		bs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	bookingTmp.BookingD.IdS = uuid4Str

	booking := bkgproto.Booking{BookingD: bookingTmp.BookingD, BookingT: bookingT, CrUpdUser: bookingTmp.CrUpdUser, CrUpdTime: crUpdTime}
	return &booking, nil
}

// UpdateBookingByReferenceCarrierBookingRequestReference - UpdateBookingByReferenceCarrierBookingRequestReference
func (bs *BkgService) UpdateBookingByReferenceCarrierBookingRequestReference(ctx context.Context, in *bkgproto.UpdateBookingByReferenceCarrierBookingRequestReferenceRequest) (*bkgproto.UpdateBookingByReferenceCarrierBookingRequestReferenceResponse, error) {
	db := bs.DBService.DB
	tn := common.GetTimeDetails()
	stmt, err := db.PreparexContext(ctx, updateBookingSQL)
	if err != nil {
		bs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = bs.DBService.InsUpd(ctx, in.GetUserEmail(), in.GetRequestId(), func(tx *sqlx.Tx) error {
		_, err = tx.StmtxContext(ctx, stmt).ExecContext(ctx,
			in.DocumentStatus,
			in.ReceiptTypeAtOrigin,
			in.DeliveryTypeAtDestination,
			in.CargoMovementTypeAtOrigin,
			in.CargoMovementTypeAtDestination,
			in.ServiceContractReference,
			in.PaymentTermCode,
			tn,
			in.CarrierBookingRequestReference)
		if err != nil {
			bs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			err1 := stmt.Close()
			if err1 != nil {
				bs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err1))
				return err1
			}
			return err
		}
		return nil
	})

	if err != nil {
		bs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = stmt.Close()
	if err != nil {
		bs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	return &bkgproto.UpdateBookingByReferenceCarrierBookingRequestReferenceResponse{}, nil
}
