package ovsservices

import (
	"context"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	commonproto "github.com/cloudfresco/sc-dcsa/internal/protogen/common/v1"
	ovsproto "github.com/cloudfresco/sc-dcsa/internal/protogen/ovs/v3"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"
	partyservice "github.com/cloudfresco/sc-dcsa/internal/services/partyservices"
	commonstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/common/v1"
	ovsstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/ovs/v3"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// VesselScheduleService - For accessing VesselSchedule services
type VesselScheduleService struct {
	log               *zap.Logger
	DBService         *common.DBService
	RedisService      *common.RedisService
	UserServiceClient partyproto.UserServiceClient
	ovsproto.UnimplementedVesselScheduleServiceServer
}

// NewVesselScheduleService - Create VesselSchedule service
func NewVesselScheduleService(log *zap.Logger, dbOpt *common.DBService, redisOpt *common.RedisService, userServiceClient partyproto.UserServiceClient) *VesselScheduleService {
	return &VesselScheduleService{
		log:               log,
		DBService:         dbOpt,
		RedisService:      redisOpt,
		UserServiceClient: userServiceClient,
	}
}

// insertVesselScheduleSQL - insert VesselScheduleSQL Query
const insertVesselScheduleSQL = `insert into vessel_schedules
	  ( 
  uuid4,
  vessel_id,
  service_id,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at
  )
  values (:uuid4,
:vessel_id,
:service_id,
:status_code,
:created_by_user_id,
:updated_by_user_id,
:created_at,
:updated_at);`

// selectVesselSchedulesSQL - select VesselSchedulesSQL Query
const selectVesselSchedulesSQL = `select 
  id,
  uuid4,
  vessel_id,
  service_id,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at from vessel_schedules`

// CreateVesselSchedule - Create  VesselSchedule
func (vs *VesselScheduleService) CreateVesselSchedule(ctx context.Context, in *ovsproto.CreateVesselScheduleRequest) (*ovsproto.CreateVesselScheduleResponse, error) {
	user, err := partyservice.GetUserWithNewContext(ctx, in.UserId, in.UserEmail, in.RequestId, vs.UserServiceClient)
	if err != nil {
		vs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	ttime := common.GetTimeDetails()
	tn := common.TimeToTimestamp(ttime)

	vesselScheduleD := ovsproto.VesselScheduleD{}
	vesselScheduleD.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		vs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	vesselScheduleD.VesselId = in.VesselId
	vesselScheduleD.ServiceId = in.ServiceId

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	vesselSchedule := ovsproto.VesselSchedule{VesselScheduleD: &vesselScheduleD, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	err = vs.insertVesselSchedule(ctx, insertVesselScheduleSQL, &vesselSchedule, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		vs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	vesselScheduleResponse := ovsproto.CreateVesselScheduleResponse{}
	vesselScheduleResponse.VesselSchedule = &vesselSchedule
	return &vesselScheduleResponse, nil
}

// insertVesselSchedule - Insert VesselSchedule
func (vs *VesselScheduleService) insertVesselSchedule(ctx context.Context, insertVesselScheduleSQL string, vesselSchedule *ovsproto.VesselSchedule, userEmail string, requestID string) error {
	vesselScheduleTmp, err := vs.crVesselScheduleStruct(ctx, vesselSchedule, userEmail, requestID)
	if err != nil {
		vs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	err = vs.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertVesselScheduleSQL, vesselScheduleTmp)
		if err != nil {
			vs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			vs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		vesselSchedule.VesselScheduleD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(vesselSchedule.VesselScheduleD.Uuid4)
		if err != nil {
			vs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		vesselSchedule.VesselScheduleD.IdS = uuid4Str
		return nil
	})
	if err != nil {
		vs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// crVesselScheduleStruct - process VesselSchedule details
func (vs *VesselScheduleService) crVesselScheduleStruct(ctx context.Context, vesselSchedule *ovsproto.VesselSchedule, userEmail string, requestID string) (*ovsstruct.VesselSchedule, error) {
	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(vesselSchedule.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(vesselSchedule.CrUpdTime.UpdatedAt)
	vesselScheduleTmp := ovsstruct.VesselSchedule{VesselScheduleD: vesselSchedule.VesselScheduleD, CrUpdUser: vesselSchedule.CrUpdUser, CrUpdTime: crUpdTime}

	return &vesselScheduleTmp, nil
}

// GetVesselSchedules - Get  VesselSchedules
func (vs *VesselScheduleService) GetVesselSchedules(ctx context.Context, in *ovsproto.GetVesselSchedulesRequest) (*ovsproto.GetVesselSchedulesResponse, error) {
	limit := in.GetLimit()
	nextCursor := in.GetNextCursor()
	if limit == "" {
		limit = vs.DBService.LimitSQLRows
	}
	query := ""
	if nextCursor == "" {
		query = " order by id desc " + " limit " + limit + ";"
	} else {
		nextCursor = common.DecodeCursor(nextCursor)
		query = " where id <= " + nextCursor + " order by id desc " + " limit " + limit + ";"
	}

	vesselSchedules := []*ovsproto.VesselSchedule{}

	nselectVesselSchedulesSQL := selectVesselSchedulesSQL + query

	rows, err := vs.DBService.DB.QueryxContext(ctx, nselectVesselSchedulesSQL)
	if err != nil {
		vs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	for rows.Next() {

		vesselScheduleTmp := ovsstruct.VesselSchedule{}
		err = rows.StructScan(&vesselScheduleTmp)
		if err != nil {
			vs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		getRequest := commonproto.GetRequest{}
		getRequest.UserEmail = in.UserEmail
		getRequest.RequestId = in.RequestId
		vesselSchedule, err := vs.getVesselScheduleStruct(ctx, &getRequest, vesselScheduleTmp)
		if err != nil {
			vs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		vesselSchedules = append(vesselSchedules, vesselSchedule)

	}

	vesselSchedulesResponse := ovsproto.GetVesselSchedulesResponse{}
	if len(vesselSchedules) != 0 {
		next := vesselSchedules[len(vesselSchedules)-1].VesselScheduleD.Id
		next--
		nextc := common.EncodeCursor(next)
		vesselSchedulesResponse = ovsproto.GetVesselSchedulesResponse{VesselSchedules: vesselSchedules, NextCursor: nextc}
	} else {
		vesselSchedulesResponse = ovsproto.GetVesselSchedulesResponse{VesselSchedules: vesselSchedules, NextCursor: "0"}
	}
	return &vesselSchedulesResponse, nil
}

// GetVesselSchedule - Get VesselSchedule
func (vs *VesselScheduleService) GetVesselSchedule(ctx context.Context, inReq *ovsproto.GetVesselScheduleRequest) (*ovsproto.GetVesselScheduleResponse, error) {
	in := inReq.GetRequest
	uuid4byte, err := common.UUIDStrToBytes(in.Id)
	if err != nil {
		vs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	nselectVesselSchedulesSQL := selectVesselSchedulesSQL + ` where uuid4 = ?`
	row := vs.DBService.DB.QueryRowxContext(ctx, nselectVesselSchedulesSQL, uuid4byte)
	vesselScheduleTmp := ovsstruct.VesselSchedule{}
	err = row.StructScan(&vesselScheduleTmp)
	if err != nil {
		vs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	getRequest := commonproto.GetRequest{}
	getRequest.UserEmail = in.UserEmail
	getRequest.RequestId = in.RequestId
	vesselSchedule, err := vs.getVesselScheduleStruct(ctx, &getRequest, vesselScheduleTmp)
	if err != nil {
		vs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	vesselScheduleResponse := ovsproto.GetVesselScheduleResponse{}
	vesselScheduleResponse.VesselSchedule = vesselSchedule
	return &vesselScheduleResponse, nil
}

// GetVesselScheduleByPk - Get VesselSchedule By Primary key(Id)
func (vs *VesselScheduleService) GetVesselScheduleByPk(ctx context.Context, inReq *ovsproto.GetVesselScheduleByPkRequest) (*ovsproto.GetVesselScheduleByPkResponse, error) {
	in := inReq.GetByIdRequest
	nselectVesselSchedulesSQL := selectVesselSchedulesSQL + ` where id = ?;`
	row := vs.DBService.DB.QueryRowxContext(ctx, nselectVesselSchedulesSQL, in.Id)
	vesselScheduleTmp := ovsstruct.VesselSchedule{}
	err := row.StructScan(&vesselScheduleTmp)
	if err != nil {
		vs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	getRequest := commonproto.GetRequest{}
	getRequest.UserEmail = in.UserEmail
	getRequest.RequestId = in.RequestId
	vesselSchedule, err := vs.getVesselScheduleStruct(ctx, &getRequest, vesselScheduleTmp)
	if err != nil {
		vs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	vesselScheduleResponse := ovsproto.GetVesselScheduleByPkResponse{}
	vesselScheduleResponse.VesselSchedule = vesselSchedule
	return &vesselScheduleResponse, nil
}

// getVesselScheduleStruct - Get vesselSchedule
func (vs *VesselScheduleService) getVesselScheduleStruct(ctx context.Context, in *commonproto.GetRequest, vesselScheduleTmp ovsstruct.VesselSchedule) (*ovsproto.VesselSchedule, error) {
	crUpdTime := new(commonproto.CrUpdTime)
	crUpdTime.CreatedAt = common.TimeToTimestamp(vesselScheduleTmp.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimeToTimestamp(vesselScheduleTmp.CrUpdTime.UpdatedAt)

	uuid4Str, err := common.UUIDBytesToStr(vesselScheduleTmp.VesselScheduleD.Uuid4)
	if err != nil {
		vs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	vesselScheduleTmp.VesselScheduleD.IdS = uuid4Str

	vesselSchedule := ovsproto.VesselSchedule{VesselScheduleD: vesselScheduleTmp.VesselScheduleD, CrUpdUser: vesselScheduleTmp.CrUpdUser, CrUpdTime: crUpdTime}

	return &vesselSchedule, nil
}
