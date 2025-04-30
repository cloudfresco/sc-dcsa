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

// LegService - For accessing Leg services
type LegService struct {
	log               *zap.Logger
	DBService         *common.DBService
	RedisService      *common.RedisService
	UserServiceClient partyproto.UserServiceClient
	ovsproto.UnimplementedLegServiceServer
}

// NewLegService - Create Leg service
func NewLegService(log *zap.Logger, dbOpt *common.DBService, redisOpt *common.RedisService, userServiceClient partyproto.UserServiceClient) *LegService {
	return &LegService{
		log:               log,
		DBService:         dbOpt,
		RedisService:      redisOpt,
		UserServiceClient: userServiceClient,
	}
}

// InsertLegSQL - insert LegSQL query
const InsertLegSQL = `insert into legs
	  ( 
uuid4,
sequence_number,
mode_of_transport,
vessel_operator_smdg_liner_code,
vessel_imo_number,
vessel_name,
carrier_service_name,
universal_service_reference,
carrier_service_code,
universal_import_voyage_reference,
universal_export_voyage_reference,
carrier_import_voyage_number,
carrier_export_voyage_number,
departure_id,
arrival_id,
point_to_point_routing_id,
status_code,
created_by_user_id,
updated_by_user_id,
created_at,
updated_at
  )
  values (:uuid4,
:sequence_number,
:mode_of_transport,
:vessel_operator_smdg_liner_code,
:vessel_imo_number,
:vessel_name,
:carrier_service_name,
:universal_service_reference,
:carrier_service_code,
:universal_import_voyage_reference,
:universal_export_voyage_reference,
:carrier_import_voyage_number,
:carrier_export_voyage_number,
:departure_id,
:arrival_id,
:point_to_point_routing_id,
:status_code,
:created_by_user_id,
:updated_by_user_id,
:created_at,
:updated_at);`

// selectLegsSQL - select LegsSQL Query
const selectLegsSQL = `select 
  id,
uuid4,
sequence_number,
mode_of_transport,
vessel_operator_smdg_liner_code,
vessel_imo_number,
vessel_name,
carrier_service_name,
universal_service_reference,
carrier_service_code,
universal_import_voyage_reference,
universal_export_voyage_reference,
carrier_import_voyage_number,
carrier_export_voyage_number,
departure_id,
arrival_id,
point_to_point_routing_id,
status_code,
created_by_user_id,
updated_by_user_id,
created_at,
updated_at from legs`

