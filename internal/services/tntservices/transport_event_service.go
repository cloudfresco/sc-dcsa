package tntservices

import (
	"context"
	"errors"
	"time"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	commonproto "github.com/cloudfresco/sc-dcsa/internal/proto-gen/common/v1"
	tntproto "github.com/cloudfresco/sc-dcsa/internal/proto-gen/tnt/v3"
	tntstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/tnt/v3"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// InsertTransportEventSQL - insert TransportEventSQL query
const InsertTransportEventSQL = `insert into transport_events
	  ( 
event_id,
event_classifier_code,
transport_event_type_code,
delay_reason_code,
change_remark,
transport_call_id,
event_created_date_time,
event_date_time
  )
  values (:event_id,
:event_classifier_code,
:transport_event_type_code,
:delay_reason_code,
:change_remark,
:transport_call_id,
:event_created_date_time,
:event_date_time);`

// selectTransportEventsSQL - select TransportEventsSQL Query
const selectTransportEventsSQL = `select 
  id,
  event_id,
  event_classifier_code,
  transport_event_type_code,
  delay_reason_code,
  change_remark,
  transport_call_id,
  event_created_date_time,
  event_date_time from transport_events`

// InsertTransportEventTypeSQL - insert TransportEventTypeSQL query
const InsertTransportEventTypeSQL = `insert into transport_event_types
	  ( 
transport_event_type_code,
transport_event_type_name,
transport_event_type_description,
transport_event_id
  )
  values (:transport_event_type_code,
:transport_event_type_name,
:transport_event_type_description,
:transport_event_id);`

// CreateTransportEvent - Create TransportEvent
func (es *EventService) CreateTransportEvent(ctx context.Context, in *tntproto.CreateTransportEventRequest) (*tntproto.CreateTransportEventResponse, error) {
	transportEvent, err := es.ProcessTransportEvent(ctx, in)
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = es.insertTransportEvent(ctx, InsertTransportEventSQL, transportEvent, InsertTransportEventTypeSQL, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	transportEventResponse := tntproto.CreateTransportEventResponse{}
	transportEventResponse.TransportEvent = transportEvent

	return &transportEventResponse, nil
}

// ProcessTransportEvent - Process TransportEvent
func (es *EventService) ProcessTransportEvent(ctx context.Context, in *tntproto.CreateTransportEventRequest) (*tntproto.TransportEvent, error) {
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

	transportEventD := tntproto.TransportEventD{}
	transportEventD.EventIdS, err = common.GetUUIDBytes()
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	transportEventD.EventClassifierCode = in.EventClassifierCode
	transportEventD.TransportEventTypeCode = in.TransportEventTypeCode
	transportEventD.DelayReasonCode = in.DelayReasonCode
	transportEventD.ChangeRemark = in.ChangeRemark
	transportEventD.TransportCallId = in.TransportCallId

	transportEventT := tntproto.TransportEventT{}
	transportEventT.EventCreatedDateTime = common.TimeToTimestamp(eventCreatedDateTime.UTC().Truncate(time.Second))
	transportEventT.EventDateTime = common.TimeToTimestamp(eventDateTime.UTC().Truncate(time.Second))

	transportEvent := tntproto.TransportEvent{TransportEventD: &transportEventD, TransportEventT: &transportEventT}

	inTransportEventType := in.TransportEventType
	transportEventType := tntproto.TransportEventType{}
	transportEventType.TransportEventTypeCode = inTransportEventType.TransportEventTypeCode
	transportEventType.TransportEventTypeName = inTransportEventType.TransportEventTypeName
	transportEventType.TransportEventTypeDescription = inTransportEventType.TransportEventTypeDescription
	transportEvent.TransportEventType = &transportEventType

	return &transportEvent, nil
}

// insertTransportEvent - Insert TransportEvent
func (es *EventService) insertTransportEvent(ctx context.Context, insertTransportEventSQL string, transportEvent *tntproto.TransportEvent, insertTransportEventTypeSQL string, userEmail string, requestID string) error {
	select {
	case <-ctx.Done():
		err := errors.New("Client closed connection")
		es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	default:
		transportEventTmp, err := es.CrTransportEventStruct(ctx, transportEvent, userEmail, requestID)
		if err != nil {
			es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		err = es.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
			res, err := tx.NamedExecContext(ctx, insertTransportEventSQL, transportEventTmp)
			if err != nil {
				es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
				return err
			}

			uID, err := res.LastInsertId()
			if err != nil {
				es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
				return err
			}
			transportEvent.TransportEventD.Id = uint32(uID)

			transportEvent.TransportEventType.TransportEventId = uint32(uID)

			_, err = tx.NamedExecContext(ctx, insertTransportEventTypeSQL, transportEvent.TransportEventType)

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
}

// CrTransportEventStruct - process TransportEvent details
func (es *EventService) CrTransportEventStruct(ctx context.Context, transportEvent *tntproto.TransportEvent, userEmail string, requestID string) (*tntstruct.TransportEvent, error) {
	select {
	case <-ctx.Done():
		err := errors.New("Client closed connection")
		es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return nil, err
	default:
		transportEventT := new(tntstruct.TransportEventT)
		transportEventT.EventCreatedDateTime = common.TimestampToTime(transportEvent.TransportEventT.EventCreatedDateTime)
		transportEventT.EventDateTime = common.TimestampToTime(transportEvent.TransportEventT.EventDateTime)

		transportEventTmp := tntstruct.TransportEvent{TransportEventD: transportEvent.TransportEventD, TransportEventT: transportEventT}

		return &transportEventTmp, nil
	}
}

// LoadTransportRelatedEntities - Get TransportEvents
func (es *EventService) LoadTransportRelatedEntities(ctx context.Context, in *tntproto.LoadTransportRelatedEntitiesRequest) (*tntproto.LoadTransportRelatedEntitiesResponse, error) {
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

		transportEvents := []*tntproto.TransportEvent{}

		nselectTransportEventsSQL := selectTransportEventsSQL + query

		rows, err := es.DBService.DB.QueryxContext(ctx, nselectTransportEventsSQL)
		if err != nil {
			es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		for rows.Next() {

			transportEventTmp := tntstruct.TransportEvent{}
			err = rows.StructScan(&transportEventTmp)
			if err != nil {
				es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
				return nil, err
			}

			getRequest := commonproto.GetRequest{}
			getRequest.UserEmail = in.UserEmail
			getRequest.RequestId = in.RequestId
			transportEvent, err := es.getTransportEventStruct(ctx, &getRequest, transportEventTmp)
			if err != nil {
				es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
				return nil, err
			}
			transportEvents = append(transportEvents, transportEvent)

		}

		transportEventResponse := tntproto.LoadTransportRelatedEntitiesResponse{}
		if len(transportEvents) != 0 {
			next := transportEvents[len(transportEvents)-1].TransportEventD.Id
			next--
			nextc := common.EncodeCursor(next)
			transportEventResponse = tntproto.LoadTransportRelatedEntitiesResponse{TransportEvents: transportEvents, NextCursor: nextc}
		} else {
			transportEventResponse = tntproto.LoadTransportRelatedEntitiesResponse{TransportEvents: transportEvents, NextCursor: "0"}
		}
		return &transportEventResponse, nil
	}
}

// getTransportEventStruct - Get TransportEvent header
func (es *EventService) getTransportEventStruct(ctx context.Context, in *commonproto.GetRequest, transportEventTmp tntstruct.TransportEvent) (*tntproto.TransportEvent, error) {
	transportEventT := new(tntproto.TransportEventT)
	transportEventT.EventCreatedDateTime = common.TimeToTimestamp(transportEventTmp.TransportEventT.EventCreatedDateTime)
	transportEventT.EventDateTime = common.TimeToTimestamp(transportEventTmp.TransportEventT.EventDateTime)

	transportEvent := tntproto.TransportEvent{TransportEventD: transportEventTmp.TransportEventD, TransportEventT: transportEventT}
	return &transportEvent, nil
}