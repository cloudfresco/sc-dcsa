package tntservices

import (
	"context"
	"errors"
	"net"
	"os"
	"time"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	"github.com/cloudfresco/sc-dcsa/internal/config"
	commonproto "github.com/cloudfresco/sc-dcsa/internal/protogen/common/v1"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"
	tntproto "github.com/cloudfresco/sc-dcsa/internal/protogen/tnt/v3"
	tntstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/tnt/v3"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// InsertEquipmentEventSQL - insert EquipmentEventSQL query
const InsertEquipmentEventSQL = `insert into equipment_events
	  ( 
event_id,
event_classifier_code,
equipment_event_type_code,
equipment_reference,
empty_indicator_code,
transport_call_id,
event_location,
event_created_date_time,
event_date_time
  )
  values (:event_id,
:event_classifier_code,
:equipment_event_type_code,
:equipment_reference,
:empty_indicator_code,
:transport_call_id,
:event_location,
:event_created_date_time,
:event_date_time);`

// selectEquipmentEventSQL - select EquipmentEventSQL Query
const selectEquipmentEventsSQL = `select 
  id,
  event_id,
  event_classifier_code,
  equipment_event_type_code,
  equipment_reference,
  empty_indicator_code,
  transport_call_id,
  event_location,  
  event_created_date_time,
  event_date_time from equipment_events`

// InsertEquipmentEventTypeSQL - insert EquipmentEventTypeSQL query
const InsertEquipmentEventTypeSQL = `insert into equipment_event_types
	  ( 
equipment_event_type_code,
equipment_event_type_name,
equipment_event_type_description,
equipment_event_id
  )
  values (:equipment_event_type_code,
:equipment_event_type_name,
:equipment_event_type_description,
:equipment_event_id);`

// StartTntServer - Start Tnt server
func StartTntServer(log *zap.Logger, isTest bool, pwd string, dbOpt *config.DBOptions, redisOpt *config.RedisOptions, mailerOpt *config.MailerOptions, grpcServerOpt *config.GrpcServerOptions, jwtOpt *config.JWTOptions, oauthOpt *config.OauthOptions, userOpt *config.UserOptions, uptraceOpt *config.UptraceOptions, dbService *common.DBService, redisService *common.RedisService, mailerService common.MailerIntf) {
	common.SetJWTOpt(jwtOpt)

	creds, err := common.GetSrvCred(log, isTest, pwd, grpcServerOpt)
	if err != nil {
		os.Exit(1)
	}

	userCreds, err := common.GetClientCred(log, isTest, pwd, grpcServerOpt)
	if err != nil {
		os.Exit(1)
	}

	var srvOpts []grpc.ServerOption

	userConn, err := grpc.NewClient(grpcServerOpt.GrpcUserServerPort, grpc.WithTransportCredentials(userCreds), grpc.WithStatsHandler(otelgrpc.NewClientHandler()))
	if err != nil {
		log.Error("Error", zap.Error(err))
		os.Exit(1)
	}

	srvOpts = append(srvOpts, grpc.Creds(creds))

	srvOpts = append(srvOpts, grpc.StatsHandler(otelgrpc.NewServerHandler()))

	uc := partyproto.NewUserServiceClient(userConn)
	eventService := NewEventService(log, dbService, redisService, uc)
	eventSubscriptionService := NewEventSubscriptionService(log, dbService, redisService, uc)
	sealService := NewSealService(log, dbService, redisService, uc)

	lis, err := net.Listen("tcp", grpcServerOpt.GrpcTntServerPort)
	if err != nil {
		log.Error("Error", zap.Error(err))
		os.Exit(1)
	}

	srv := grpc.NewServer(srvOpts...)
	tntproto.RegisterEventServiceServer(srv, eventService)
	tntproto.RegisterEventSubscriptionServiceServer(srv, eventSubscriptionService)
	tntproto.RegisterSealServiceServer(srv, sealService)

	if err := srv.Serve(lis); err != nil {
		log.Error("Error", zap.Error(err))
		os.Exit(1)
	}
}

// CreateEquipmentEvent - Create EquipmentEvent
func (es *EventService) CreateEquipmentEvent(ctx context.Context, in *tntproto.CreateEquipmentEventRequest) (*tntproto.CreateEquipmentEventResponse, error) {
	equipmentEvent, err := es.ProcessEquipmentEvent(ctx, in)
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = es.insertEquipmentEvent(ctx, InsertEquipmentEventSQL, equipmentEvent, InsertEquipmentEventTypeSQL, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	equipmentEventResponse := tntproto.CreateEquipmentEventResponse{}
	equipmentEventResponse.EquipmentEvent = equipmentEvent
	return &equipmentEventResponse, nil
}

// ProcessEquipmentEvent - Process EquipmentEvent
func (es *EventService) ProcessEquipmentEvent(ctx context.Context, in *tntproto.CreateEquipmentEventRequest) (*tntproto.EquipmentEvent, error) {
	eventCreatedDateTime, err := time.Parse(common.Layout, in.EventCreatedDateTime)
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	eventDateTime, err := time.Parse(common.Layout, in.EventDateTime)
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	equipmentEventD := tntproto.EquipmentEventD{}
	equipmentEventD.EventIdS, err = common.GetUUIDBytes()
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	equipmentEventD.EventClassifierCode = in.EventClassifierCode
	equipmentEventD.EquipmentEventTypeCode = in.EquipmentEventTypeCode
	equipmentEventD.EquipmentReference = in.EquipmentReference
	equipmentEventD.EmptyIndicatorCode = in.EmptyIndicatorCode
	equipmentEventD.TransportCallId = in.TransportCallId
	equipmentEventD.EventLocation = in.EventLocation

	equipmentEventT := tntproto.EquipmentEventT{}
	equipmentEventT.EventCreatedDateTime = common.TimeToTimestamp(eventCreatedDateTime.UTC().Truncate(time.Second))
	equipmentEventT.EventDateTime = common.TimeToTimestamp(eventDateTime.UTC().Truncate(time.Second))

	equipmentEvent := tntproto.EquipmentEvent{EquipmentEventD: &equipmentEventD, EquipmentEventT: &equipmentEventT}

	inEquipmentEventType := in.EquipmentEventType
	equipmentEventType := tntproto.EquipmentEventType{}
	equipmentEventType.EquipmentEventTypeCode = inEquipmentEventType.EquipmentEventTypeCode
	equipmentEventType.EquipmentEventTypeName = inEquipmentEventType.EquipmentEventTypeName
	equipmentEventType.EquipmentEventTypeDescription = inEquipmentEventType.EquipmentEventTypeDescription
	equipmentEvent.EquipmentEventType = &equipmentEventType
	return &equipmentEvent, nil
}

// insertEquipmentEvent - Insert EquipmentEvent
func (es *EventService) insertEquipmentEvent(ctx context.Context, insertEquipmentEventSQL string, equipmentEvent *tntproto.EquipmentEvent, insertEquipmentEventTypeSQL string, userEmail string, requestID string) error {
	equipmentEventTmp, err := es.CrEquipmentEventStruct(ctx, equipmentEvent, userEmail, requestID)
	if err != nil {
		es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	err = es.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertEquipmentEventSQL, equipmentEventTmp)
		if err != nil {
			es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		equipmentEvent.EquipmentEventD.Id = uint32(uID)

		equipmentEvent.EquipmentEventType.EquipmentEventId = uint32(uID)
		_, err = tx.NamedExecContext(ctx, insertEquipmentEventTypeSQL, equipmentEvent.EquipmentEventType)

		if err != nil {
			es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		return nil
	})
	if err != nil {
		es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// CrEquipmentEventStruct - process EquipmentEvent details
func (es *EventService) CrEquipmentEventStruct(ctx context.Context, equipmentEvent *tntproto.EquipmentEvent, userEmail string, requestID string) (*tntstruct.EquipmentEvent, error) {
	equipmentEventT := new(tntstruct.EquipmentEventT)
	equipmentEventT.EventCreatedDateTime = common.TimestampToTime(equipmentEvent.EquipmentEventT.EventCreatedDateTime)
	equipmentEventT.EventDateTime = common.TimestampToTime(equipmentEvent.EquipmentEventT.EventDateTime)

	equipmentEventTmp := tntstruct.EquipmentEvent{EquipmentEventD: equipmentEvent.EquipmentEventD, EquipmentEventT: equipmentEventT}

	return &equipmentEventTmp, nil
}

// LoadRelatedEntities - Get EquipmentEvents
func (es *EventService) LoadRelatedEntities(ctx context.Context, in *tntproto.LoadRelatedEntitiesRequest) (*tntproto.LoadRelatedEntitiesResponse, error) {
	select {
	case <-ctx.Done():
		err := errors.New("Client closed connection")
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	default:
		limit := in.GetLimit()
		nextCursor := in.GetNextCursor()
		if limit == "" {
			limit = es.DBService.LimitSQLRows
		}
		query := ""
		if nextCursor == "" {
			query = " order by id desc " + " limit " + limit + ";"
		} else {
			nextCursor = common.DecodeCursor(nextCursor)
			query = " where id <= " + nextCursor + " order by id desc " + " limit " + limit + ";"
		}

		equipmentEvents := []*tntproto.EquipmentEvent{}

		nselectEquipmentEventsSQL := selectEquipmentEventsSQL + query

		rows, err := es.DBService.DB.QueryxContext(ctx, nselectEquipmentEventsSQL)
		if err != nil {
			es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		for rows.Next() {

			equipmentEventTmp := tntstruct.EquipmentEvent{}
			err = rows.StructScan(&equipmentEventTmp)
			if err != nil {
				es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
				return nil, err
			}

			getRequest := commonproto.GetRequest{}
			getRequest.UserEmail = in.UserEmail
			getRequest.RequestId = in.RequestId
			equipmentEvent, err := es.getEquipmentEventStruct(ctx, &getRequest, equipmentEventTmp)
			if err != nil {
				es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
				return nil, err
			}
			equipmentEvents = append(equipmentEvents, equipmentEvent)

		}

		equipmentEventResponse := tntproto.LoadRelatedEntitiesResponse{}
		if len(equipmentEvents) != 0 {
			next := equipmentEvents[len(equipmentEvents)-1].EquipmentEventD.Id
			next--
			nextc := common.EncodeCursor(next)
			equipmentEventResponse = tntproto.LoadRelatedEntitiesResponse{EquipmentEvents: equipmentEvents, NextCursor: nextc}
		} else {
			equipmentEventResponse = tntproto.LoadRelatedEntitiesResponse{EquipmentEvents: equipmentEvents, NextCursor: "0"}
		}
		return &equipmentEventResponse, nil
	}
}

// getEquipmentEventStruct - Get EquipmentEvent header
func (es *EventService) getEquipmentEventStruct(ctx context.Context, in *commonproto.GetRequest, equipmentEventTmp tntstruct.EquipmentEvent) (*tntproto.EquipmentEvent, error) {
	equipmentEventT := new(tntproto.EquipmentEventT)
	equipmentEventT.EventCreatedDateTime = common.TimeToTimestamp(equipmentEventTmp.EquipmentEventT.EventCreatedDateTime)
	equipmentEventT.EventDateTime = common.TimeToTimestamp(equipmentEventTmp.EquipmentEventT.EventDateTime)

	equipmentEvent := tntproto.EquipmentEvent{EquipmentEventD: equipmentEventTmp.EquipmentEventD, EquipmentEventT: equipmentEventT}

	return &equipmentEvent, nil
}
