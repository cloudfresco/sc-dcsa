package ovsservices

import (
	"context"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	commonproto "github.com/cloudfresco/sc-dcsa/internal/protogen/common/v1"
	ovsproto "github.com/cloudfresco/sc-dcsa/internal/protogen/ovs/v3"
	partyservice "github.com/cloudfresco/sc-dcsa/internal/services/partyservices"
	commonstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/common/v1"
	ovsstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/ovs/v3"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// insertVesselSQL - insert VesselSQL Query
const insertVesselSQL = `insert into vessels
	  ( 
  uuid4,
  vessel_imo_number,
  vessel_name,
  vessel_flag,
  vessel_call_sign,
  is_dummy_vessel,
  vessel_operator_carrier_code,
  vessel_operator_carrier_code_list_provider,
  vessel_length,
  vessel_width,
  dimension_unit,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at
  )
  values (:uuid4,
:vessel_imo_number,
:vessel_name,
:vessel_flag,
:vessel_call_sign,
:is_dummy_vessel,
:vessel_operator_carrier_code,
:vessel_operator_carrier_code_list_provider,
:vessel_length,
:vessel_width,
:dimension_unit,
:status_code,
:created_by_user_id,
:updated_by_user_id,
:created_at,
:updated_at);`

// selectVesselsSQL - select VesselsSQL Query
const selectVesselsSQL = `select 
  id,
  uuid4,
  vessel_imo_number,
  vessel_name,
  vessel_flag,
  vessel_call_sign,
  is_dummy_vessel,
  vessel_operator_carrier_code,
  vessel_operator_carrier_code_list_provider,
  vessel_length,
  vessel_width,
  dimension_unit,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at from vessels`

// CreateVessel - Create  Vessel
func (vs *VesselScheduleService) CreateVessel(ctx context.Context, in *ovsproto.CreateVesselRequest) (*ovsproto.CreateVesselResponse, error) {
	user, err := partyservice.GetUserWithNewContext(ctx, in.UserId, in.UserEmail, in.RequestId, vs.UserServiceClient)
	if err != nil {
		vs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	ttime := common.GetTimeDetails()
	tn := common.TimeToTimestamp(ttime)

	vesselD := ovsproto.VesselD{}
	vesselD.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		vs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	vesselD.VesselImoNumber = in.VesselImoNumber
	vesselD.VesselName = in.VesselName
	vesselD.VesselFlag = in.VesselFlag
	vesselD.VesselCallSign = in.VesselCallSign
	vesselD.IsDummyVessel = in.IsDummyVessel
	vesselD.VesselOperatorCarrierCode = in.VesselOperatorCarrierCode
	vesselD.VesselOperatorCarrierCodeListProvider = in.VesselOperatorCarrierCodeListProvider
	vesselD.VesselLength = in.VesselLength
	vesselD.VesselWidth = in.VesselWidth
	vesselD.DimensionUnit = in.DimensionUnit

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	vessel := ovsproto.Vessel{VesselD: &vesselD, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	err = vs.insertVessel(ctx, insertVesselSQL, &vessel, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		vs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	vesselResponse := ovsproto.CreateVesselResponse{}
	vesselResponse.Vessel = &vessel
	return &vesselResponse, nil
}

// insertVessel - Insert Vessel
func (vs *VesselScheduleService) insertVessel(ctx context.Context, insertVesselSQL string, vessel *ovsproto.Vessel, userEmail string, requestID string) error {
	vesselTmp, err := vs.crVesselStruct(ctx, vessel, userEmail, requestID)
	if err != nil {
		vs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	err = vs.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertVesselSQL, vesselTmp)
		if err != nil {
			vs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			vs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		vessel.VesselD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(vessel.VesselD.Uuid4)
		if err != nil {
			vs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		vessel.VesselD.IdS = uuid4Str
		return nil
	})
	if err != nil {
		vs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// crVesselStruct - process Vessel details
func (vs *VesselScheduleService) crVesselStruct(ctx context.Context, vessel *ovsproto.Vessel, userEmail string, requestID string) (*ovsstruct.Vessel, error) {
	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(vessel.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(vessel.CrUpdTime.UpdatedAt)
	vesselTmp := ovsstruct.Vessel{VesselD: vessel.VesselD, CrUpdUser: vessel.CrUpdUser, CrUpdTime: crUpdTime}

	return &vesselTmp, nil
}

// GetVessels - Get  Vessels
func (vs *VesselScheduleService) GetVessels(ctx context.Context, in *ovsproto.GetVesselsRequest) (*ovsproto.GetVesselsResponse, error) {
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

	vessels := []*ovsproto.Vessel{}

	nselectVesselsSQL := selectVesselsSQL + query

	rows, err := vs.DBService.DB.QueryxContext(ctx, nselectVesselsSQL)
	if err != nil {
		vs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	for rows.Next() {

		vesselTmp := ovsstruct.Vessel{}
		err = rows.StructScan(&vesselTmp)
		if err != nil {
			vs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		getRequest := commonproto.GetRequest{}
		getRequest.UserEmail = in.UserEmail
		getRequest.RequestId = in.RequestId
		vessel, err := vs.getVesselStruct(ctx, &getRequest, vesselTmp)
		if err != nil {
			vs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		vessels = append(vessels, vessel)

	}

	vesselsResponse := ovsproto.GetVesselsResponse{}
	if len(vessels) != 0 {
		next := vessels[len(vessels)-1].VesselD.Id
		next--
		nextc := common.EncodeCursor(next)
		vesselsResponse = ovsproto.GetVesselsResponse{Vessels: vessels, NextCursor: nextc}
	} else {
		vesselsResponse = ovsproto.GetVesselsResponse{Vessels: vessels, NextCursor: "0"}
	}
	return &vesselsResponse, nil
}

// GetVessel - Get Vessel
func (vs *VesselScheduleService) GetVessel(ctx context.Context, inReq *ovsproto.GetVesselRequest) (*ovsproto.GetVesselResponse, error) {
	in := inReq.GetRequest
	uuid4byte, err := common.UUIDStrToBytes(in.Id)
	if err != nil {
		vs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	nselectVesselsSQL := selectVesselsSQL + ` where uuid4 = ?`
	row := vs.DBService.DB.QueryRowxContext(ctx, nselectVesselsSQL, uuid4byte)
	vesselTmp := ovsstruct.Vessel{}
	err = row.StructScan(&vesselTmp)
	if err != nil {
		vs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	getRequest := commonproto.GetRequest{}
	getRequest.UserEmail = in.UserEmail
	getRequest.RequestId = in.RequestId
	vessel, err := vs.getVesselStruct(ctx, &getRequest, vesselTmp)
	if err != nil {
		vs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	vesselResponse := ovsproto.GetVesselResponse{}
	vesselResponse.Vessel = vessel
	return &vesselResponse, nil
}

// GetVesselByPk - Get Vessel By Primary key(Id)
func (vs *VesselScheduleService) GetVesselByPk(ctx context.Context, inReq *ovsproto.GetVesselByPkRequest) (*ovsproto.GetVesselByPkResponse, error) {
	in := inReq.GetByIdRequest
	nselectVesselsSQL := selectVesselsSQL + ` where id = ?;`
	row := vs.DBService.DB.QueryRowxContext(ctx, nselectVesselsSQL, in.Id)
	vesselTmp := ovsstruct.Vessel{}
	err := row.StructScan(&vesselTmp)
	if err != nil {
		vs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	getRequest := commonproto.GetRequest{}
	getRequest.UserEmail = in.UserEmail
	getRequest.RequestId = in.RequestId
	vessel, err := vs.getVesselStruct(ctx, &getRequest, vesselTmp)
	if err != nil {
		vs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	vesselResponse := ovsproto.GetVesselByPkResponse{}
	vesselResponse.Vessel = vessel
	return &vesselResponse, nil
}

// getVesselStruct - Get vessel
func (vs *VesselScheduleService) getVesselStruct(ctx context.Context, in *commonproto.GetRequest, vesselTmp ovsstruct.Vessel) (*ovsproto.Vessel, error) {
	crUpdTime := new(commonproto.CrUpdTime)
	crUpdTime.CreatedAt = common.TimeToTimestamp(vesselTmp.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimeToTimestamp(vesselTmp.CrUpdTime.UpdatedAt)

	uuid4Str, err := common.UUIDBytesToStr(vesselTmp.VesselD.Uuid4)
	if err != nil {
		vs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	vesselTmp.VesselD.IdS = uuid4Str

	vessel := ovsproto.Vessel{VesselD: vesselTmp.VesselD, CrUpdUser: vesselTmp.CrUpdUser, CrUpdTime: crUpdTime}

	return &vessel, nil
}
