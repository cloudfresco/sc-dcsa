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

// InsertOperationsEventSQL - insert OperationsEventSQL query
const InsertOperationsEventSQL = `insert into operations_events
	  ( 
event_id,
event_classifier_code,
publisher,
publisher_role,
operations_event_type_code,
event_location,
transport_call_id,
port_call_service_type_code,
facility_type_code,
delay_reason_code,
vessel_position,
remark,
port_call_phase_type_code,
vessel_draft,
vessel_draft_unit,
miles_remaining_to_destination,
event_created_date_time,
event_date_time
  )
  values (:event_id,
:event_classifier_code,
:publisher,
:publisher_role,
:operations_event_type_code,
:event_location,
:transport_call_id,
:port_call_service_type_code,
:facility_type_code,
:delay_reason_code,
:vessel_position,
:remark,
:port_call_phase_type_code,
:vessel_draft,
:vessel_draft_unit,
:miles_remaining_to_destination,
:event_created_date_time,
:event_date_time);`

// selectOperationsEventsSQL - select OperationsEventsSQL Query
const selectOperationsEventsSQL = `select 
  id,
  event_id,
  event_classifier_code,
  publisher,
  publisher_role,
  operations_event_type_code,
  event_location,
  transport_call_id,
  port_call_service_type_code,
  facility_type_code,
  delay_reason_code,
  vessel_position,
  remark,
  port_call_phase_type_code,
  vessel_draft,
  vessel_draft_unit,
  miles_remaining_to_destination,
  event_created_date_time,
  event_date_time from operations_events`

// InsertOperationsEventTypeSQL - insert OperationsEventTypeSQL query
const InsertOperationsEventTypeSQL = `insert into operations_event_types
	  ( 
operations_event_type_code,
operations_event_type_name,
operations_event_type_description,
operations_event_id
  )
  values (:operations_event_type_code,
:operations_event_type_name,
:operations_event_type_description,
:operations_event_id);`

