package eventcoreservices

import (
	"context"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	commonproto "github.com/cloudfresco/sc-dcsa/internal/protogen/common/v1"
	eventcoreproto "github.com/cloudfresco/sc-dcsa/internal/protogen/eventcore/v1"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"
	partyservice "github.com/cloudfresco/sc-dcsa/internal/services/partyservices"
	commonstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/common/v1"
	eventcorestruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/eventcore/v1"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// TransportCallService - For accessing TransportCall services
type TransportCallService struct {
	log               *zap.Logger
	DBService         *common.DBService
	RedisService      *common.RedisService
	UserServiceClient partyproto.UserServiceClient
	eventcoreproto.UnimplementedTransportCallServiceServer
}

// NewTransportCallService - Create TransportCall service
func NewTransportCallService(log *zap.Logger, dbOpt *common.DBService, redisOpt *common.RedisService, userServiceClient partyproto.UserServiceClient) *TransportCallService {
	return &TransportCallService{
		log:               log,
		DBService:         dbOpt,
		RedisService:      redisOpt,
		UserServiceClient: userServiceClient,
	}
}

// insertTransportCallSQL - insert TransportCallSQL query
const insertTransportCallSQL = `insert into transport_calls
	  ( 
  uuid4,
  transport_call_reference,
  transport_call_sequence_number,
  facility_id,
  facility_type_code,
  other_facility,
  location_id,
  mode_of_transport_code,
  vessel_id,
  import_voyage_id,
  export_voyage_id,
  port_call_status_code,
  port_visit_reference,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at
  )
  values (:uuid4,
:transport_call_reference,
:transport_call_sequence_number,
:facility_id,
:facility_type_code,
:other_facility,
:location_id,
:mode_of_transport_code,
:vessel_id,
:import_voyage_id,
:export_voyage_id,
:port_call_status_code,
:port_visit_reference,
:status_code,
:created_by_user_id,
:updated_by_user_id,
:created_at,
:updated_at);`

// selectTransportCallsSQL - select TransportCallsSQL query
const selectTransportCallsSQL = `select 
  id,
  uuid4,
  transport_call_reference,
  transport_call_sequence_number,
  facility_id,
  facility_type_code,
  other_facility,
  location_id,
  mode_of_transport_code,
  vessel_id,
  import_voyage_id,
  export_voyage_id,
  port_call_status_code,
  port_visit_reference,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at from transport_calls`

// CreateTransportCall - Create TransportCall
func (tcs *TransportCallService) CreateTransportCall(ctx context.Context, in *eventcoreproto.CreateTransportCallRequest) (*eventcoreproto.CreateTransportCallResponse, error) {
	user, err := partyservice.GetUserWithNewContext(ctx, in.UserId, in.UserEmail, in.RequestId, tcs.UserServiceClient)
	if err != nil {
		tcs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	ttime := common.GetTimeDetails()
	tn := common.TimeToTimestamp(ttime)

	transportCallD := eventcoreproto.TransportCallD{}
	transportCallD.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		tcs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	transportCallD.TransportCallReference = in.TransportCallReference
	transportCallD.TransportCallSequenceNumber = in.TransportCallSequenceNumber
	transportCallD.FacilityId = in.FacilityId
	transportCallD.FacilityTypeCode = in.FacilityTypeCode
	transportCallD.OtherFacility = in.OtherFacility
	transportCallD.LocationId = in.LocationId
	transportCallD.ModeOfTransportCode = in.ModeOfTransportCode
	transportCallD.VesselId = in.VesselId
	transportCallD.ImportVoyageId = in.ImportVoyageId
	transportCallD.ExportVoyageId = in.ExportVoyageId
	transportCallD.PortCallStatusCode = in.PortCallStatusCode
	transportCallD.PortVisitReference = in.PortVisitReference

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	transportCall := eventcoreproto.TransportCall{TransportCallD: &transportCallD, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	err = tcs.insertTransportCall(ctx, insertTransportCallSQL, &transportCall, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		tcs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	transCallResponse := eventcoreproto.CreateTransportCallResponse{}
	transCallResponse.TransportCall = &transportCall
	return &transCallResponse, nil
}

// insertTransportCall - Insert Transport Call
func (tcs *TransportCallService) insertTransportCall(ctx context.Context, insertTransportCallSQL string, transportCall *eventcoreproto.TransportCall, userEmail string, requestID string) error {
	transportCallTmp, err := tcs.crTransportCallStruct(ctx, transportCall, userEmail, requestID)
	if err != nil {
		tcs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}

	err = tcs.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertTransportCallSQL, transportCallTmp)
		if err != nil {
			tcs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			tcs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		transportCall.TransportCallD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(transportCall.TransportCallD.Uuid4)
		if err != nil {
			tcs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		transportCall.TransportCallD.IdS = uuid4Str
		return nil
	})
	if err != nil {
		tcs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// crTransportCallStruct - process TransportCall details
func (tcs *TransportCallService) crTransportCallStruct(ctx context.Context, transportCall *eventcoreproto.TransportCall, userEmail string, requestID string) (*eventcorestruct.TransportCall, error) {
	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(transportCall.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(transportCall.CrUpdTime.UpdatedAt)
	transportCallTmp := eventcorestruct.TransportCall{TransportCallD: transportCall.TransportCallD, CrUpdUser: transportCall.CrUpdUser, CrUpdTime: crUpdTime}

	return &transportCallTmp, nil
}

// GetTransportCalls - Get  TransportCalls
func (tcs *TransportCallService) GetTransportCalls(ctx context.Context, in *eventcoreproto.GetTransportCallsRequest) (*eventcoreproto.GetTransportCallsResponse, error) {
	limit := in.GetLimit()
	nextCursor := in.GetNextCursor()
	if limit == "" {
		limit = tcs.DBService.LimitSQLRows
	}
	query := ""
	if nextCursor == "" {
		query = " order by id desc " + " limit " + limit + ";"
	} else {
		nextCursor = common.DecodeCursor(nextCursor)
		query = " where id <= " + nextCursor + " order by id desc " + " limit " + limit + ";"
	}

	transportCalls := []*eventcoreproto.TransportCall{}

	nselectTransportCallsSQL := selectTransportCallsSQL + query

	rows, err := tcs.DBService.DB.QueryxContext(ctx, nselectTransportCallsSQL)
	if err != nil {
		tcs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	for rows.Next() {

		transportCallTmp := eventcorestruct.TransportCall{}
		err = rows.StructScan(&transportCallTmp)
		if err != nil {
			tcs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		getRequest := commonproto.GetRequest{}
		getRequest.UserEmail = in.UserEmail
		getRequest.RequestId = in.RequestId
		transportCall, err := tcs.getTransportCallStruct(ctx, &getRequest, transportCallTmp)
		if err != nil {
			tcs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		transportCalls = append(transportCalls, transportCall)

	}

	transCallResponse := eventcoreproto.GetTransportCallsResponse{}
	if len(transportCalls) != 0 {
		next := transportCalls[len(transportCalls)-1].TransportCallD.Id
		next--
		nextc := common.EncodeCursor(next)
		transCallResponse = eventcoreproto.GetTransportCallsResponse{TransportCalls: transportCalls, NextCursor: nextc}
	} else {
		transCallResponse = eventcoreproto.GetTransportCallsResponse{TransportCalls: transportCalls, NextCursor: "0"}
	}
	return &transCallResponse, nil
}

// FindTransportCall - Find TransportCall
func (tcs *TransportCallService) FindTransportCall(ctx context.Context, inReq *eventcoreproto.FindTransportCallRequest) (*eventcoreproto.FindTransportCallResponse, error) {
	in := inReq.GetRequest
	uuid4byte, err := common.UUIDStrToBytes(in.Id)
	if err != nil {
		tcs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	nselectTransportCallsSQL := selectTransportCallsSQL + ` where uuid4 = ?;`
	row := tcs.DBService.DB.QueryRowxContext(ctx, nselectTransportCallsSQL, uuid4byte)
	transportCallTmp := eventcorestruct.TransportCall{}
	err = row.StructScan(&transportCallTmp)
	if err != nil {
		tcs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	getRequest := commonproto.GetRequest{}
	getRequest.UserEmail = in.UserEmail
	getRequest.RequestId = in.RequestId
	transportCall, err := tcs.getTransportCallStruct(ctx, &getRequest, transportCallTmp)
	if err != nil {
		tcs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	transCallResponse := eventcoreproto.FindTransportCallResponse{}
	transCallResponse.TransportCall = transportCall
	return &transCallResponse, nil
}

// GetTransportCallByPk - Get TransportCall By Primary key(Id)
func (tcs *TransportCallService) GetTransportCallByPk(ctx context.Context, inReq *eventcoreproto.GetTransportCallByPkRequest) (*eventcoreproto.GetTransportCallByPkResponse, error) {
	in := inReq.GetByIdRequest
	nselectTransportCallsSQL := selectTransportCallsSQL + ` where id = ?;`
	row := tcs.DBService.DB.QueryRowxContext(ctx, nselectTransportCallsSQL, in.Id)
	transportCallTmp := eventcorestruct.TransportCall{}
	err := row.StructScan(&transportCallTmp)
	if err != nil {
		tcs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	getRequest := commonproto.GetRequest{}
	getRequest.UserEmail = in.UserEmail
	getRequest.RequestId = in.RequestId
	transportCall, err := tcs.getTransportCallStruct(ctx, &getRequest, transportCallTmp)
	if err != nil {
		tcs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	transCallResponse := eventcoreproto.GetTransportCallByPkResponse{}
	transCallResponse.TransportCall = transportCall
	return &transCallResponse, nil
}

// getTransportCallStruct - Get TransportCall struct
func (tcs *TransportCallService) getTransportCallStruct(ctx context.Context, in *commonproto.GetRequest, transportCallTmp eventcorestruct.TransportCall) (*eventcoreproto.TransportCall, error) {
	crUpdTime := new(commonproto.CrUpdTime)
	crUpdTime.CreatedAt = common.TimeToTimestamp(transportCallTmp.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimeToTimestamp(transportCallTmp.CrUpdTime.UpdatedAt)

	uuid4Str, err := common.UUIDBytesToStr(transportCallTmp.TransportCallD.Uuid4)
	if err != nil {
		tcs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	transportCallTmp.TransportCallD.IdS = uuid4Str
	transportCall := eventcoreproto.TransportCall{TransportCallD: transportCallTmp.TransportCallD, CrUpdUser: transportCallTmp.CrUpdUser, CrUpdTime: crUpdTime}
	return &transportCall, nil
}
