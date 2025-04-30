package eblservices

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

// IssueRequestResponseService - For accessing Issuance services
type IssueRequestResponseService struct {
	log               *zap.Logger
	DBService         *common.DBService
	RedisService      *common.RedisService
	UserServiceClient partyproto.UserServiceClient
	eblproto.UnimplementedIssueRequestResponseServiceServer
}

// NewIssueRequestResponseService - Create Shipping service
func NewIssueRequestResponseService(log *zap.Logger, dbOpt *common.DBService, redisOpt *common.RedisService, userServiceClient partyproto.UserServiceClient) *IssueRequestResponseService {
	return &IssueRequestResponseService{
		log:               log,
		DBService:         dbOpt,
		RedisService:      redisOpt,
		UserServiceClient: userServiceClient,
	}
}

// insertIssuanceRequestResponseSQL - Insert IssuanceRequestResponseSQL Query
const insertIssuanceRequestResponseSQL = `insert into issuance_request_responses
	  ( 
  uuid4,
  transport_document_reference,
  issuance_response_code,
  reason,
  issuance_request_id,
  created_date_time,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at
  )
  values (:uuid4,
:transport_document_reference,
:issuance_response_code,
:reason,
:issuance_request_id,
:created_date_time,
:status_code,
:created_by_user_id,
:updated_by_user_id,
:created_at,
:updated_at);`

// updateIssuanceRequestResponseSQL - update IssuanceRequestResponseSQL query
const updateIssuanceRequestResponseSQL = `update issuance_request_responses set 
  transport_document_reference = ?,
  issuance_response_code = ?,
  reason = ?,
  updated_at = ? where uuid4 = ?;`

// CreateIssuanceRequestResponse - Create  IssuanceRequestResponse
func (is *IssueRequestResponseService) CreateIssuanceRequestResponse(ctx context.Context, in *eblproto.CreateIssuanceRequestResponseRequest) (*eblproto.CreateIssuanceRequestResponseResponse, error) {
	user, err := partyservice.GetUserWithNewContext(ctx, in.UserId, in.UserEmail, in.RequestId, is.UserServiceClient)
	if err != nil {
		is.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	ttime := common.GetTimeDetails()
	tn := common.TimeToTimestamp(ttime)

	createdDateTime, err := time.Parse(common.Layout, in.CreatedDateTime)
	if err != nil {
		is.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	issuanceRequestResponseD := eblproto.IssuanceRequestResponseD{}
	issuanceRequestResponseD.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		is.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	issuanceRequestResponseD.TransportDocumentReference = in.TransportDocumentReference
	issuanceRequestResponseD.IssuanceResponseCode = in.IssuanceResponseCode
	issuanceRequestResponseD.Reason = in.Reason
	issuanceRequestResponseD.IssuanceRequestId = in.IssuanceRequestId

	issuanceRequestResponseT := eblproto.IssuanceRequestResponseT{}
	issuanceRequestResponseT.CreatedDateTime = common.TimeToTimestamp(createdDateTime.UTC().Truncate(time.Second))

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	issuanceRequestResponse := eblproto.IssuanceRequestResponse{IssuanceRequestResponseD: &issuanceRequestResponseD, IssuanceRequestResponseT: &issuanceRequestResponseT, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	err = is.insertIssuanceRequestResponse(ctx, insertIssuanceRequestResponseSQL, &issuanceRequestResponse, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		is.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	issuanceRequestResponseResponse := eblproto.CreateIssuanceRequestResponseResponse{}
	issuanceRequestResponseResponse.IssuanceRequestResponse = &issuanceRequestResponse
	return &issuanceRequestResponseResponse, nil
}

// insertIssuanceRequestResponse - Insert IssuanceRequestResponse
func (is *IssueRequestResponseService) insertIssuanceRequestResponse(ctx context.Context, insertIssuanceRequestResponseSQL string, issuanceRequestResponse *eblproto.IssuanceRequestResponse, userEmail string, requestID string) error {
	issuanceRequestResponseTmp, err := is.CrIssuanceRequestResponseStruct(ctx, issuanceRequestResponse, userEmail, requestID)
	if err != nil {
		is.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}

	err = is.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertIssuanceRequestResponseSQL, issuanceRequestResponseTmp)
		if err != nil {
			is.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			is.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		issuanceRequestResponse.IssuanceRequestResponseD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(issuanceRequestResponse.IssuanceRequestResponseD.Uuid4)
		if err != nil {
			is.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		issuanceRequestResponse.IssuanceRequestResponseD.IdS = uuid4Str
		return nil
	})
	if err != nil {
		is.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// CrIssuanceRequestResponseStruct - process IssuanceRequestResponse details
func (is *IssueRequestResponseService) CrIssuanceRequestResponseStruct(ctx context.Context, issuanceRequestResponse *eblproto.IssuanceRequestResponse, userEmail string, requestID string) (*eblstruct.IssuanceRequestResponse, error) {
	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(issuanceRequestResponse.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(issuanceRequestResponse.CrUpdTime.UpdatedAt)

	issuanceRequestResponseT := new(eblstruct.IssuanceRequestResponseT)
	issuanceRequestResponseT.CreatedDateTime = common.TimestampToTime(issuanceRequestResponse.IssuanceRequestResponseT.CreatedDateTime)

	issuanceRequestResponseTmp := eblstruct.IssuanceRequestResponse{IssuanceRequestResponseD: issuanceRequestResponse.IssuanceRequestResponseD, IssuanceRequestResponseT: issuanceRequestResponseT, CrUpdUser: issuanceRequestResponse.CrUpdUser, CrUpdTime: crUpdTime}

	return &issuanceRequestResponseTmp, nil
}

// UpdateIssuanceRequestResponse - Update IssuanceRequestResponse
func (is *IssueRequestResponseService) UpdateIssuanceRequestResponse(ctx context.Context, in *eblproto.UpdateIssuanceRequestResponseRequest) (*eblproto.UpdateIssuanceRequestResponseResponse, error) {
	db := is.DBService.DB
	tn := common.GetTimeDetails()

	uuid4byte, err := common.UUIDStrToBytes(in.Id)
	if err != nil {
		is.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	stmt, err := db.PreparexContext(ctx, updateIssuanceRequestResponseSQL)
	if err != nil {
		is.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = is.DBService.InsUpd(ctx, in.GetUserEmail(), in.GetRequestId(), func(tx *sqlx.Tx) error {
		_, err = tx.StmtxContext(ctx, stmt).ExecContext(ctx,
			in.TransportDocumentReference,
			in.IssuanceResponseCode,
			in.Reason,
			tn,
			uuid4byte)
		if err != nil {
			is.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			err1 := stmt.Close()
			if err1 != nil {
				is.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err1))
				return err1
			}
			return err
		}
		return nil
	})

	if err != nil {
		is.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = stmt.Close()
	if err != nil {
		is.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	return &eblproto.UpdateIssuanceRequestResponseResponse{}, nil
}
