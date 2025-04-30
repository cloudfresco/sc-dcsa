package tntservices

import (
	"context"
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

// EventService - For accessing Transport Document services
type EventService struct {
	log               *zap.Logger
	DBService         *common.DBService
	RedisService      *common.RedisService
	UserServiceClient partyproto.UserServiceClient
	tntproto.UnimplementedEventServiceServer
}

// NewEventService - Create Transport Document service
func NewEventService(log *zap.Logger, dbOpt *common.DBService, redisOpt *common.RedisService, userServiceClient partyproto.UserServiceClient) *EventService {
	return &EventService{
		log:               log,
		DBService:         dbOpt,
		RedisService:      redisOpt,
		UserServiceClient: userServiceClient,
	}
}

// InsertEquipmentEventSQL - insert EquipmentEventSQL query
const InsertEquipmentEventSQL = `insert into equipment_events
	  ( 
event_id_s,
event_classifier_code,
equipment_event_type_code,
equipment_reference,
empty_indicator_code,
transport_call_id,
event_location,
event_created_date_time,
event_date_time
  )
  values (:event_id_s,
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
  event_id_s,
  event_classifier_code,
  equipment_event_type_code,
  equipment_reference,
  empty_indicator_code,
  transport_call_id,
  event_location,  
  event_created_date_time,
  event_date_time from equipment_events`

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
	equipmentEvents, err := es.ProcessEquipmentEvent(ctx, in)
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = es.insertEquipmentEvent(ctx, InsertEquipmentEventSQL, equipmentEvents, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	equipmentEventResponse := tntproto.CreateEquipmentEventResponse{}
	return &equipmentEventResponse, nil
}

// ProcessEquipmentEvent - Process EquipmentEvent
func (es *EventService) ProcessEquipmentEvent(ctx context.Context, inReq *tntproto.CreateEquipmentEventRequest) ([]*tntproto.EquipmentEvent, error) {
	equipmentEvents := []*tntproto.EquipmentEvent{}
	for _, in := range inReq.EquipmentEventRequests {
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
		equipmentEvents = append(equipmentEvents, &equipmentEvent)
	}
	return equipmentEvents, nil
}

// insertEquipmentEvent - Insert EquipmentEvent
func (es *EventService) insertEquipmentEvent(ctx context.Context, insertEquipmentEventSQL string, equipmentEvents []*tntproto.EquipmentEvent, userEmail string, requestID string) error {
	equipmentEventTmps := []*tntstruct.EquipmentEvent{}
	for _, equipmentEvent := range equipmentEvents {
		equipmentEventTmp, err := es.CrEquipmentEventStruct(ctx, equipmentEvent, userEmail, requestID)
		if err != nil {
			es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		equipmentEventTmps = append(equipmentEventTmps, equipmentEventTmp)
	}

	err := es.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		_, err := tx.NamedExecContext(ctx, insertEquipmentEventSQL, equipmentEventTmps)
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

// LoadEquipmentRelatedEntities - Get EquipmentEvents
func (es *EventService) LoadEquipmentRelatedEntities(ctx context.Context, in *tntproto.LoadEquipmentRelatedEntitiesRequest) (*tntproto.LoadEquipmentRelatedEntitiesResponse, error) {
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

	equipmentEventResponse := tntproto.LoadEquipmentRelatedEntitiesResponse{}
	if len(equipmentEvents) != 0 {
		next := equipmentEvents[len(equipmentEvents)-1].EquipmentEventD.Id
		next--
		nextc := common.EncodeCursor(next)
		equipmentEventResponse = tntproto.LoadEquipmentRelatedEntitiesResponse{EquipmentEvents: equipmentEvents, NextCursor: nextc}
	} else {
		equipmentEventResponse = tntproto.LoadEquipmentRelatedEntitiesResponse{EquipmentEvents: equipmentEvents, NextCursor: "0"}
	}
	return &equipmentEventResponse, nil
}

// getEquipmentEventStruct - Get EquipmentEvent header
func (es *EventService) getEquipmentEventStruct(ctx context.Context, in *commonproto.GetRequest, equipmentEventTmp tntstruct.EquipmentEvent) (*tntproto.EquipmentEvent, error) {
	equipmentEventT := new(tntproto.EquipmentEventT)
	equipmentEventT.EventCreatedDateTime = common.TimeToTimestamp(equipmentEventTmp.EquipmentEventT.EventCreatedDateTime)
	equipmentEventT.EventDateTime = common.TimeToTimestamp(equipmentEventTmp.EquipmentEventT.EventDateTime)

	equipmentEvent := tntproto.EquipmentEvent{EquipmentEventD: equipmentEventTmp.EquipmentEventD, EquipmentEventT: equipmentEventT}

	return &equipmentEvent, nil
}
