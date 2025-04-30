package eblservices

import (
	"context"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	commonproto "github.com/cloudfresco/sc-dcsa/internal/protogen/common/v1"
	eblproto "github.com/cloudfresco/sc-dcsa/internal/protogen/ebl/v1"
	partyservice "github.com/cloudfresco/sc-dcsa/internal/services/partyservices"
	commonstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/common/v1"
	eblstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/ebl/v1"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// insertTransportSQL - Insert TransportSQL Query
const insertTransportSQL = `insert into transports
	  ( 
  uuid4,
  transport_reference,
  transport_name,
  load_transport_call_id,
  discharge_transport_call_id,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at
  )
  values (:uuid4,
:transport_reference,
:transport_name,
:load_transport_call_id,
:discharge_transport_call_id,
:status_code,
:created_by_user_id,
:updated_by_user_id,
:created_at,
:updated_at);`

// selectTransportSQL - select TransportSQL Query
const selectTransportsSQL = `select 
  id,
  uuid4,
  transport_reference,
  transport_name,
  load_transport_call_id,
  discharge_transport_call_id,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at from transports`

// CreateTransport - Create  Transport
func (ss *ShippingService) CreateTransport(ctx context.Context, in *eblproto.CreateTransportRequest) (*eblproto.CreateTransportResponse, error) {
	user, err := partyservice.GetUserWithNewContext(ctx, in.UserId, in.UserEmail, in.RequestId, ss.UserServiceClient)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	ttime := common.GetTimeDetails()
	tn := common.TimeToTimestamp(ttime)

	transportD := eblproto.TransportD{}
	transportD.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	transportD.TransportReference = in.TransportReference
	transportD.TransportName = in.TransportName
	transportD.LoadTransportCallId = in.LoadTransportCallId
	transportD.DischargeTransportCallId = in.DischargeTransportCallId

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	transport := eblproto.Transport{TransportD: &transportD, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	err = ss.insertTransport(ctx, insertTransportSQL, &transport, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	transportResponse := eblproto.CreateTransportResponse{}
	transportResponse.Transport = &transport
	return &transportResponse, nil
}

// insertTransport - Insert Transport
func (ss *ShippingService) insertTransport(ctx context.Context, insertTransportSQL string, transport *eblproto.Transport, userEmail string, requestID string) error {
	transportTmp, err := ss.CrTransportStruct(ctx, transport, userEmail, requestID)
	if err != nil {
		ss.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}

	err = ss.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertTransportSQL, transportTmp)
		if err != nil {
			ss.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			ss.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		transport.TransportD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(transport.TransportD.Uuid4)
		if err != nil {
			ss.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		transport.TransportD.IdS = uuid4Str
		return nil
	})
	if err != nil {
		ss.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// CrTransportStruct - process Transport details
func (ss *ShippingService) CrTransportStruct(ctx context.Context, transport *eblproto.Transport, userEmail string, requestID string) (*eblstruct.Transport, error) {
	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(transport.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(transport.CrUpdTime.UpdatedAt)

	transportTmp := eblstruct.Transport{TransportD: transport.TransportD, CrUpdUser: transport.CrUpdUser, CrUpdTime: crUpdTime}

	return &transportTmp, nil
}

// GetTransports - Get  Transports
func (ss *ShippingService) GetTransports(ctx context.Context, in *eblproto.GetTransportsRequest) (*eblproto.GetTransportsResponse, error) {
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

	transports := []*eblproto.Transport{}

	nselectTransportsSQL := selectTransportsSQL + query

	rows, err := ss.DBService.DB.QueryxContext(ctx, nselectTransportsSQL)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	for rows.Next() {

		transportTmp := eblstruct.Transport{}
		err = rows.StructScan(&transportTmp)
		if err != nil {
			ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		getRequest := commonproto.GetRequest{}
		getRequest.UserEmail = in.UserEmail
		getRequest.RequestId = in.RequestId
		transport, err := ss.getTransportStruct(ctx, &getRequest, transportTmp)
		if err != nil {
			ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		transports = append(transports, transport)

	}

	transportsResponse := eblproto.GetTransportsResponse{}
	if len(transports) != 0 {
		next := transports[len(transports)-1].TransportD.Id
		next--
		nextc := common.EncodeCursor(next)
		transportsResponse = eblproto.GetTransportsResponse{Transports: transports, NextCursor: nextc}
	} else {
		transportsResponse = eblproto.GetTransportsResponse{Transports: transports, NextCursor: "0"}
	}
	return &transportsResponse, nil
}

// GetTransport - Get Transport
func (ss *ShippingService) GetTransport(ctx context.Context, inReq *eblproto.GetTransportRequest) (*eblproto.GetTransportResponse, error) {
	in := inReq.GetRequest
	uuid4byte, err := common.UUIDStrToBytes(in.Id)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	nselectTransportsSQL := selectTransportsSQL + ` where uuid4 = ?;`
	row := ss.DBService.DB.QueryRowxContext(ctx, nselectTransportsSQL, uuid4byte)
	transportTmp := eblstruct.Transport{}
	err = row.StructScan(&transportTmp)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	getRequest := commonproto.GetRequest{}
	getRequest.UserEmail = in.UserEmail
	getRequest.RequestId = in.RequestId
	transport, err := ss.getTransportStruct(ctx, &getRequest, transportTmp)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	transportResponse := eblproto.GetTransportResponse{}
	transportResponse.Transport = transport
	return &transportResponse, nil
}

// GetTransportByPk - Get Transport By Primary key(Id)
func (ss *ShippingService) GetTransportByPk(ctx context.Context, inReq *eblproto.GetTransportByPkRequest) (*eblproto.GetTransportByPkResponse, error) {
	in := inReq.GetByIdRequest
	nselectTransportsSQL := selectTransportsSQL + ` where id = ?;`
	row := ss.DBService.DB.QueryRowxContext(ctx, nselectTransportsSQL, in.Id)
	transportTmp := eblstruct.Transport{}
	err := row.StructScan(&transportTmp)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	getRequest := commonproto.GetRequest{}
	getRequest.UserEmail = in.UserEmail
	getRequest.RequestId = in.RequestId
	transport, err := ss.getTransportStruct(ctx, &getRequest, transportTmp)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	transportResponse := eblproto.GetTransportByPkResponse{}
	transportResponse.Transport = transport
	return &transportResponse, nil
}

// GetTransportStruct - Get Transport header
func (ss *ShippingService) getTransportStruct(ctx context.Context, in *commonproto.GetRequest, transportTmp eblstruct.Transport) (*eblproto.Transport, error) {
	crUpdTime := new(commonproto.CrUpdTime)
	crUpdTime.CreatedAt = common.TimeToTimestamp(transportTmp.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimeToTimestamp(transportTmp.CrUpdTime.UpdatedAt)

	uuid4Str, err := common.UUIDBytesToStr(transportTmp.TransportD.Uuid4)
	if err != nil {
		ss.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	transportTmp.TransportD.IdS = uuid4Str

	transport := eblproto.Transport{TransportD: transportTmp.TransportD, CrUpdUser: transportTmp.CrUpdUser, CrUpdTime: crUpdTime}

	return &transport, nil
}
