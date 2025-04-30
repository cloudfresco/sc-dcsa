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

// SurrenderRequestAnswerService - For accessing Surrender services
type SurrenderRequestAnswerService struct {
	log               *zap.Logger
	DBService         *common.DBService
	RedisService      *common.RedisService
	UserServiceClient partyproto.UserServiceClient
	eblproto.UnimplementedSurrenderRequestAnswerServiceServer
}

// NewSurrenderRequestAnswerService - Create Shipping service
func NewSurrenderRequestAnswerService(log *zap.Logger, dbOpt *common.DBService, redisOpt *common.RedisService, userServiceClient partyproto.UserServiceClient) *SurrenderRequestAnswerService {
	return &SurrenderRequestAnswerService{
		log:               log,
		DBService:         dbOpt,
		RedisService:      redisOpt,
		UserServiceClient: userServiceClient,
	}
}

// insertSurrenderRequestAnswerSQL - Insert SurrenderRequestAnswerSQL Query
const insertSurrenderRequestAnswerSQL = `insert into surrender_request_answers
	  ( 
  uuid4,
  surrender_request_reference,
  action,
  comments,
  created_date_time,
  surrender_request_id,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at
  )
  values (:uuid4,
:surrender_request_reference,
:action,
:comments,
:created_date_time,
:surrender_request_id,
:status_code,
:created_by_user_id,
:updated_by_user_id,
:created_at,
:updated_at);`

// updateSurrenderRequestAnswerSQL - update SurrenderRequestAnswerSQL query
const updateSurrenderRequestAnswerSQL = `update surrender_request_answers set 
  surrender_request_reference = ?,
  action = ?,
  comments = ?,
  updated_at = ? where uuid4 = ?;`

// CreateSurrenderRequestAnswer - Create  SurrenderRequestAnswer
func (ss *SurrenderRequestAnswerService) CreateSurrenderRequestAnswer(ctx context.Context, in *eblproto.CreateSurrenderRequestAnswerRequest) (*eblproto.CreateSurrenderRequestAnswerResponse, error) {
	user, err := partyservice.GetUserWithNewContext(ctx, in.UserId, in.UserEmail, in.RequestId, ss.UserServiceClient)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	ttime := common.GetTimeDetails()
	tn := common.TimeToTimestamp(ttime)

	createdDateTime, err := time.Parse(common.Layout, in.CreatedDateTime)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	surrenderRequestAnswerD := eblproto.SurrenderRequestAnswerD{}
	surrenderRequestAnswerD.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	surrenderRequestAnswerD.SurrenderRequestReference = in.SurrenderRequestReference
	surrenderRequestAnswerD.Action = in.Action
	surrenderRequestAnswerD.Comments = in.Comments
	surrenderRequestAnswerD.SurrenderRequestId = in.SurrenderRequestId

	surrenderRequestAnswerT := eblproto.SurrenderRequestAnswerT{}
	surrenderRequestAnswerT.CreatedDateTime = common.TimeToTimestamp(createdDateTime.UTC().Truncate(time.Second))

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	surrenderRequestAnswer := eblproto.SurrenderRequestAnswer{SurrenderRequestAnswerD: &surrenderRequestAnswerD, SurrenderRequestAnswerT: &surrenderRequestAnswerT, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	err = ss.insertSurrenderRequestAnswer(ctx, insertSurrenderRequestAnswerSQL, &surrenderRequestAnswer, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	surrenderRequestAnswerResponse := eblproto.CreateSurrenderRequestAnswerResponse{}
	surrenderRequestAnswerResponse.SurrenderRequestAnswer = &surrenderRequestAnswer
	return &surrenderRequestAnswerResponse, nil
}

// insertSurrenderRequestAnswer - Insert SurrenderRequestAnswer
func (ss *SurrenderRequestAnswerService) insertSurrenderRequestAnswer(ctx context.Context, insertSurrenderRequestAnswerSQL string, surrenderRequestAnswer *eblproto.SurrenderRequestAnswer, userEmail string, requestID string) error {
	surrenderRequestAnswerTmp, err := ss.CrSurrenderRequestAnswerStruct(ctx, surrenderRequestAnswer, userEmail, requestID)
	if err != nil {
		ss.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}

	err = ss.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertSurrenderRequestAnswerSQL, surrenderRequestAnswerTmp)
		if err != nil {
			ss.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			ss.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		surrenderRequestAnswer.SurrenderRequestAnswerD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(surrenderRequestAnswer.SurrenderRequestAnswerD.Uuid4)
		if err != nil {
			ss.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		surrenderRequestAnswer.SurrenderRequestAnswerD.IdS = uuid4Str
		return nil
	})
	if err != nil {
		ss.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// CrSurrenderRequestAnswerStruct - process SurrenderRequestAnswer details
func (ss *SurrenderRequestAnswerService) CrSurrenderRequestAnswerStruct(ctx context.Context, surrenderRequestAnswer *eblproto.SurrenderRequestAnswer, userEmail string, requestID string) (*eblstruct.SurrenderRequestAnswer, error) {
	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(surrenderRequestAnswer.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(surrenderRequestAnswer.CrUpdTime.UpdatedAt)

	surrenderRequestAnswerT := new(eblstruct.SurrenderRequestAnswerT)
	surrenderRequestAnswerT.CreatedDateTime = common.TimestampToTime(surrenderRequestAnswer.SurrenderRequestAnswerT.CreatedDateTime)

	surrenderRequestAnswerTmp := eblstruct.SurrenderRequestAnswer{SurrenderRequestAnswerD: surrenderRequestAnswer.SurrenderRequestAnswerD, SurrenderRequestAnswerT: surrenderRequestAnswerT, CrUpdUser: surrenderRequestAnswer.CrUpdUser, CrUpdTime: crUpdTime}

	return &surrenderRequestAnswerTmp, nil
}

// UpdateSurrenderRequestAnswer - Update SurrenderRequestAnswer
func (ss *SurrenderRequestAnswerService) UpdateSurrenderRequestAnswer(ctx context.Context, in *eblproto.UpdateSurrenderRequestAnswerRequest) (*eblproto.UpdateSurrenderRequestAnswerResponse, error) {
	db := ss.DBService.DB
	tn := common.GetTimeDetails()

	uuid4byte, err := common.UUIDStrToBytes(in.Id)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	stmt, err := db.PreparexContext(ctx, updateSurrenderRequestAnswerSQL)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = ss.DBService.InsUpd(ctx, in.GetUserEmail(), in.GetRequestId(), func(tx *sqlx.Tx) error {
		_, err = tx.StmtxContext(ctx, stmt).ExecContext(ctx,
			in.SurrenderRequestReference,
			in.Action,
			in.Comments,
			tn,
			uuid4byte)
		if err != nil {
			ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			err1 := stmt.Close()
			if err1 != nil {
				ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err1))
				return err1
			}
			return err
		}
		return nil
	})

	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = stmt.Close()
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	return &eblproto.UpdateSurrenderRequestAnswerResponse{}, nil
}