// CreateLeg - Create Leg
func (lg *LegService) CreateLeg(ctx context.Context, in *ovsproto.CreateLegRequest) (*ovsproto.CreateLegResponse, error) {
	leg, err := lg.ProcessLeg(ctx, in)
	if err != nil {
		lg.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = lg.insertLeg(ctx, InsertLegSQL, leg, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		lg.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	legResponse := ovsproto.CreateLegResponse{}
	legResponse.Leg = leg
	return &legResponse, nil
}

// ProcessLeg - Process Leg
func (lg *LegService) ProcessLeg(ctx context.Context, in *ovsproto.CreateLegRequest) (*ovsproto.Leg, error) {
	user, err := partyservice.GetUserWithNewContext(ctx, in.UserId, in.UserEmail, in.RequestId, lg.UserServiceClient)
	if err != nil {
		lg.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	ttime := common.GetTimeDetails()
	tn := common.TimeToTimestamp(ttime)

	legD := ovsproto.LegD{}
	legD.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		lg.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	legD.SequenceNumber = in.SequenceNumber
	legD.ModeOfTransport = in.ModeOfTransport
	legD.VesselOperatorSmdgLinerCode = in.VesselOperatorSmdgLinerCode
	legD.VesselImoNumber = in.VesselImoNumber
	legD.VesselName = in.VesselName
	legD.CarrierServiceName = in.CarrierServiceName
	legD.UniversalServiceReference = in.UniversalServiceReference
	legD.CarrierServiceCode = in.CarrierServiceCode
	legD.UniversalImportVoyageReference = in.UniversalImportVoyageReference
	legD.UniversalExportVoyageReference = in.UniversalExportVoyageReference
	legD.CarrierImportVoyageNumber = in.CarrierImportVoyageNumber
	legD.CarrierExportVoyageNumber = in.CarrierExportVoyageNumber
	legD.DepartureId = in.DepartureId
	legD.ArrivalId = in.ArrivalId
	legD.PointToPointRoutingId = in.PointToPointRoutingId

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	leg := ovsproto.Leg{LegD: &legD, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}
	return &leg, nil
}

// insertLeg - Insert Leg
func (lg *LegService) insertLeg(ctx context.Context, insertLegSQL string, leg *ovsproto.Leg, userEmail string, requestID string) error {
	legTmp, err := lg.CrLegStruct(ctx, leg, userEmail, requestID)
	if err != nil {
		lg.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	err = lg.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertLegSQL, legTmp)
		if err != nil {
			lg.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			lg.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		leg.LegD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(leg.LegD.Uuid4)
		if err != nil {
			lg.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		leg.LegD.IdS = uuid4Str
		return nil
	})
	if err != nil {
		lg.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// CrLegStruct - process Leg details
func (lg *LegService) CrLegStruct(ctx context.Context, leg *ovsproto.Leg, userEmail string, requestID string) (*ovsstruct.Leg, error) {
	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(leg.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(leg.CrUpdTime.UpdatedAt)
	legTmp := ovsstruct.Leg{LegD: leg.LegD, CrUpdUser: leg.CrUpdUser, CrUpdTime: crUpdTime}

	return &legTmp, nil
}

// GetLegs - Get Legs
func (lg *LegService) GetLegs(ctx context.Context, in *ovsproto.GetLegsRequest) (*ovsproto.GetLegsResponse, error) {
	limit := in.GetLimit()
	nextCursor := in.GetNextCursor()
	if limit == "" {
		limit = lg.DBService.LimitSQLRows
	}
	query := ""
	if nextCursor == "" {
		query = " order by id desc " + " limit " + limit + ";"
	} else {
		nextCursor = common.DecodeCursor(nextCursor)
		query = " where id <= " + nextCursor + " order by id desc " + " limit " + limit + ";"
	}

	legs := []*ovsproto.Leg{}

	nselectLegsSQL := selectLegsSQL + query

	rows, err := lg.DBService.DB.QueryxContext(ctx, nselectLegsSQL)
	if err != nil {
		lg.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	for rows.Next() {

		legTmp := ovsstruct.Leg{}
		err = rows.StructScan(&legTmp)
		if err != nil {
			lg.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		getRequest := commonproto.GetRequest{}
		getRequest.UserEmail = in.UserEmail
		getRequest.RequestId = in.RequestId
		leg, err := lg.getLegStruct(ctx, &getRequest, legTmp)
		if err != nil {
			lg.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		legs = append(legs, leg)

	}

	legsResponse := ovsproto.GetLegsResponse{}
	if len(legs) != 0 {
		next := legs[len(legs)-1].LegD.Id
		next--
		nextc := common.EncodeCursor(next)
		legsResponse = ovsproto.GetLegsResponse{Legs: legs, NextCursor: nextc}
	} else {
		legsResponse = ovsproto.GetLegsResponse{Legs: legs, NextCursor: "0"}
	}
	return &legsResponse, nil
}

// getLegStruct - Get Leg header
func (lg *LegService) getLegStruct(ctx context.Context, in *commonproto.GetRequest, legTmp ovsstruct.Leg) (*ovsproto.Leg, error) {
	crUpdTime := new(commonproto.CrUpdTime)
	crUpdTime.CreatedAt = common.TimeToTimestamp(legTmp.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimeToTimestamp(legTmp.CrUpdTime.UpdatedAt)

	uuid4Str, err := common.UUIDBytesToStr(legTmp.LegD.Uuid4)
	if err != nil {
		lg.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	legTmp.LegD.IdS = uuid4Str

	leg := ovsproto.Leg{LegD: legTmp.LegD, CrUpdUser: legTmp.CrUpdUser, CrUpdTime: crUpdTime}

	return &leg, nil
}

// GetLeg - Get Leg
func (lg *LegService) GetLeg(ctx context.Context, inReq *ovsproto.GetLegRequest) (*ovsproto.GetLegResponse, error) {
	in := inReq.GetRequest
	uuid4byte, err := common.UUIDStrToBytes(in.Id)
	if err != nil {
		lg.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	nselectLegsSQL := selectLegsSQL + ` where uuid4 = ?`
	row := lg.DBService.DB.QueryRowxContext(ctx, nselectLegsSQL, uuid4byte)
	legTmp := ovsstruct.Leg{}
	err = row.StructScan(&legTmp)
	if err != nil {
		lg.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	getRequest := commonproto.GetRequest{}
	getRequest.UserEmail = in.UserEmail
	getRequest.RequestId = in.RequestId
	leg, err := lg.getLegStruct(ctx, &getRequest, legTmp)
	if err != nil {
		lg.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	legResponse := ovsproto.GetLegResponse{}
	legResponse.Leg = leg
	return &legResponse, nil
}

// GetLegByPk - Get Leg By Primary key(Id)
func (lg *LegService) GetLegByPk(ctx context.Context, inReq *ovsproto.GetLegByPkRequest) (*ovsproto.GetLegByPkResponse, error) {
	in := inReq.GetByIdRequest
	nselectLegsSQL := selectLegsSQL + ` where id = ?;`
	row := lg.DBService.DB.QueryRowxContext(ctx, nselectLegsSQL, in.Id)
	legTmp := ovsstruct.Leg{}
	err := row.StructScan(&legTmp)
	if err != nil {
		lg.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	getRequest := commonproto.GetRequest{}
	getRequest.UserEmail = in.UserEmail
	getRequest.RequestId = in.RequestId
	leg, err := lg.getLegStruct(ctx, &getRequest, legTmp)
	if err != nil {
		lg.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	legResponse := ovsproto.GetLegByPkResponse{}
	legResponse.Leg = leg
	return &legResponse, nil
}
