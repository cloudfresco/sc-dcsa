package bkgservices

// https://github.com/dcsaorg/DCSA-BKG/blob/master/src/main/java/org/dcsa/bkg/service/impl/BkgShipmentSummaryServiceImpl.java
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

// BkgShipmentSummaryService - For accessing BkgShipmentSummary services
type BkgShipmentSummaryService struct {
	log               *zap.Logger
	DBService         *common.DBService
	RedisService      *common.RedisService
	UserServiceClient partyproto.UserServiceClient
	bkgproto.UnimplementedBkgShipmentSummaryServiceServer
}

// NewBkgShipmentSummaryService - Create BkgShipmentSummary Service
func NewBkgShipmentSummaryService(log *zap.Logger, dbOpt *common.DBService, redisOpt *common.RedisService, userServiceClient partyproto.UserServiceClient) *BkgShipmentSummaryService {
	return &BkgShipmentSummaryService{
		log:               log,
		DBService:         dbOpt,
		RedisService:      redisOpt,
		UserServiceClient: userServiceClient,
	}
}

// insertBkgShipmentSummarySQL - insert BkgShipmentSummarySQL query
const insertBkgShipmentSummarySQL = `insert into shipment_summaries
	  ( 
  uuid4,
  carrier_booking_reference,
  terms_and_conditions,
  carrier_booking_request_reference,
  booking_document_status,
  shipment_created_date_time,
  shipment_updated_date_time,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at)
  values (:uuid4,
  :carrier_booking_reference,
  :terms_and_conditions,
  :carrier_booking_request_reference,
  :booking_document_status,
  :shipment_created_date_time,
  :shipment_updated_date_time,
  :status_code,
  :created_by_user_id,
  :updated_by_user_id,
  :created_at,
  :updated_at);`

// selectBkgShipmentSummariesSQL - select BkgShipmentSummariesSQL query
const selectBkgShipmentSummariesSQL = `select 
  id,
  uuid4,
  carrier_booking_reference,
  terms_and_conditions,
  carrier_booking_request_reference,
  booking_document_status,
  shipment_created_date_time,
  shipment_updated_date_time,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at from shipment_summaries`

