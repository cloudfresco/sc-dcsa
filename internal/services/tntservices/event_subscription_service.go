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

// EventSubscriptionService - For accessing Transport Document services
type EventSubscriptionService struct {
	log               *zap.Logger
	DBService         *common.DBService
	RedisService      *common.RedisService
	UserServiceClient partyproto.UserServiceClient
	tntproto.UnimplementedEventSubscriptionServiceServer
}

// NewEventSubscriptionService - Create Transport Document service
func NewEventSubscriptionService(log *zap.Logger, dbOpt *common.DBService, redisOpt *common.RedisService, userServiceClient partyproto.UserServiceClient) *EventSubscriptionService {
	return &EventSubscriptionService{
		log:               log,
		DBService:         dbOpt,
		RedisService:      redisOpt,
		UserServiceClient: userServiceClient,
	}
}

// insertEventSubscriptionSQL - insert EventSubscriptionSQL query
const insertEventSubscriptionSQL = `insert into event_subscriptions
	  (
subscription_id,
callback_url,
document_reference,
equipment_reference,
transport_call_reference,
vessel_imo_number,
carrier_export_voyage_number,
universal_export_voyage_reference,
carrier_service_code,
universal_service_reference,
un_location_code,
secret,
status_code,
created_by_user_id,
updated_by_user_id,
created_at,
updated_at
  )
  values (:subscription_id,
:callback_url,
:document_reference,
:equipment_reference,
:transport_call_reference,
:vessel_imo_number,
:carrier_export_voyage_number,
:universal_export_voyage_reference,
:carrier_service_code,
:universal_service_reference,
:un_location_code,
:secret,
:status_code,
:created_by_user_id,
:updated_by_user_id,
:created_at,
:updated_at);`

// selectEventSubscriptionsSQL - select EventSubscriptionsSQL Query
const selectEventSubscriptionsSQL = `select 
  id,
  subscription_id,
  callback_url,
  document_reference,
  equipment_reference,
  transport_call_reference,
  vessel_imo_number,
  carrier_export_voyage_number,
  universal_export_voyage_reference,
  carrier_service_code,
  universal_service_reference,
  un_location_code,
  secret,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at from event_subscriptions`

// updateEventSubscriptionSQL - update EventSubscriptionSQL Query
const updateEventSubscriptionSQL = `update event_subscriptions set 
  callback_url = ?,
  document_reference = ?,
  equipment_reference = ?,
  transport_call_reference = ?,
  vessel_imo_number = ?,
  carrier_export_voyage_number = ?,
  universal_export_voyage_reference = ?,
  carrier_service_code = ?,
  universal_service_reference = ? where subscription_id = ?;`

// deleteEventSubscriptionSQL - delete EventSubscriptionSQL Query
const deleteEventSubscriptionSQL = `delete from event_subscriptions where subscription_id = ?;`

