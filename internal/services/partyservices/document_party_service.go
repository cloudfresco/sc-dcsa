package partyservices

import (
	"context"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	commonproto "github.com/cloudfresco/sc-dcsa/internal/protogen/common/v1"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"
	commonstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/common/v1"
	partystruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/party/v1"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// DocumentPartyService - For accessing Document Party services
type DocumentPartyService struct {
	log               *zap.Logger
	DBService         *common.DBService
	RedisService      *common.RedisService
	UserServiceClient partyproto.UserServiceClient
	partyproto.UnimplementedDocumentPartyServiceServer
}

// NewDocumentPartyService - Create Document Party service
func NewDocumentPartyService(log *zap.Logger, dbOpt *common.DBService, redisOpt *common.RedisService, userServiceClient partyproto.UserServiceClient) *DocumentPartyService {
	return &DocumentPartyService{
		log:               log,
		DBService:         dbOpt,
		RedisService:      redisOpt,
		UserServiceClient: userServiceClient,
	}
}

// insertDocumentPartySQL - insert DocumentPartySQL query
const insertDocumentPartySQL = `insert into document_parties
	  ( 
  uuid4,
  party_id,
  shipping_instruction_id,
  shipment_id,
  party_function,
  is_to_be_notified,
  booking_id, 
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at
  )
  values (:uuid4,
:party_id,
:shipping_instruction_id,
:shipment_id,
:party_function,
:is_to_be_notified,
:booking_id,
:status_code,
    :created_by_user_id,
    :updated_by_user_id,
    :created_at,
    :updated_at);`

// selectDocumentPartiesSQL - select DocumentPartiesSQL query
const selectDocumentPartiesSQL = `select 
  id,
  uuid4,
  party_id,
  shipping_instruction_id,
  shipment_id,
  party_function,
  is_to_be_notified,
  booking_id,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at from document_parties`

// CreateDocumentParty - CreateDocumentParty
func (dps *DocumentPartyService) CreateDocumentParty(ctx context.Context, in *partyproto.CreateDocumentPartyRequest) (*partyproto.CreateDocumentPartyResponse, error) {
	documentParty, err := dps.ProcessDocumentPartyRequest(ctx, in)
	if err != nil {
		dps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = dps.insertDocumentParty(ctx, insertDocumentPartySQL, documentParty, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		dps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	documentPartyResponse := partyproto.CreateDocumentPartyResponse{}
	documentPartyResponse.DocumentParty = documentParty
	return &documentPartyResponse, nil
}

// ProcessDocumentPartyRequest - Process DocumentPartyRequest
func (dps *DocumentPartyService) ProcessDocumentPartyRequest(ctx context.Context, in *partyproto.CreateDocumentPartyRequest) (*partyproto.DocumentParty, error) {
	user, err := GetUserWithNewContext(ctx, in.UserId, in.UserEmail, in.RequestId, dps.UserServiceClient)
	if err != nil {
		dps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	ttime := common.GetTimeDetails()
	tn := common.TimeToTimestamp(ttime)

	documentPartyD := partyproto.DocumentPartyD{}
	documentPartyD.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		dps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	documentPartyD.PartyId = in.PartyId
	documentPartyD.ShippingInstructionId = in.ShippingInstructionId
	documentPartyD.ShipmentId = in.ShipmentId
	documentPartyD.PartyFunction = in.PartyFunction
	documentPartyD.IsToBeNotified = in.IsToBeNotified
	documentPartyD.BookingId = in.BookingId

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	documentParty := partyproto.DocumentParty{DocumentPartyD: &documentPartyD, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	return &documentParty, nil
}

// insertDocumentParty - Insert Document Party
func (dps *DocumentPartyService) insertDocumentParty(ctx context.Context, insertDocumentPartySQL string, documentParty *partyproto.DocumentParty, userEmail string, requestID string) error {
	documentPartyTmp, err := dps.crDocumentPartyStruct(ctx, documentParty, userEmail, requestID)
	if err != nil {
		dps.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	err = dps.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertDocumentPartySQL, documentPartyTmp)
		if err != nil {
			dps.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			dps.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		documentParty.DocumentPartyD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(documentParty.DocumentPartyD.Uuid4)
		if err != nil {
			dps.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		documentParty.DocumentPartyD.IdS = uuid4Str
		return nil
	})

	if err != nil {
		dps.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// crDocumentPartyStruct - process DocumentParty details
func (dps *DocumentPartyService) crDocumentPartyStruct(ctx context.Context, documentParty *partyproto.DocumentParty, userEmail string, requestID string) (*partystruct.DocumentParty, error) {
	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(documentParty.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(documentParty.CrUpdTime.UpdatedAt)

	documentPartyTmp := partystruct.DocumentParty{DocumentPartyD: documentParty.DocumentPartyD, CrUpdUser: documentParty.CrUpdUser, CrUpdTime: crUpdTime}

	return &documentPartyTmp, nil
}

// CreateDocumentPartiesByBookingID - CreateDocumentPartiesByBookingID
func (dps *DocumentPartyService) CreateDocumentPartiesByBookingID(ctx context.Context, inReq *partyproto.CreateDocumentPartiesByBookingIDRequest) (*partyproto.CreateDocumentPartiesByBookingIDResponse, error) {
	in := inReq.CreateDocumentPartyRequest
	documentParty, err := dps.ProcessDocumentPartyRequest(ctx, in)
	if err != nil {
		dps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = dps.insertDocumentParty(ctx, insertDocumentPartySQL, documentParty, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		dps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	documentPartyResponse := partyproto.CreateDocumentPartiesByBookingIDResponse{}
	documentPartyResponse.DocumentParty = documentParty
	return &documentPartyResponse, nil
}

// CreateDocumentPartiesByShippingInstructionID - CreateDocumentPartiesByShippingInstructionID
func (dps *DocumentPartyService) CreateDocumentPartiesByShippingInstructionID(ctx context.Context, inReq *partyproto.CreateDocumentPartiesByShippingInstructionIDRequest) (*partyproto.CreateDocumentPartiesByShippingInstructionIDResponse, error) {
	in := inReq.CreateDocumentPartyRequest
	documentParty, err := dps.ProcessDocumentPartyRequest(ctx, in)
	if err != nil {
		dps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = dps.insertDocumentParty(ctx, insertDocumentPartySQL, documentParty, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		dps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	documentPartyResponse := partyproto.CreateDocumentPartiesByShippingInstructionIDResponse{}
	documentPartyResponse.DocumentParty = documentParty
	return &documentPartyResponse, nil
}

// FetchDocumentPartiesByBookingID - Get DocumentParties
func (dps *DocumentPartyService) FetchDocumentPartiesByBookingID(ctx context.Context, in *partyproto.FetchDocumentPartiesByBookingIDRequest) (*partyproto.FetchDocumentPartiesByBookingIDResponse, error) {
	limit := in.GetLimit()
	nextCursor := in.GetNextCursor()
	if limit == "" {
		limit = dps.DBService.LimitSQLRows
	}
	query := "booking_id = ?"
	if nextCursor == "" {
		query = query + " order by id desc " + " limit " + limit + ";"
	} else {
		nextCursor = common.DecodeCursor(nextCursor)
		query = query + " " + "and" + " " + "id <= " + nextCursor + " order by id desc " + " limit " + limit + ";"
	}

	documentParties := []*partyproto.DocumentParty{}

	nselectDocumentPartiesSQL := selectDocumentPartiesSQL + ` where ` + query

	rows, err := dps.DBService.DB.QueryxContext(ctx, nselectDocumentPartiesSQL, in.BookingId)
	if err != nil {
		dps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	for rows.Next() {

		documentPartyTmp := partystruct.DocumentParty{}
		err = rows.StructScan(&documentPartyTmp)
		if err != nil {
			dps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		in1 := commonproto.GetRequest{}
		in1.UserEmail = in.UserEmail
		in1.RequestId = in.RequestId
		documentParty, err := dps.getDocumentPartyStruct(ctx, &in1, documentPartyTmp)
		if err != nil {
			dps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		documentParties = append(documentParties, documentParty)

	}

	documentPartyResponse := partyproto.FetchDocumentPartiesByBookingIDResponse{}
	if len(documentParties) != 0 {
		next := documentParties[len(documentParties)-1].DocumentPartyD.Id
		next--
		nextc := common.EncodeCursor(next)
		documentPartyResponse = partyproto.FetchDocumentPartiesByBookingIDResponse{DocumentParties: documentParties, NextCursor: nextc}
	} else {
		documentPartyResponse = partyproto.FetchDocumentPartiesByBookingIDResponse{DocumentParties: documentParties, NextCursor: "0"}
	}
	return &documentPartyResponse, nil
}

// FetchDocumentPartiesByByShippingInstructionID - Get DocumentParties
func (dps *DocumentPartyService) FetchDocumentPartiesByByShippingInstructionID(ctx context.Context, in *partyproto.FetchDocumentPartiesByByShippingInstructionIDRequest) (*partyproto.FetchDocumentPartiesByByShippingInstructionIDResponse, error) {
	limit := in.GetLimit()
	nextCursor := in.GetNextCursor()
	if limit == "" {
		limit = dps.DBService.LimitSQLRows
	}
	query := "shipping_instruction_id = ?"
	if nextCursor == "" {
		query = query + " order by id desc " + " limit " + limit + ";"
	} else {
		nextCursor = common.DecodeCursor(nextCursor)
		query = query + " " + "and" + " " + "id <= " + nextCursor + " order by id desc " + " limit " + limit + ";"
	}

	documentParties := []*partyproto.DocumentParty{}

	nselectDocumentPartiesSQL := selectDocumentPartiesSQL + ` where ` + query

	rows, err := dps.DBService.DB.QueryxContext(ctx, nselectDocumentPartiesSQL, in.ShippingInstructionId)
	if err != nil {
		dps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	for rows.Next() {

		documentPartyTmp := partystruct.DocumentParty{}
		err = rows.StructScan(&documentPartyTmp)
		if err != nil {
			dps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		in1 := commonproto.GetRequest{}
		in1.UserEmail = in.UserEmail
		in1.RequestId = in.RequestId
		documentParty, err := dps.getDocumentPartyStruct(ctx, &in1, documentPartyTmp)
		if err != nil {
			dps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		documentParties = append(documentParties, documentParty)

	}

	documentPartyResponse := partyproto.FetchDocumentPartiesByByShippingInstructionIDResponse{}
	if len(documentParties) != 0 {
		next := documentParties[len(documentParties)-1].DocumentPartyD.Id
		next--
		nextc := common.EncodeCursor(next)
		documentPartyResponse = partyproto.FetchDocumentPartiesByByShippingInstructionIDResponse{DocumentParties: documentParties, NextCursor: nextc}
	} else {
		documentPartyResponse = partyproto.FetchDocumentPartiesByByShippingInstructionIDResponse{DocumentParties: documentParties, NextCursor: "0"}
	}
	return &documentPartyResponse, nil
}

// getDocumentPartyStruct - Get Document party
func (dps *DocumentPartyService) getDocumentPartyStruct(ctx context.Context, in *commonproto.GetRequest, documentPartyTmp partystruct.DocumentParty) (*partyproto.DocumentParty, error) {
	uuid4Str, err := common.UUIDBytesToStr(documentPartyTmp.DocumentPartyD.Uuid4)
	if err != nil {
		dps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	documentPartyTmp.DocumentPartyD.IdS = uuid4Str

	crUpdTime := new(commonproto.CrUpdTime)
	crUpdTime.CreatedAt = common.TimeToTimestamp(documentPartyTmp.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimeToTimestamp(documentPartyTmp.CrUpdTime.UpdatedAt)

	documentParty := partyproto.DocumentParty{DocumentPartyD: documentPartyTmp.DocumentPartyD, CrUpdUser: documentPartyTmp.CrUpdUser, CrUpdTime: crUpdTime}

	return &documentParty, nil
}

// ResolveDocumentPartiesForShippingInstructionID - ResolveDocumentPartiesForShippingInstructionID
func (dps *DocumentPartyService) ResolveDocumentPartiesForShippingInstructionID(ctx context.Context, in *partyproto.ResolveDocumentPartiesForShippingInstructionIDRequest) (*partyproto.ResolveDocumentPartiesForShippingInstructionIDResponse, error) {
	return &partyproto.ResolveDocumentPartiesForShippingInstructionIDResponse{}, nil
}
