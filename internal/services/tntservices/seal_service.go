package tntservices

import (
	"context"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	commonproto "github.com/cloudfresco/sc-dcsa/internal/protogen/common/v1"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"
	tntproto "github.com/cloudfresco/sc-dcsa/internal/protogen/tnt/v3"
	partyservice "github.com/cloudfresco/sc-dcsa/internal/services/partyservices"
	commonstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/common/v1"
	tntstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/tnt/v3"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// SealService - For accessing Transport Document services
type SealService struct {
	log               *zap.Logger
	DBService         *common.DBService
	RedisService      *common.RedisService
	UserServiceClient partyproto.UserServiceClient
	tntproto.UnimplementedSealServiceServer
}

// NewSealService - Create Transport Document service
func NewSealService(log *zap.Logger, dbOpt *common.DBService, redisOpt *common.RedisService, userServiceClient partyproto.UserServiceClient) *SealService {
	return &SealService{
		log:               log,
		DBService:         dbOpt,
		RedisService:      redisOpt,
		UserServiceClient: userServiceClient,
	}
}

// insertSealSQL - insert ChargeSQL query
const insertSealSQL = `insert into seals
	  ( 
uuid4,
utilized_transport_equipment_id,
seal_number,
seal_source_code,
seal_type_code,
status_code,
created_by_user_id,
updated_by_user_id,
created_at,
updated_at)
  values (:uuid4,
:utilized_transport_equipment_id,
:seal_number,
:seal_source_code,
:seal_type_code,
:status_code,
:created_by_user_id,
:updated_by_user_id,
:created_at,
:updated_at);`

// selectSealsSQL - select SealsSQL Query
const selectSealsSQL = `select 
  id,
  uuid4,
  utilized_transport_equipment_id,
  seal_number,
  seal_source_code,
  seal_type_code,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at from seals)`

// CreateSeal - Create Seal
func (ss *SealService) CreateSeal(ctx context.Context, in *tntproto.CreateSealRequest) (*tntproto.CreateSealResponse, error) {
	user, err := partyservice.GetUserWithNewContext(ctx, in.UserId, in.UserEmail, in.RequestId, ss.UserServiceClient)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	ttime := common.GetTimeDetails()
	tn := common.TimeToTimestamp(ttime)

	sealD := tntproto.SealD{}
	sealD.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	sealD.UtilizedTransportEquipmentId = in.UtilizedTransportEquipmentId
	sealD.SealNumber = in.SealNumber
	sealD.SealSourceCode = in.SealSourceCode
	sealD.SealTypeCode = in.SealTypeCode

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	seal := tntproto.Seal{SealD: &sealD, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	err = ss.insertSeal(ctx, insertSealSQL, &seal, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	sealResponse := tntproto.CreateSealResponse{}
	sealResponse.Seal = &seal
	return &sealResponse, nil
}

// insertSeal - Insert seal
func (ss *SealService) insertSeal(ctx context.Context, insertSealSQL string, seal *tntproto.Seal, userEmail string, requestID string) error {
	sealTmp, err := ss.crSealStruct(ctx, seal, userEmail, requestID)
	if err != nil {
		ss.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	err = ss.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertSealSQL, sealTmp)
		if err != nil {
			ss.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			ss.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		seal.SealD.Id = uint32(uID)

		uuid4Str, err := common.UUIDBytesToStr(seal.SealD.Uuid4)
		if err != nil {
			ss.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		seal.SealD.IdS = uuid4Str

		return nil
	})
	if err != nil {
		ss.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// crSealStruct - process Seal details
func (ss *SealService) crSealStruct(ctx context.Context, seal *tntproto.Seal, userEmail string, requestID string) (*tntstruct.Seal, error) {
	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(seal.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(seal.CrUpdTime.UpdatedAt)

	sealTmp := tntstruct.Seal{SealD: seal.SealD, CrUpdUser: seal.CrUpdUser, CrUpdTime: crUpdTime}
	return &sealTmp, nil
}

// GetSeals - Get Seals
func (ss *SealService) GetSeals(ctx context.Context, in *tntproto.GetSealsRequest) (*tntproto.GetSealsResponse, error) {
	limit := in.GetLimit()
	nextCursor := in.GetNextCursor()
	if limit == "" {
		limit = ss.DBService.LimitSQLRows
	}
	query := ""
	if nextCursor == "" {
		query = " order by id desc " + " limit " + limit + ";"
	} else {
		nextCursor = common.DecodeCursor(nextCursor)
		query = " where id <= " + nextCursor + " order by id desc " + " limit " + limit + ";"
	}

	seals := []*tntproto.Seal{}

	nselectSealsSQL := selectSealsSQL + query

	rows, err := ss.DBService.DB.QueryxContext(ctx, nselectSealsSQL)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	for rows.Next() {

		sealTmp := tntstruct.Seal{}
		err = rows.StructScan(&sealTmp)
		if err != nil {
			ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		getRequest := commonproto.GetRequest{}
		getRequest.UserEmail = in.UserEmail
		getRequest.RequestId = in.RequestId
		seal, err := ss.getSealStruct(ctx, &getRequest, sealTmp)
		if err != nil {
			ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		seals = append(seals, seal)

	}

	sealResponse := tntproto.GetSealsResponse{}
	if len(seals) != 0 {
		next := seals[len(seals)-1].SealD.Id
		next--
		nextc := common.EncodeCursor(next)
		sealResponse = tntproto.GetSealsResponse{Seals: seals, NextCursor: nextc}
	} else {
		sealResponse = tntproto.GetSealsResponse{Seals: seals, NextCursor: "0"}
	}
	return &sealResponse, nil
}

// getSealStruct - Get Seal header
func (ss *SealService) getSealStruct(ctx context.Context, in *commonproto.GetRequest, sealTmp tntstruct.Seal) (*tntproto.Seal, error) {
	crUpdTime := new(commonproto.CrUpdTime)
	crUpdTime.CreatedAt = common.TimeToTimestamp(sealTmp.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimeToTimestamp(sealTmp.CrUpdTime.UpdatedAt)

	uuid4Str, err := common.UUIDBytesToStr(sealTmp.SealD.Uuid4)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	sealTmp.SealD.IdS = uuid4Str

	seal := tntproto.Seal{SealD: sealTmp.SealD, CrUpdUser: sealTmp.CrUpdUser, CrUpdTime: crUpdTime}

	return &seal, nil
}