// GetEventSubscriptions - GetEventSubscriptions
func (ess *EventSubscriptionService) GetEventSubscriptions(ctx context.Context, in *tntproto.GetEventSubscriptionsRequest) (*tntproto.GetEventSubscriptionsResponse, error) {
	limit := in.GetLimit()
	nextCursor := in.GetNextCursor()
	if limit == "" {
		limit = ess.DBService.LimitSQLRows
	}
	query := "status_code = ?"
	if nextCursor == "" {
		query = query + " order by id desc " + " limit " + limit + ";"
	} else {
		nextCursor = common.DecodeCursor(nextCursor)
		query = query + " " + "and" + " " + "id <= " + nextCursor + " order by id desc " + " limit " + limit + ";"
	}

	eventSubscriptions := []*tntproto.EventSubscription{}

	nselectEventSubscriptionsSQL := selectEventSubscriptionsSQL + ` where ` + query

	rows, err := ess.DBService.DB.QueryxContext(ctx, nselectEventSubscriptionsSQL, "active")
	if err != nil {
		ess.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	for rows.Next() {

		eventSubscriptionTmp := tntstruct.EventSubscription{}
		err = rows.StructScan(&eventSubscriptionTmp)
		if err != nil {
			ess.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		getRequest := commonproto.GetRequest{}
		getRequest.UserEmail = in.UserEmail
		getRequest.RequestId = in.RequestId
		eventSubscription, err := ess.getEventSubscriptionStruct(ctx, &getRequest, eventSubscriptionTmp)
		if err != nil {
			ess.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		eventSubscriptions = append(eventSubscriptions, eventSubscription)

	}

	eventSubscriptionsResponse := tntproto.GetEventSubscriptionsResponse{}
	if len(eventSubscriptions) != 0 {
		next := eventSubscriptions[len(eventSubscriptions)-1].EventSubscriptionD.Id
		next--
		nextc := common.EncodeCursor(next)
		eventSubscriptionsResponse = tntproto.GetEventSubscriptionsResponse{EventSubscriptions: eventSubscriptions, NextCursor: nextc}
	} else {
		eventSubscriptionsResponse = tntproto.GetEventSubscriptionsResponse{EventSubscriptions: eventSubscriptions, NextCursor: "0"}
	}
	return &eventSubscriptionsResponse, nil
}

// CreateEventSubscription - Create EventSubscription
func (ess *EventSubscriptionService) CreateEventSubscription(ctx context.Context, in *tntproto.CreateEventSubscriptionRequest) (*tntproto.CreateEventSubscriptionResponse, error) {
	user, err := partyservice.GetUserWithNewContext(ctx, in.UserId, in.UserEmail, in.RequestId, ess.UserServiceClient)
	if err != nil {
		ess.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	ttime := common.GetTimeDetails()
	tn := common.TimeToTimestamp(ttime)

	eventSubscriptionD := tntproto.EventSubscriptionD{}
	eventSubscriptionD.SubscriptionId, err = common.GetUUIDBytes()
	if err != nil {
		ess.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	eventSubscriptionD.CallbackUrl = in.CallbackUrl
	eventSubscriptionD.DocumentReference = in.DocumentReference
	eventSubscriptionD.EquipmentReference = in.EquipmentReference
	eventSubscriptionD.TransportCallReference = in.TransportCallReference
	eventSubscriptionD.VesselImoNumber = in.VesselImoNumber
	eventSubscriptionD.CarrierExportVoyageNumber = in.CarrierExportVoyageNumber
	eventSubscriptionD.UniversalExportVoyageReference = in.UniversalExportVoyageReference
	eventSubscriptionD.CarrierServiceCode = in.CarrierServiceCode
	eventSubscriptionD.UniversalServiceReference = in.UniversalServiceReference
	eventSubscriptionD.UnLocationCode = in.UnLocationCode
	eventSubscriptionD.Secret = in.Secret

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	eventSubscription := tntproto.EventSubscription{EventSubscriptionD: &eventSubscriptionD, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	err = ess.insertEventSubscription(ctx, insertEventSubscriptionSQL, &eventSubscription, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		ess.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	eventSubscriptionResponse := tntproto.CreateEventSubscriptionResponse{}
	eventSubscriptionResponse.EventSubscription = &eventSubscription
	return &eventSubscriptionResponse, nil
}

// insertEventSubscription - Insert EventSubscription
func (ess *EventSubscriptionService) insertEventSubscription(ctx context.Context, insertEventSubscriptionSQL string, eventSubscription *tntproto.EventSubscription, userEmail string, requestID string) error {
	eventSubscriptionTmp, err := ess.CrEventSubscriptionStruct(ctx, eventSubscription, userEmail, requestID)
	if err != nil {
		ess.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	err = ess.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertEventSubscriptionSQL, eventSubscriptionTmp)
		if err != nil {
			ess.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			ess.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		eventSubscription.EventSubscriptionD.Id = uint32(uID)
		return nil
	})
	if err != nil {
		ess.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// CrEventSubscriptionStruct - process EventSubscription details
func (ess *EventSubscriptionService) CrEventSubscriptionStruct(ctx context.Context, eventSubscription *tntproto.EventSubscription, userEmail string, requestID string) (*tntstruct.EventSubscription, error) {
	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(eventSubscription.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(eventSubscription.CrUpdTime.UpdatedAt)

	eventSubscriptionTmp := tntstruct.EventSubscription{EventSubscriptionD: eventSubscription.EventSubscriptionD, CrUpdUser: eventSubscription.CrUpdUser, CrUpdTime: crUpdTime}

	return &eventSubscriptionTmp, nil
}

// FindEventSubscriptionByID - find By ID
func (ess *EventSubscriptionService) FindEventSubscriptionByID(ctx context.Context, inReq *tntproto.FindEventSubscriptionByIDRequest) (*tntproto.FindEventSubscriptionByIDResponse, error) {
	in := inReq.GetRequest
	uuid4byte, err := common.UUIDStrToBytes(in.Id)
	if err != nil {
		ess.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	nselectEventSubscriptionsSQL := selectEventSubscriptionsSQL + ` where uuid4 = ?;`
	row := ess.DBService.DB.QueryRowxContext(ctx, nselectEventSubscriptionsSQL, uuid4byte)
	eventSubscriptionTmp := tntstruct.EventSubscription{}
	err = row.StructScan(&eventSubscriptionTmp)
	if err != nil {
		ess.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	eventSubscription, err := ess.getEventSubscriptionStruct(ctx, in, eventSubscriptionTmp)
	if err != nil {
		ess.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	eventSubscriptionResponse := tntproto.FindEventSubscriptionByIDResponse{}
	eventSubscriptionResponse.EventSubscription = eventSubscription
	return &eventSubscriptionResponse, nil
}

// getEventSubscriptionStruct - Get EventSubscription header
func (ess *EventSubscriptionService) getEventSubscriptionStruct(ctx context.Context, in *commonproto.GetRequest, eventSubscriptionTmp tntstruct.EventSubscription) (*tntproto.EventSubscription, error) {
	crUpdTime := new(commonproto.CrUpdTime)
	crUpdTime.CreatedAt = common.TimeToTimestamp(eventSubscriptionTmp.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimeToTimestamp(eventSubscriptionTmp.CrUpdTime.UpdatedAt)

	eventSubscription := tntproto.EventSubscription{EventSubscriptionD: eventSubscriptionTmp.EventSubscriptionD, CrUpdUser: eventSubscriptionTmp.CrUpdUser, CrUpdTime: crUpdTime}

	return &eventSubscription, nil
}

// UpdateEventSubscription - Update EventSubscription
func (ess *EventSubscriptionService) UpdateEventSubscription(ctx context.Context, in *tntproto.UpdateEventSubscriptionRequest) (*tntproto.UpdateEventSubscriptionResponse, error) {
	uuid4byte, err := common.UUIDStrToBytes(in.SubscriptionId)
	if err != nil {
		ess.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	db := ess.DBService.DB
	stmt, err := db.PreparexContext(ctx, updateEventSubscriptionSQL)
	if err != nil {
		ess.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = ess.DBService.InsUpd(ctx, in.GetUserEmail(), in.GetRequestId(), func(tx *sqlx.Tx) error {
		_, err = tx.StmtxContext(ctx, stmt).ExecContext(ctx,
			in.CallbackUrl,
			in.DocumentReference,
			in.EquipmentReference,
			in.TransportCallReference,
			in.VesselImoNumber,
			in.CarrierExportVoyageNumber,
			in.UniversalExportVoyageReference,
			in.CarrierServiceCode,
			in.UniversalServiceReference,
			uuid4byte)
		if err != nil {
			ess.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			err1 := stmt.Close()
			if err1 != nil {
				ess.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err1))
				return err1
			}
			return err
		}
		return nil
	})

	if err != nil {
		ess.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = stmt.Close()
	if err != nil {
		ess.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	return &tntproto.UpdateEventSubscriptionResponse{}, nil
}

// DeleteEventSubscriptionByID - DeleteEventSubscriptionByID EventSubscription
func (ess *EventSubscriptionService) DeleteEventSubscriptionByID(ctx context.Context, inReq *tntproto.DeleteEventSubscriptionByIDRequest) (*tntproto.DeleteEventSubscriptionByIDResponse, error) {
	in := inReq.GetRequest
	uuid4byte, err := common.UUIDStrToBytes(in.Id)
	if err != nil {
		ess.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	db := ess.DBService.DB
	stmt, err := db.PreparexContext(ctx, deleteEventSubscriptionSQL)
	if err != nil {
		ess.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = ess.DBService.InsUpd(ctx, in.GetUserEmail(), in.GetRequestId(), func(tx *sqlx.Tx) error {
		_, err = tx.StmtxContext(ctx, stmt).ExecContext(ctx,
			uuid4byte)
		if err != nil {
			ess.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			err1 := stmt.Close()
			if err1 != nil {
				ess.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err1))
				return err1
			}
			return err
		}
		return nil
	})

	if err != nil {
		ess.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = stmt.Close()
	if err != nil {
		ess.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	return &tntproto.DeleteEventSubscriptionByIDResponse{}, nil
}