// CreateBkgShipmentSummary - Create BkgShipmentSummary
func (bss *BkgShipmentSummaryService) CreateBkgShipmentSummary(ctx context.Context, in *bkgproto.CreateBkgShipmentSummaryRequest) (*bkgproto.CreateBkgShipmentSummaryResponse, error) {
	user, err := partyservice.GetUserWithNewContext(ctx, in.UserId, in.UserEmail, in.RequestId, bss.UserServiceClient)
	if err != nil {
		bss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	ttime := common.GetTimeDetails()
	tn := common.TimeToTimestamp(ttime)

	shipmentCreatedDateTime, err := time.Parse(common.Layout, in.ShipmentCreatedDateTime)
	if err != nil {
		bss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	shipmentUpdatedDateTime, err := time.Parse(common.Layout, in.ShipmentUpdatedDateTime)
	if err != nil {
		bss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	bkgShipmentSummaryD := bkgproto.BkgShipmentSummaryD{}
	bkgShipmentSummaryD.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		bss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	bkgShipmentSummaryD.CarrierBookingReference = in.CarrierBookingReference
	bkgShipmentSummaryD.TermsAndConditions = in.TermsAndConditions
	bkgShipmentSummaryD.CarrierBookingRequestReference = in.CarrierBookingRequestReference
	bkgShipmentSummaryD.BookingDocumentStatus = in.BookingDocumentStatus

	bkgShipmentSummaryT := bkgproto.BkgShipmentSummaryT{}
	bkgShipmentSummaryT.ShipmentCreatedDateTime = common.TimeToTimestamp(shipmentCreatedDateTime.UTC().Truncate(time.Second))
	bkgShipmentSummaryT.ShipmentUpdatedDateTime = common.TimeToTimestamp(shipmentUpdatedDateTime.UTC().Truncate(time.Second))

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	bkgShipmentSummary := bkgproto.BkgShipmentSummary{BkgShipmentSummaryD: &bkgShipmentSummaryD, BkgShipmentSummaryT: &bkgShipmentSummaryT, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	err = bss.insertBkgShipmentSummary(ctx, insertBkgShipmentSummarySQL, &bkgShipmentSummary, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		bss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	bkgShipmentSummaryResponse := bkgproto.CreateBkgShipmentSummaryResponse{}
	bkgShipmentSummaryResponse.BkgShipmentSummary = &bkgShipmentSummary
	return &bkgShipmentSummaryResponse, nil
}

// insertBkgShipmentSummary - Insert Booking Summary
func (bss *BkgShipmentSummaryService) insertBkgShipmentSummary(ctx context.Context, insertBkgShipmentSummarySQL string, bkgShipmentSummary *bkgproto.BkgShipmentSummary, userEmail string, requestID string) error {
	bkgShipmentSummaryTmp, err := bss.crBkgShipmentSummaryStruct(ctx, bkgShipmentSummary, userEmail, requestID)
	if err != nil {
		bss.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}

	err = bss.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertBkgShipmentSummarySQL, bkgShipmentSummaryTmp)
		if err != nil {
			bss.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			bss.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		bkgShipmentSummary.BkgShipmentSummaryD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(bkgShipmentSummary.BkgShipmentSummaryD.Uuid4)
		if err != nil {
			bss.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		bkgShipmentSummary.BkgShipmentSummaryD.IdS = uuid4Str
		return nil
	})

	if err != nil {
		bss.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// crBkgShipmentSummaryStruct - process BkgShipmentSummary details
func (bss *BkgShipmentSummaryService) crBkgShipmentSummaryStruct(ctx context.Context, bkgShipmentSummary *bkgproto.BkgShipmentSummary, userEmail string, requestID string) (*bkgstruct.BkgShipmentSummary, error) {
	bkgShipmentSummaryT := new(bkgstruct.BkgShipmentSummaryT)
	bkgShipmentSummaryT.ShipmentCreatedDateTime = common.TimestampToTime(bkgShipmentSummary.BkgShipmentSummaryT.ShipmentCreatedDateTime)
	bkgShipmentSummaryT.ShipmentUpdatedDateTime = common.TimestampToTime(bkgShipmentSummary.BkgShipmentSummaryT.ShipmentUpdatedDateTime)

	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(bkgShipmentSummary.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(bkgShipmentSummary.CrUpdTime.UpdatedAt)

	bkgShipmentSummaryTmp := bkgstruct.BkgShipmentSummary{BkgShipmentSummaryD: bkgShipmentSummary.BkgShipmentSummaryD, BkgShipmentSummaryT: bkgShipmentSummaryT, CrUpdUser: bkgShipmentSummary.CrUpdUser, CrUpdTime: crUpdTime}

	return &bkgShipmentSummaryTmp, nil
}

// GetBkgShipmentSummaries - Get BkgShipmentSummaries
func (bss *BkgShipmentSummaryService) GetBkgShipmentSummaries(ctx context.Context, in *bkgproto.GetBkgShipmentSummariesRequest) (*bkgproto.GetBkgShipmentSummariesResponse, error) {
	limit := in.GetLimit()
	nextCursor := in.GetNextCursor()
	if limit == "" {
		limit = bss.DBService.LimitSQLRows
	}
	query := "status_code = ?"
	if nextCursor == "" {
		query = query + " order by id desc " + " limit " + limit + ";"
	} else {
		nextCursor = common.DecodeCursor(nextCursor)
		query = query + " " + "and" + " " + "id <= " + nextCursor + " order by id desc " + " limit " + limit + ";"
	}

	bkgShipSummaries := []*bkgproto.BkgShipmentSummary{}

	nselectBkgShipmentSummariesSQL := selectBkgShipmentSummariesSQL + ` where ` + query

	rows, err := bss.DBService.DB.QueryxContext(ctx, nselectBkgShipmentSummariesSQL, "active")
	if err != nil {
		bss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	for rows.Next() {

		bkgShipmentSummaryTmp := bkgstruct.BkgShipmentSummary{}
		err = rows.StructScan(&bkgShipmentSummaryTmp)
		if err != nil {
			bss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		getRequest := commonproto.GetRequest{}
		getRequest.UserEmail = in.UserEmail
		getRequest.RequestId = in.RequestId
		bkgShipmentSummary, err := bss.getBkgShipmentSummaryStruct(ctx, &getRequest, bkgShipmentSummaryTmp)
		if err != nil {
			bss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		bkgShipSummaries = append(bkgShipSummaries, bkgShipmentSummary)

	}

	bkgShipmentSummaryResponse := bkgproto.GetBkgShipmentSummariesResponse{}
	if len(bkgShipSummaries) != 0 {
		next := bkgShipSummaries[len(bkgShipSummaries)-1].BkgShipmentSummaryD.Id
		next--
		nextc := common.EncodeCursor(next)
		bkgShipmentSummaryResponse = bkgproto.GetBkgShipmentSummariesResponse{BkgShipmentSummaries: bkgShipSummaries, NextCursor: nextc}
	} else {
		bkgShipmentSummaryResponse = bkgproto.GetBkgShipmentSummariesResponse{BkgShipmentSummaries: bkgShipSummaries, NextCursor: "0"}
	}
	return &bkgShipmentSummaryResponse, nil
}

// getBkgShipmentSummaryStruct - Get shipment Summary header
func (bss *BkgShipmentSummaryService) getBkgShipmentSummaryStruct(ctx context.Context, in *commonproto.GetRequest, bkgShipmentSummaryTmp bkgstruct.BkgShipmentSummary) (*bkgproto.BkgShipmentSummary, error) {
	bkgShipmentSummaryT := new(bkgproto.BkgShipmentSummaryT)
	bkgShipmentSummaryT.ShipmentCreatedDateTime = common.TimeToTimestamp(bkgShipmentSummaryTmp.BkgShipmentSummaryT.ShipmentCreatedDateTime)
	bkgShipmentSummaryT.ShipmentUpdatedDateTime = common.TimeToTimestamp(bkgShipmentSummaryTmp.BkgShipmentSummaryT.ShipmentUpdatedDateTime)

	crUpdTime := new(commonproto.CrUpdTime)
	crUpdTime.CreatedAt = common.TimeToTimestamp(bkgShipmentSummaryTmp.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimeToTimestamp(bkgShipmentSummaryTmp.CrUpdTime.UpdatedAt)

	uuid4Str, err := common.UUIDBytesToStr(bkgShipmentSummaryTmp.BkgShipmentSummaryD.Uuid4)
	if err != nil {
		bss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	bkgShipmentSummaryTmp.BkgShipmentSummaryD.IdS = uuid4Str

	bkgShipmentSummary := bkgproto.BkgShipmentSummary{BkgShipmentSummaryD: bkgShipmentSummaryTmp.BkgShipmentSummaryD, BkgShipmentSummaryT: bkgShipmentSummaryT, CrUpdUser: bkgShipmentSummaryTmp.CrUpdUser, CrUpdTime: crUpdTime}

	return &bkgShipmentSummary, nil
}
