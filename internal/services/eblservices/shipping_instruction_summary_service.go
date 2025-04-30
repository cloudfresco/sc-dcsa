package eblservices

// https://github.com/dcsaorg/DCSA-EBL/blob/master/src/main/java/org/dcsa/ebl/service/ShippingInstructionService.java
import (
	"context"
	"time"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	commonproto "github.com/cloudfresco/sc-dcsa/internal/protogen/common/v1"
	eblproto "github.com/cloudfresco/sc-dcsa/internal/protogen/ebl/v1"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"
	partyservice "github.com/cloudfresco/sc-dcsa/internal/services/partyservices"
	commonstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/common/v1"
	eblstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/ebl/v1"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// InsertShippingInstructionSummarySQL - Insert ShippingInstructionSummarySQL Query
const insertShippingInstructionSummarySQL = `insert into shipping_instruction_summaries
	  ( 
  uuid4,
  shipping_instruction_reference,
  document_status,
  amend_to_transport_document,
  transport_document_type_code,
  is_shipped_onboard_type,
  number_of_copies,
  number_of_originals,
  is_electronic,
  is_to_order,
  are_charges_displayed_on_originals,
  are_charges_displayed_on_copies,
  displayed_name_for_place_of_receipt,
  displayed_name_for_port_of_load,
  displayed_name_for_port_of_discharge,
  displayed_name_for_place_of_delivery,
  shipping_instruction_created_date_time,
  shipping_instruction_updated_date_time,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at)
  values (:uuid4,
  :shipping_instruction_reference,
  :document_status,
  :amend_to_transport_document,
  :transport_document_type_code,
  :is_shipped_onboard_type,
  :number_of_copies,
  :number_of_originals,
  :is_electronic,
  :is_to_order,
  :are_charges_displayed_on_originals,
  :are_charges_displayed_on_copies,
  :displayed_name_for_place_of_receipt,
  :displayed_name_for_port_of_load,
  :displayed_name_for_port_of_discharge,
  :displayed_name_for_place_of_delivery,
  :shipping_instruction_created_date_time,
  :shipping_instruction_updated_date_time,
:status_code,
:created_by_user_id,
:updated_by_user_id,
:created_at,
:updated_at);`

// selectShippingInstructionSummarySQL - select ShippingInstructionSummarySQL Query
const selectShippingInstructionSummariesSQL = `select 
  id,
  uuid4,
  shipping_instruction_reference,
  document_status,
  amend_to_transport_document,
  transport_document_type_code,
  is_shipped_onboard_type,
  number_of_copies,
  number_of_originals,
  is_electronic,
  is_to_order,
  are_charges_displayed_on_originals,
  are_charges_displayed_on_copies,
  displayed_name_for_place_of_receipt,
  displayed_name_for_port_of_load,
  displayed_name_for_port_of_discharge,
  displayed_name_for_place_of_delivery,  
  shipping_instruction_created_date_time,
  shipping_instruction_updated_date_time,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at from shipping_instruction_summaries`

// ShippingInstructionSummaryService - For accessing Shipping Instruction Summary services
type ShippingInstructionSummaryService struct {
	log               *zap.Logger
	DBService         *common.DBService
	RedisService      *common.RedisService
	UserServiceClient partyproto.UserServiceClient
	eblproto.UnimplementedShippingInstructionSummaryServiceServer
}

// NewShippingInstructionSummaryService - Create Shipping Instruction Summary service
func NewShippingInstructionSummaryService(log *zap.Logger, dbOpt *common.DBService, redisOpt *common.RedisService, userServiceClient partyproto.UserServiceClient) *ShippingInstructionSummaryService {
	return &ShippingInstructionSummaryService{
		log:               log,
		DBService:         dbOpt,
		RedisService:      redisOpt,
		UserServiceClient: userServiceClient,
	}
}

// CreateShippingInstructionSummary - Create ShippingInstructionSummary
func (sis *ShippingInstructionSummaryService) CreateShippingInstructionSummary(ctx context.Context, in *eblproto.CreateShippingInstructionSummaryRequest) (*eblproto.CreateShippingInstructionSummaryResponse, error) {
	user, err := partyservice.GetUserWithNewContext(ctx, in.UserId, in.UserEmail, in.RequestId, sis.UserServiceClient)
	if err != nil {
		sis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	ttime := common.GetTimeDetails()
	tn := common.TimeToTimestamp(ttime)

	shippingInstructionCreatedDateTime, err := time.Parse(common.Layout, in.ShippingInstructionCreatedDateTime)
	if err != nil {
		sis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	shippingInstructionUpdatedDateTime, err := time.Parse(common.Layout, in.ShippingInstructionUpdatedDateTime)
	if err != nil {
		sis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	shippingInstructionSummaryD := eblproto.ShippingInstructionSummaryD{}
	shippingInstructionSummaryD.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		sis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	shippingInstructionSummaryD.ShippingInstructionReference = in.ShippingInstructionReference
	shippingInstructionSummaryD.DocumentStatus = in.DocumentStatus
	shippingInstructionSummaryD.AmendToTransportDocument = in.AmendToTransportDocument
	shippingInstructionSummaryD.TransportDocumentTypeCode = in.TransportDocumentTypeCode
	shippingInstructionSummaryD.IsShippedOnboardType = in.IsShippedOnboardType
	shippingInstructionSummaryD.NumberOfCopies = in.NumberOfCopies
	shippingInstructionSummaryD.NumberOfOriginals = in.NumberOfOriginals
	shippingInstructionSummaryD.IsElectronic = in.IsElectronic
	shippingInstructionSummaryD.IsToOrder = in.IsToOrder
	shippingInstructionSummaryD.AreChargesDisplayedOnOriginals = in.AreChargesDisplayedOnOriginals
	shippingInstructionSummaryD.AreChargesDisplayedOnCopies = in.AreChargesDisplayedOnCopies
	shippingInstructionSummaryD.DisplayedNameForPlaceOfReceipt = in.DisplayedNameForPlaceOfReceipt
	shippingInstructionSummaryD.DisplayedNameForPortOfLoad = in.DisplayedNameForPortOfLoad
	shippingInstructionSummaryD.DisplayedNameForPortOfDischarge = in.DisplayedNameForPortOfDischarge
	shippingInstructionSummaryD.DisplayedNameForPlaceOfDelivery = in.DisplayedNameForPlaceOfDelivery

	shippingInstructionSummaryT := eblproto.ShippingInstructionSummaryT{}
	shippingInstructionSummaryT.ShippingInstructionCreatedDateTime = common.TimeToTimestamp(shippingInstructionCreatedDateTime.UTC().Truncate(time.Second))
	shippingInstructionSummaryT.ShippingInstructionUpdatedDateTime = common.TimeToTimestamp(shippingInstructionUpdatedDateTime.UTC().Truncate(time.Second))

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	shippingInstructionSummary := eblproto.ShippingInstructionSummary{ShippingInstructionSummaryD: &shippingInstructionSummaryD, ShippingInstructionSummaryT: &shippingInstructionSummaryT, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	err = sis.insertShippingInstructionSummary(ctx, insertShippingInstructionSummarySQL, &shippingInstructionSummary, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		sis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	shipInstSumResponse := eblproto.CreateShippingInstructionSummaryResponse{}
	shipInstSumResponse.ShippingInstructionSummary = &shippingInstructionSummary
	return &shipInstSumResponse, nil
}

// insertShippingInstructionSummary - Insert ShippingInstructionSummary
func (sis *ShippingInstructionSummaryService) insertShippingInstructionSummary(ctx context.Context, insertShippingInstructionSummarySQL string, shippingInstructionSummary *eblproto.ShippingInstructionSummary, userEmail string, requestID string) error {
	shippingInstructionSummaryTmp, err := sis.crShippingInstructionSummaryStruct(ctx, shippingInstructionSummary, userEmail, requestID)
	if err != nil {
		sis.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}

	err = sis.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertShippingInstructionSummarySQL, shippingInstructionSummaryTmp)
		if err != nil {
			sis.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			sis.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		shippingInstructionSummary.ShippingInstructionSummaryD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(shippingInstructionSummary.ShippingInstructionSummaryD.Uuid4)
		if err != nil {
			sis.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		shippingInstructionSummary.ShippingInstructionSummaryD.IdS = uuid4Str

		return nil
	})

	if err != nil {
		sis.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// crShippingInstructionSummaryStruct - process ShippingInstructionSummary details
func (sis *ShippingInstructionSummaryService) crShippingInstructionSummaryStruct(ctx context.Context, shippingInstructionSummary *eblproto.ShippingInstructionSummary, userEmail string, requestID string) (*eblstruct.ShippingInstructionSummary, error) {
	shippingInstructionSummaryT := new(eblstruct.ShippingInstructionSummaryT)
	shippingInstructionSummaryT.ShippingInstructionCreatedDateTime = common.TimestampToTime(shippingInstructionSummary.ShippingInstructionSummaryT.ShippingInstructionCreatedDateTime)
	shippingInstructionSummaryT.ShippingInstructionUpdatedDateTime = common.TimestampToTime(shippingInstructionSummary.ShippingInstructionSummaryT.ShippingInstructionUpdatedDateTime)

	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(shippingInstructionSummary.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(shippingInstructionSummary.CrUpdTime.UpdatedAt)

	shippingInstructionSummaryTmp := eblstruct.ShippingInstructionSummary{ShippingInstructionSummaryD: shippingInstructionSummary.ShippingInstructionSummaryD, ShippingInstructionSummaryT: shippingInstructionSummaryT, CrUpdUser: shippingInstructionSummary.CrUpdUser, CrUpdTime: crUpdTime}

	return &shippingInstructionSummaryTmp, nil
}

// GetShippingInstructionSummaries - Get ShippingInstructionSummaries
func (sis *ShippingInstructionSummaryService) GetShippingInstructionSummaries(ctx context.Context, in *eblproto.GetShippingInstructionSummariesRequest) (*eblproto.GetShippingInstructionSummariesResponse, error) {
	limit := in.GetLimit()
	nextCursor := in.GetNextCursor()
	if limit == "" {
		limit = sis.DBService.LimitSQLRows
	}
	query := ""
	if nextCursor == "" {
		query = " order by id desc " + " limit " + limit + ";"
	} else {
		nextCursor = common.DecodeCursor(nextCursor)
		query = " where id <= " + nextCursor + " order by id desc " + " limit " + limit + ";"
	}

	shippingInstructionSummaries := []*eblproto.ShippingInstructionSummary{}

	nselectShippingInstructionSummariesSQL := selectShippingInstructionSummariesSQL + query

	rows, err := sis.DBService.DB.QueryxContext(ctx, nselectShippingInstructionSummariesSQL)
	if err != nil {
		sis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	for rows.Next() {

		shippingInstructionSummaryTmp := eblstruct.ShippingInstructionSummary{}
		err = rows.StructScan(&shippingInstructionSummaryTmp)
		if err != nil {
			sis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		getRequest := commonproto.GetRequest{}
		getRequest.UserEmail = in.UserEmail
		getRequest.RequestId = in.RequestId
		shippingInstructionSummary, err := sis.getShippingInstructionSummaryStruct(ctx, &getRequest, shippingInstructionSummaryTmp)
		if err != nil {
			sis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		shippingInstructionSummaries = append(shippingInstructionSummaries, shippingInstructionSummary)

	}

	shipInstSumsResponse := eblproto.GetShippingInstructionSummariesResponse{}
	if len(shippingInstructionSummaries) != 0 {
		next := shippingInstructionSummaries[len(shippingInstructionSummaries)-1].ShippingInstructionSummaryD.Id
		next--
		nextc := common.EncodeCursor(next)
		shipInstSumsResponse = eblproto.GetShippingInstructionSummariesResponse{ShippingInstructionSummaries: shippingInstructionSummaries, NextCursor: nextc}
	} else {
		shipInstSumsResponse = eblproto.GetShippingInstructionSummariesResponse{ShippingInstructionSummaries: shippingInstructionSummaries, NextCursor: "0"}
	}
	return &shipInstSumsResponse, nil
}

// getShippingInstructionSummaryStruct - Get shippingInstructionSummary header
func (sis *ShippingInstructionSummaryService) getShippingInstructionSummaryStruct(ctx context.Context, in *commonproto.GetRequest, shippingInstructionSummaryTmp eblstruct.ShippingInstructionSummary) (*eblproto.ShippingInstructionSummary, error) {
	shippingInstructionSummaryT := new(eblproto.ShippingInstructionSummaryT)
	shippingInstructionSummaryT.ShippingInstructionCreatedDateTime = common.TimeToTimestamp(shippingInstructionSummaryTmp.ShippingInstructionSummaryT.ShippingInstructionCreatedDateTime)
	shippingInstructionSummaryT.ShippingInstructionUpdatedDateTime = common.TimeToTimestamp(shippingInstructionSummaryTmp.ShippingInstructionSummaryT.ShippingInstructionUpdatedDateTime)

	crUpdTime := new(commonproto.CrUpdTime)
	crUpdTime.CreatedAt = common.TimeToTimestamp(shippingInstructionSummaryTmp.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimeToTimestamp(shippingInstructionSummaryTmp.CrUpdTime.UpdatedAt)

	uuid4Str, err := common.UUIDBytesToStr(shippingInstructionSummaryTmp.ShippingInstructionSummaryD.Uuid4)
	if err != nil {
		sis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	shippingInstructionSummaryTmp.ShippingInstructionSummaryD.IdS = uuid4Str

	shippingInstructionSummary := eblproto.ShippingInstructionSummary{ShippingInstructionSummaryD: shippingInstructionSummaryTmp.ShippingInstructionSummaryD, ShippingInstructionSummaryT: shippingInstructionSummaryT, CrUpdUser: shippingInstructionSummaryTmp.CrUpdUser, CrUpdTime: crUpdTime}

	return &shippingInstructionSummary, nil
}
