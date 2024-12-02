package tntservices

import (
	"context"
	"errors"
	"time"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	commonproto "github.com/cloudfresco/sc-dcsa/internal/proto-gen/common/v1"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/proto-gen/party/v1"
	tntproto "github.com/cloudfresco/sc-dcsa/internal/proto-gen/tnt/v3"
	tntstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/tnt/v3"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
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

const InsertEventSQL = ``

const InsertEventClassifierSQL = ``

// selectEventsSQL - select EventsSQL Query
const selectEventsSQL = `select 
  id,
  event_id,
  event_classifier_code,
  event_created_date_time,
  event_date_time from events`

// CreateEvent - Create Event
func (es *EventService) CreateEvent(ctx context.Context, in *tntproto.CreateEventRequest) (*tntproto.CreateEventResponse, error) {
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

	eventD := tntproto.EventD{}
	eventD.EventId, err = common.GetUUIDBytes()
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	eventD.EventClassifierCode = in.EventClassifierCode

	eventT := tntproto.EventT{}
	eventT.EventCreatedDateTime = common.TimeToTimestamp(eventCreatedDateTime.UTC().Truncate(time.Second))
	eventT.EventDateTime = common.TimeToTimestamp(eventDateTime.UTC().Truncate(time.Second))

	event := tntproto.Event{EventD: &eventD, EventT: &eventT}

	eventClassifier, err := es.ProcessEventClassifier(ctx, in.EventClassifier)
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	equipmentEvent := &tntproto.EquipmentEvent{}
	operationsEvent := &tntproto.OperationsEvent{}
	shipmentEvent := &tntproto.ShipmentEvent{}
	transportEvent := &tntproto.TransportEvent{}

	if in.EventType == "Equipment" {
		equipmentEvent, err = es.ProcessEquipmentEvent(ctx, in.EquipmentEvent)
		if err != nil {
			es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
	} else if in.EventType == "Operations" {

		operationsEvent, err = es.ProcessOperationsEvent(ctx, in.OperationsEvent)
		if err != nil {
			es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
	} else if in.EventType == "Shipment" {

		shipmentEvent, err = es.ProcessShipmentEvent(ctx, in.ShipmentEvent)
		if err != nil {
			es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
	} else if in.EventType == "Transport" {

		transportEvent, err = es.ProcessTransportEvent(ctx, in.TransportEvent)
		if err != nil {
			es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
	}

	err = es.insertEvent(ctx, in.EventType, InsertEventSQL, &event, eventClassifier, InsertEventClassifierSQL, equipmentEvent, InsertEquipmentEventSQL, operationsEvent, InsertOperationsEventSQL, shipmentEvent, InsertShipmentEventSQL, transportEvent, InsertTransportEventSQL, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	eventResponse := tntproto.CreateEventResponse{}
	eventResponse.Event = &event
	return &eventResponse, nil
}

// insertEvent - Insert Event
func (es *EventService) insertEvent(ctx context.Context, eventType string, insertEventSQL string, event *tntproto.Event, eventClassifier *tntproto.EventClassifier, insertEventClassifierSQL string, equipmentEvent *tntproto.EquipmentEvent, insertEquipmentEventSQL string, operationsEvent *tntproto.OperationsEvent, insertOperationsEventSQL string, shipmentEvent *tntproto.ShipmentEvent, insertShipmentEventSQL string, transportEvent *tntproto.TransportEvent, insertTransportEventSQL string, userEmail string, requestID string) error {
	eventTmp, err := es.CrEventStruct(ctx, event, userEmail, requestID)
	if err != nil {
		es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	err = es.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertEventSQL, eventTmp)
		if err != nil {
			es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		event.EventD.Id = uint32(uID)

		eventClassifier.EventId = event.EventD.Id
		_, err = tx.NamedExecContext(ctx, insertEventClassifierSQL, eventClassifier)
		if err != nil {
			es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		if eventType == "Equipment" {
			equipmentEvent.EquipmentEventD.EventId = event.EventD.Id
			_, err = tx.NamedExecContext(ctx, insertEquipmentEventSQL, equipmentEvent)
			if err != nil {
				es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
				return err
			}
		} else if eventType == "Operations" {

			operationsEvent.OperationsEventD.EventId = event.EventD.Id
			_, err = tx.NamedExecContext(ctx, insertOperationsEventSQL, operationsEvent)
			if err != nil {
				es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
				return err
			}
		} else if eventType == "Shipment" {

			shipmentEvent.ShipmentEventD.EventId = event.EventD.Id
			_, err = tx.NamedExecContext(ctx, insertShipmentEventSQL, shipmentEvent)
			if err != nil {
				es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
				return err
			}
		} else if eventType == "Transport" {

			transportEvent.TransportEventD.EventId = event.EventD.Id
			_, err = tx.NamedExecContext(ctx, insertTransportEventSQL, transportEvent)
			if err != nil {
				es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
				return err
			}
		}
		return nil
	})
	if err != nil {
		es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// GetEvents - Get Events
func (es *EventService) GetEvents(ctx context.Context, in *tntproto.GetEventsRequest) (*tntproto.GetEventsResponse, error) {
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

		events := []*tntproto.Event{}

		nselectEventsSQL := selectEventsSQL + query

		rows, err := es.DBService.DB.QueryxContext(ctx, nselectEventsSQL)
		if err != nil {
			es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		for rows.Next() {

			eventTmp := tntstruct.Event{}
			err = rows.StructScan(&eventTmp)
			if err != nil {
				es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
				return nil, err
			}

			getRequest := commonproto.GetRequest{}
			getRequest.UserEmail = in.UserEmail
			getRequest.RequestId = in.RequestId
			event, err := es.getEventStruct(ctx, &getRequest, eventTmp)
			if err != nil {
				es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
				return nil, err
			}
			events = append(events, event)

		}

		eventResponse := tntproto.GetEventsResponse{}
		if len(events) != 0 {
			next := events[len(events)-1].EventD.Id
			next--
			nextc := common.EncodeCursor(next)
			eventResponse = tntproto.GetEventsResponse{Events: events, NextCursor: nextc}
		} else {
			eventResponse = tntproto.GetEventsResponse{Events: events, NextCursor: "0"}
		}
		return &eventResponse, nil
	}
}

// GetEvent - find By ID
func (es *EventService) GetEvent(ctx context.Context, inReq *tntproto.GetEventRequest) (*tntproto.GetEventResponse, error) {
	in := inReq.GetRequest
	select {
	case <-ctx.Done():
		err := errors.New("Client closed connection")
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	default:
		uuid4byte, err := common.UUIDStrToBytes(in.Id)
		if err != nil {
			es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		nselectEventsSQL := selectEventsSQL + ` where uuid4 = ?;`
		row := es.DBService.DB.QueryRowxContext(ctx, nselectEventsSQL, uuid4byte)
		eventTmp := tntstruct.Event{}
		err = row.StructScan(&eventTmp)
		if err != nil {
			es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		event, err := es.getEventStruct(ctx, in, eventTmp)
		if err != nil {
			es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		eventResponse := tntproto.GetEventResponse{}
		eventResponse.Event = event
		return &eventResponse, nil
	}
}

// getEventStruct - Get Event header
func (es *EventService) getEventStruct(ctx context.Context, in *commonproto.GetRequest, eventTmp tntstruct.Event) (*tntproto.Event, error) {
	eventT := new(tntproto.EventT)
	eventT.EventCreatedDateTime = common.TimeToTimestamp(eventTmp.EventT.EventCreatedDateTime)
	eventT.EventDateTime = common.TimeToTimestamp(eventTmp.EventT.EventDateTime)

	event := tntproto.Event{EventD: eventTmp.EventD, EventT: eventT}

	return &event, nil
}

// CrEventStruct - process Event details
func (es *EventService) CrEventStruct(ctx context.Context, event *tntproto.Event, userEmail string, requestID string) (*tntstruct.Event, error) {
	eventT := new(tntstruct.EventT)
	eventT.EventCreatedDateTime = common.TimestampToTime(event.EventT.EventCreatedDateTime)
	eventT.EventDateTime = common.TimestampToTime(event.EventT.EventDateTime)

	eventTmp := tntstruct.Event{EventD: event.EventD, EventT: eventT}
	return &eventTmp, nil
}

// ProcessEventClassifier - Process EventClassifier
func (es *EventService) ProcessEventClassifier(ctx context.Context, in *tntproto.CreateEventClassifierRequest) (*tntproto.EventClassifier, error) {
	eventClassifier := tntproto.EventClassifier{}
	eventClassifier.EventClassifierCode = in.EventClassifierCode
	eventClassifier.EventClassifierName = in.EventClassifierName
	eventClassifier.EventClassifierName = in.EventClassifierDescription
	return &eventClassifier, nil
}
