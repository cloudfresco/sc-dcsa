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

// insertFacilitySQL - insert FacilitySQL query
const insertFacilitySQL = `insert into facilities
	  ( 
  facility_name,
  un_location_code,
  facility_bic_code,
  facility_smdg_code,
  location_id,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at)
  values (
:facility_name,
:un_location_code,
:facility_bic_code,
:facility_smdg_code,
:location_id,
:status_code,
:created_by_user_id,
:updated_by_user_id,
:created_at,
:updated_at);`

// selectFacilitiesSQL - select FacilitiesSQL query
/*const selectFacilitiesSQL = `select
  id,
  facility_name,
  un_location_code,
  facility_bic_code,
  facility_smdg_code,
  location_id,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at from facilities`*/

// CreateFacility - CreateFacility
func (ps *PartyService) CreateFacility(ctx context.Context, in *partyproto.CreateFacilityRequest) (*partyproto.CreateFacilityResponse, error) {
	facility, err := ps.ProcessFacilityRequest(ctx, in)
	if err != nil {
		ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = ps.insertFacility(ctx, insertFacilitySQL, facility, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	facilityResponse := partyproto.CreateFacilityResponse{}
	facilityResponse.Facility = facility
	return &facilityResponse, nil
}

// ProcessFacilityRequest - ProcessFacilityRequest
func (ps *PartyService) ProcessFacilityRequest(ctx context.Context, in *partyproto.CreateFacilityRequest) (*partyproto.Facility, error) {
	user, err := GetUserWithNewContext(ctx, in.UserId, in.UserEmail, in.RequestId, ps.UserServiceClient)
	if err != nil {
		ps.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	ttime := common.GetTimeDetails()
	tn := common.TimeToTimestamp(ttime)

	facilityD := partyproto.FacilityD{}

	facilityD.FacilityName = in.FacilityName
	facilityD.UnLocationCode = in.UnLocationCode
	facilityD.FacilityBicCode = in.FacilityBicCode
	facilityD.FacilitySmdgCode = in.FacilitySmdgCode
	facilityD.LocationId = in.LocationId

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	facility := partyproto.Facility{FacilityD: &facilityD, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	return &facility, nil
}

// insertFacility - Insert Document Party
func (ps *PartyService) insertFacility(ctx context.Context, insertFacilitySQL string, facility *partyproto.Facility, userEmail string, requestID string) error {
	facilityTmp, err := ps.crFacilityStruct(ctx, facility, userEmail, requestID)
	if err != nil {
		ps.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	err = ps.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertFacilitySQL, facilityTmp)
		if err != nil {
			ps.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			ps.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		facility.FacilityD.Id = uint32(uID)
		return nil
	})
	if err != nil {
		ps.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// crFacilityStruct - process Facility details
func (ps *PartyService) crFacilityStruct(ctx context.Context, facility *partyproto.Facility, userEmail string, requestID string) (*partystruct.Facility, error) {
	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(facility.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(facility.CrUpdTime.UpdatedAt)

	facilityTmp := partystruct.Facility{FacilityD: facility.FacilityD, CrUpdUser: facility.CrUpdUser, CrUpdTime: crUpdTime}
	return &facilityTmp, nil
}