// CreateOperationsEvent - Create OperationsEvent
func (es *EventService) CreateOperationsEvent(ctx context.Context, in *tntproto.CreateOperationsEventRequest) (*tntproto.CreateOperationsEventResponse, error) {
	operationsEvent, err := es.ProcessOperationsEvent(ctx, in)
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = es.insertOperationsEvent(ctx, InsertOperationsEventSQL, operationsEvent, InsertOperationsEventTypeSQL, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	operationsEventResponse := tntproto.CreateOperationsEventResponse{}
	operationsEventResponse.OperationsEvent = operationsEvent
	return &operationsEventResponse, nil
}

// ProcessOperationsEvent - Process OperationsEvent
func (es *EventService) ProcessOperationsEvent(ctx context.Context, in *tntproto.CreateOperationsEventRequest) (*tntproto.OperationsEvent, error) {
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

	operationsEventD := tntproto.OperationsEventD{}
	operationsEventD.EventIdS, err = common.GetUUIDBytes()
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	operationsEventD.EventClassifierCode = in.EventClassifierCode
	operationsEventD.Publisher = in.Publisher
	operationsEventD.PublisherRole = in.PublisherRole
	operationsEventD.OperationsEventTypeCode = in.OperationsEventTypeCode
	operationsEventD.EventLocation = in.EventLocation
	operationsEventD.TransportCallId = in.TransportCallId
	operationsEventD.PortCallServiceTypeCode = in.PortCallServiceTypeCode
	operationsEventD.FacilityTypeCode = in.FacilityTypeCode
	operationsEventD.DelayReasonCode = in.DelayReasonCode
	operationsEventD.VesselPosition = in.VesselPosition
	operationsEventD.Remark = in.Remark
	operationsEventD.PortCallPhaseTypeCode = in.PortCallPhaseTypeCode
	operationsEventD.VesselDraft = in.VesselDraft
	operationsEventD.VesselDraftUnit = in.VesselDraftUnit
	operationsEventD.MilesRemainingToDestination = in.MilesRemainingToDestination

	operationsEventT := tntproto.OperationsEventT{}
	operationsEventT.EventCreatedDateTime = common.TimeToTimestamp(eventCreatedDateTime.UTC().Truncate(time.Second))
	operationsEventT.EventDateTime = common.TimeToTimestamp(eventDateTime.UTC().Truncate(time.Second))

	operationsEvent := tntproto.OperationsEvent{OperationsEventD: &operationsEventD, OperationsEventT: &operationsEventT}

	inOperationsEventType := in.OperationsEventType
	operationsEventType := tntproto.OperationsEventType{}
	operationsEventType.OperationsEventTypeCode = inOperationsEventType.OperationsEventTypeCode
	operationsEventType.OperationsEventTypeName = inOperationsEventType.OperationsEventTypeName
	operationsEventType.OperationsEventTypeDescription = inOperationsEventType.OperationsEventTypeDescription
	operationsEvent.OperationsEventType = &operationsEventType

	return &operationsEvent, nil
}

// insertOperationsEvent - Insert OperationsEvent
func (es *EventService) insertOperationsEvent(ctx context.Context, insertOperationsEventSQL string, operationsEvent *tntproto.OperationsEvent, insertOperationsEventTypeSQL string, userEmail string, requestID string) error {
	operationsEventTmp, err := es.CrOperationsEventStruct(ctx, operationsEvent, userEmail, requestID)
	if err != nil {
		es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	err = es.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertOperationsEventSQL, operationsEventTmp)
		if err != nil {
			es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		operationsEvent.OperationsEventD.Id = uint32(uID)

		operationsEvent.OperationsEventType.OperationsEventId = uint32(uID)
		_, err = tx.NamedExecContext(ctx, insertOperationsEventTypeSQL, operationsEvent.OperationsEventType)

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

// CrOperationsEventStruct - process OperationsEvent details
func (es *EventService) CrOperationsEventStruct(ctx context.Context, operationsEvent *tntproto.OperationsEvent, userEmail string, requestID string) (*tntstruct.OperationsEvent, error) {
	operationsEventT := new(tntstruct.OperationsEventT)
	operationsEventT.EventCreatedDateTime = common.TimestampToTime(operationsEvent.OperationsEventT.EventCreatedDateTime)
	operationsEventT.EventDateTime = common.TimestampToTime(operationsEvent.OperationsEventT.EventDateTime)

	operationsEventTmp := tntstruct.OperationsEvent{OperationsEventD: operationsEvent.OperationsEventD, OperationsEventT: operationsEventT}

	return &operationsEventTmp, nil
}

// LoadOperationsRelatedEntities - Get OperationsEvents
func (es *EventService) LoadOperationsRelatedEntities(ctx context.Context, in *tntproto.LoadOperationsRelatedEntitiesRequest) (*tntproto.LoadOperationsRelatedEntitiesResponse, error) {
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

		operationsEvents := []*tntproto.OperationsEvent{}

		nselectOperationsEventsSQL := selectOperationsEventsSQL + query

		rows, err := es.DBService.DB.QueryxContext(ctx, nselectOperationsEventsSQL)
		if err != nil {
			es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		for rows.Next() {

			operationsEventTmp := tntstruct.OperationsEvent{}
			err = rows.StructScan(&operationsEventTmp)
			if err != nil {
				es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
				return nil, err
			}

			getRequest := commonproto.GetRequest{}
			getRequest.UserEmail = in.UserEmail
			getRequest.RequestId = in.RequestId
			operationsEvent, err := es.getOperationsEventStruct(ctx, &getRequest, operationsEventTmp)
			if err != nil {
				es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
				return nil, err
			}
			operationsEvents = append(operationsEvents, operationsEvent)

		}

		operationsEventsResponse := tntproto.LoadOperationsRelatedEntitiesResponse{}
		if len(operationsEvents) != 0 {
			next := operationsEvents[len(operationsEvents)-1].OperationsEventD.Id
			next--
			nextc := common.EncodeCursor(next)
			operationsEventsResponse = tntproto.LoadOperationsRelatedEntitiesResponse{OperationsEvents: operationsEvents, NextCursor: nextc}
		} else {
			operationsEventsResponse = tntproto.LoadOperationsRelatedEntitiesResponse{OperationsEvents: operationsEvents, NextCursor: "0"}
		}
		return &operationsEventsResponse, nil
	}
}

// getOperationsEventStruct - Get OperationsEvent header
func (es *EventService) getOperationsEventStruct(ctx context.Context, in *commonproto.GetRequest, operationsEventTmp tntstruct.OperationsEvent) (*tntproto.OperationsEvent, error) {
	operationsEventT := new(tntproto.OperationsEventT)
	operationsEventT.EventCreatedDateTime = common.TimeToTimestamp(operationsEventTmp.OperationsEventT.EventCreatedDateTime)
	operationsEventT.EventDateTime = common.TimeToTimestamp(operationsEventTmp.OperationsEventT.EventDateTime)

	operationsEvent := tntproto.OperationsEvent{OperationsEventD: operationsEventTmp.OperationsEventD, OperationsEventT: operationsEventT}

	return &operationsEvent, nil
}