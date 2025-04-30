package jitservices

import (
	"context"
	"time"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	commonproto "github.com/cloudfresco/sc-dcsa/internal/protogen/common/v1"
	jitproto "github.com/cloudfresco/sc-dcsa/internal/protogen/jit/v1"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"
	jitstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/jit/v1"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// TimestampService - For accessing Transport Document services
type TimestampService struct {
	log               *zap.Logger
	DBService         *common.DBService
	RedisService      *common.RedisService
	UserServiceClient partyproto.UserServiceClient
	jitproto.UnimplementedTimestampServiceServer
}

// NewTimestampService - Create Transport Document service
func NewTimestampService(log *zap.Logger, dbOpt *common.DBService, redisOpt *common.RedisService, userServiceClient partyproto.UserServiceClient) *TimestampService {
	return &TimestampService{
		log:               log,
		DBService:         dbOpt,
		RedisService:      redisOpt,
		UserServiceClient: userServiceClient,
	}
}

// insertTimestampSQL - insert TimestampSQL Query
const insertTimestampSQL = `insert into timestamps
	  ( 
uuid4,
event_type_code,
event_classifier_code,
delay_reason_code,
change_remark,
event_date_time)
  values (:uuid4,
:event_type_code,
:event_classifier_code,
:delay_reason_code,
:change_remark,
:event_date_time);`

// selectTimestampsSQL - select TimestampsSQL Query
const selectTimestampsSQL = `select 
  id,
  uuid4,
  event_type_code,
  event_classifier_code,
  delay_reason_code,
  change_remark,
  event_date_time from timestamps`

// CreateTimestamp - Create Timestamp
func (ts *TimestampService) CreateTimestamp(ctx context.Context, in *jitproto.CreateTimestampRequest) (*jitproto.CreateTimestampResponse, error) {
	eventDateTime, err := time.Parse(common.Layout, in.EventDateTime)
	if err != nil {
		ts.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	timestampD := jitproto.TimestampD{}
	timestampD.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		ts.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	timestampD.EventTypeCode = in.EventTypeCode
	timestampD.EventClassifierCode = in.EventClassifierCode
	timestampD.DelayReasonCode = in.DelayReasonCode
	timestampD.ChangeRemark = in.ChangeRemark

	timestampT := jitproto.TimestampT{}
	timestampT.EventDateTime = common.TimeToTimestamp(eventDateTime.UTC().Truncate(time.Second))

	tstamp := jitproto.Timestamp{TimestampD: &timestampD, TimestampT: &timestampT}

	err = ts.insertTimestamp(ctx, insertTimestampSQL, &tstamp, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		ts.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	timestampResponse := jitproto.CreateTimestampResponse{}
	timestampResponse.Timestamp1 = &tstamp
	return &timestampResponse, nil
}

// insertTimestamp - Insert Timestamp
func (ts *TimestampService) insertTimestamp(ctx context.Context, insertTimestampSQL string, tstamp *jitproto.Timestamp, userEmail string, requestID string) error {
	timestampTmp, err := ts.crTimestampStruct(ctx, tstamp, userEmail, requestID)
	if err != nil {
		ts.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}

	err = ts.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertTimestampSQL, timestampTmp)
		if err != nil {
			ts.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			ts.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		tstamp.TimestampD.Id = uint32(uID)
		return nil
	})
	if err != nil {
		ts.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// crTimestampStruct - process Timestamp details
func (ts *TimestampService) crTimestampStruct(ctx context.Context, tstamp *jitproto.Timestamp, userEmail string, requestID string) (*jitstruct.Timestamp, error) {
	timestampTmpT := new(jitstruct.TimestampT)
	timestampTmpT.EventDateTime = common.TimestampToTime(tstamp.TimestampT.EventDateTime)
	timestampTmp := jitstruct.Timestamp{TimestampD: tstamp.TimestampD, TimestampT: timestampTmpT}
	return &timestampTmp, nil
}

// GetTimestamps - Get Timestamps
func (ts *TimestampService) GetTimestamps(ctx context.Context, in *jitproto.GetTimestampsRequest) (*jitproto.GetTimestampsResponse, error) {
	limit := in.GetLimit()
	nextCursor := in.GetNextCursor()
	if limit == "" {
		limit = ts.DBService.LimitSQLRows
	}
	query := ""
	if nextCursor == "" {
		query = " order by id desc " + " limit " + limit + ";"
	} else {
		nextCursor = common.DecodeCursor(nextCursor)
		query = " where id <= " + nextCursor + " order by id desc " + " limit " + limit + ";"
	}

	tstamps := []*jitproto.Timestamp{}

	nselectTimestampsSQL := selectTimestampsSQL + query

	rows, err := ts.DBService.DB.QueryxContext(ctx, nselectTimestampsSQL)
	if err != nil {
		ts.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	for rows.Next() {

		timestampTmp := jitstruct.Timestamp{}
		err = rows.StructScan(&timestampTmp)
		if err != nil {
			ts.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		getRequest := commonproto.GetRequest{}
		getRequest.UserEmail = in.UserEmail
		getRequest.RequestId = in.RequestId
		tstamp, err := ts.getTimestampStruct(ctx, &getRequest, timestampTmp)
		if err != nil {
			ts.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		tstamps = append(tstamps, tstamp)

	}

	tstampsResponse := jitproto.GetTimestampsResponse{}
	if len(tstamps) != 0 {
		next := tstamps[len(tstamps)-1].TimestampD.Id
		next--
		nextc := common.EncodeCursor(next)
		tstampsResponse = jitproto.GetTimestampsResponse{Timestamps: tstamps, NextCursor: nextc}
	} else {
		tstampsResponse = jitproto.GetTimestampsResponse{Timestamps: tstamps, NextCursor: "0"}
	}
	return &tstampsResponse, nil
}

// getTimestampStruct - Get Timestamp header
func (ts *TimestampService) getTimestampStruct(ctx context.Context, in *commonproto.GetRequest, timestampTmp jitstruct.Timestamp) (*jitproto.Timestamp, error) {
	timestampT := new(jitproto.TimestampT)
	timestampT.EventDateTime = common.TimeToTimestamp(timestampTmp.TimestampT.EventDateTime)

	tstamp := jitproto.Timestamp{TimestampD: timestampTmp.TimestampD, TimestampT: timestampT}
	return &tstamp, nil
}
