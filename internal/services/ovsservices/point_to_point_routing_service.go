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

// insertPointToPointRoutingSQL - insert PointToPointRoutingSQL Query
const insertPointToPointRoutingSQL = `insert into point_to_point_routings
	  ( 
uuid4,
sequence_number,
place_of_receipt_id,
place_of_delivery_id,
status_code,
created_by_user_id,
updated_by_user_id,
created_at,
updated_at
  )
  values (:uuid4,
:sequence_number,
:place_of_receipt_id,
:place_of_delivery_id,
:status_code,
:created_by_user_id,
:updated_by_user_id,
:created_at,
:updated_at);`

// selectPointToPointRoutingsSQL - select PointToPointRoutingsSQL Query
const selectPointToPointRoutingsSQL = `select 
  id,
  uuid4,
  sequence_number,
  place_of_receipt_id,
  place_of_delivery_id,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at from point_to_point_routings`

// CreatePointToPointRouting - Create  PointToPointRouting
func (lg *LegService) CreatePointToPointRouting(ctx context.Context, in *ovsproto.CreatePointToPointRoutingRequest) (*ovsproto.CreatePointToPointRoutingResponse, error) {
	user, err := partyservice.GetUserWithNewContext(ctx, in.UserId, in.UserEmail, in.RequestId, lg.UserServiceClient)
	if err != nil {
		lg.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	ttime := common.GetTimeDetails()
	tn := common.TimeToTimestamp(ttime)

	pointToPointRoutingD := ovsproto.PointToPointRoutingD{}
	pointToPointRoutingD.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		lg.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	pointToPointRoutingD.SequenceNumber = in.SequenceNumber
	pointToPointRoutingD.PlaceOfReceiptId = in.PlaceOfReceiptId
	pointToPointRoutingD.PlaceOfDeliveryId = in.PlaceOfDeliveryId

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	pointToPointRouting := ovsproto.PointToPointRouting{PointToPointRoutingD: &pointToPointRoutingD, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	legs := []*ovsproto.Leg{}
	for _, leg := range in.Legs {
		leg.UserId = in.UserId
		leg.UserEmail = in.UserEmail
		leg.RequestId = in.RequestId
		leg, err := lg.ProcessLeg(ctx, leg)
		if err != nil {
			lg.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		legs = append(legs, leg)
	}

	err = lg.insertPointToPointRouting(ctx, insertPointToPointRoutingSQL, &pointToPointRouting, InsertLegSQL, legs, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		lg.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	pointToPointRoutingResponse := ovsproto.CreatePointToPointRoutingResponse{}
	pointToPointRoutingResponse.PointToPointRouting = &pointToPointRouting
	return &pointToPointRoutingResponse, nil
}

// insertPointToPointRouting - Insert PointToPointRouting
func (lg *LegService) insertPointToPointRouting(ctx context.Context, insertPointToPointRoutingSQL string, pointToPointRouting *ovsproto.PointToPointRouting, insertLegSQL string, legs []*ovsproto.Leg, userEmail string, requestID string) error {
	pointToPointRoutingTmp, err := lg.crPointToPointRoutingStruct(ctx, pointToPointRouting, userEmail, requestID)
	if err != nil {
		lg.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	err = lg.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertPointToPointRoutingSQL, pointToPointRoutingTmp)
		if err != nil {
			lg.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			lg.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		pointToPointRouting.PointToPointRoutingD.Id = uint32(uID)

		for _, leg := range legs {
			leg.LegD.PointToPointRoutingId = pointToPointRouting.PointToPointRoutingD.Id
			legTmp, err := lg.CrLegStruct(ctx, leg, userEmail, requestID)
			if err != nil {
				lg.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
				return err
			}
			_, err = tx.NamedExecContext(ctx, insertLegSQL, legTmp)
			if err != nil {
				lg.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
				return err
			}
		}

		uuid4Str, err := common.UUIDBytesToStr(pointToPointRouting.PointToPointRoutingD.Uuid4)
		if err != nil {
			lg.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		pointToPointRouting.PointToPointRoutingD.IdS = uuid4Str
		pointToPointRouting.Legs = legs
		return nil
	})
	if err != nil {
		lg.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// crPointToPointRoutingStruct - process PointToPointRouting details
func (lg *LegService) crPointToPointRoutingStruct(ctx context.Context, pointToPointRouting *ovsproto.PointToPointRouting, userEmail string, requestID string) (*ovsstruct.PointToPointRouting, error) {
	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(pointToPointRouting.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(pointToPointRouting.CrUpdTime.UpdatedAt)
	pointToPointRoutingTmp := ovsstruct.PointToPointRouting{PointToPointRoutingD: pointToPointRouting.PointToPointRoutingD, CrUpdUser: pointToPointRouting.CrUpdUser, CrUpdTime: crUpdTime}

	return &pointToPointRoutingTmp, nil
}

// GetPointToPointRoutings - Get  PointToPointRoutings
func (lg *LegService) GetPointToPointRoutings(ctx context.Context, in *ovsproto.GetPointToPointRoutingsRequest) (*ovsproto.GetPointToPointRoutingsResponse, error) {
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

	pointToPointRoutings := []*ovsproto.PointToPointRouting{}

	nselectPointToPointRoutingsSQL := selectPointToPointRoutingsSQL + query

	rows, err := lg.DBService.DB.QueryxContext(ctx, nselectPointToPointRoutingsSQL)
	if err != nil {
		lg.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	for rows.Next() {

		pointToPointRoutingTmp := ovsstruct.PointToPointRouting{}
		err = rows.StructScan(&pointToPointRoutingTmp)
		if err != nil {
			lg.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		getRequest := commonproto.GetRequest{}
		getRequest.UserEmail = in.UserEmail
		getRequest.RequestId = in.RequestId
		pointToPointRouting, err := lg.getPointToPointRoutingStruct(ctx, &getRequest, pointToPointRoutingTmp)
		if err != nil {
			lg.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		pointToPointRoutings = append(pointToPointRoutings, pointToPointRouting)

	}

	pointToPointRoutingsResponse := ovsproto.GetPointToPointRoutingsResponse{}
	if len(pointToPointRoutings) != 0 {
		next := pointToPointRoutings[len(pointToPointRoutings)-1].PointToPointRoutingD.Id
		next--
		nextc := common.EncodeCursor(next)
		pointToPointRoutingsResponse = ovsproto.GetPointToPointRoutingsResponse{PointToPointRoutings: pointToPointRoutings, NextCursor: nextc}
	} else {
		pointToPointRoutingsResponse = ovsproto.GetPointToPointRoutingsResponse{PointToPointRoutings: pointToPointRoutings, NextCursor: "0"}
	}
	return &pointToPointRoutingsResponse, nil
}

// GetPointToPointRouting - Get PointToPointRouting
func (lg *LegService) GetPointToPointRouting(ctx context.Context, inReq *ovsproto.GetPointToPointRoutingRequest) (*ovsproto.GetPointToPointRoutingResponse, error) {
	in := inReq.GetRequest
	uuid4byte, err := common.UUIDStrToBytes(in.Id)
	if err != nil {
		lg.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	nselectPointToPointRoutingsSQL := selectPointToPointRoutingsSQL + ` where uuid4 = ?`
	row := lg.DBService.DB.QueryRowxContext(ctx, nselectPointToPointRoutingsSQL, uuid4byte)
	pointToPointRoutingTmp := ovsstruct.PointToPointRouting{}
	err = row.StructScan(&pointToPointRoutingTmp)
	if err != nil {
		lg.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	getRequest := commonproto.GetRequest{}
	getRequest.UserEmail = in.UserEmail
	getRequest.RequestId = in.RequestId
	pointToPointRouting, err := lg.getPointToPointRoutingStruct(ctx, &getRequest, pointToPointRoutingTmp)
	if err != nil {
		lg.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	pointToPointRoutingResponse := ovsproto.GetPointToPointRoutingResponse{}
	pointToPointRoutingResponse.PointToPointRouting = pointToPointRouting
	return &pointToPointRoutingResponse, nil
}

// GetPointToPointRoutingByPk - Get PointToPointRouting By Primary key(Id)
func (lg *LegService) GetPointToPointRoutingByPk(ctx context.Context, inReq *ovsproto.GetPointToPointRoutingByPkRequest) (*ovsproto.GetPointToPointRoutingByPkResponse, error) {
	in := inReq.GetByIdRequest
	nselectPointToPointRoutingsSQL := selectPointToPointRoutingsSQL + ` where id = ?;`
	row := lg.DBService.DB.QueryRowxContext(ctx, nselectPointToPointRoutingsSQL, in.Id)
	pointToPointRoutingTmp := ovsstruct.PointToPointRouting{}
	err := row.StructScan(&pointToPointRoutingTmp)
	if err != nil {
		lg.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	getRequest := commonproto.GetRequest{}
	getRequest.UserEmail = in.UserEmail
	getRequest.RequestId = in.RequestId
	pointToPointRouting, err := lg.getPointToPointRoutingStruct(ctx, &getRequest, pointToPointRoutingTmp)
	if err != nil {
		lg.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	pointToPointRoutingResponse := ovsproto.GetPointToPointRoutingByPkResponse{}
	pointToPointRoutingResponse.PointToPointRouting = pointToPointRouting
	return &pointToPointRoutingResponse, nil
}

// getPointToPointRoutingStruct - Get pointToPointRouting
func (lg *LegService) getPointToPointRoutingStruct(ctx context.Context, in *commonproto.GetRequest, pointToPointRoutingTmp ovsstruct.PointToPointRouting) (*ovsproto.PointToPointRouting, error) {
	crUpdTime := new(commonproto.CrUpdTime)
	crUpdTime.CreatedAt = common.TimeToTimestamp(pointToPointRoutingTmp.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimeToTimestamp(pointToPointRoutingTmp.CrUpdTime.UpdatedAt)

	uuid4Str, err := common.UUIDBytesToStr(pointToPointRoutingTmp.PointToPointRoutingD.Uuid4)
	if err != nil {
		lg.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	pointToPointRoutingTmp.PointToPointRoutingD.IdS = uuid4Str

	pointToPointRouting := ovsproto.PointToPointRouting{PointToPointRoutingD: pointToPointRoutingTmp.PointToPointRoutingD, CrUpdUser: pointToPointRoutingTmp.CrUpdUser, CrUpdTime: crUpdTime}

	return &pointToPointRouting, nil
}
