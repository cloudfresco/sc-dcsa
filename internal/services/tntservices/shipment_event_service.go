package tntservices

import (
	"context"
	"errors"
	"time"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	bkgproto "github.com/cloudfresco/sc-dcsa/internal/proto-gen/bkg/v2"
	commonproto "github.com/cloudfresco/sc-dcsa/internal/proto-gen/common/v1"
	eblproto "github.com/cloudfresco/sc-dcsa/internal/proto-gen/ebl/v1"
	tntproto "github.com/cloudfresco/sc-dcsa/internal/proto-gen/tnt/v3"
	bkgservice "github.com/cloudfresco/sc-dcsa/internal/services/bkgservices"
	eblservice "github.com/cloudfresco/sc-dcsa/internal/services/eblservices"
	tntstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/tnt/v3"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// InsertShipmentEventSQL - insert ShipmentEventSQL query
const InsertShipmentEventSQL = `insert into shipment_events
	  ( 
event_id,
event_classifier_code,
shipment_event_type_code,
document_type_code,
document_id,
document_reference,
reason,
event_created_date_time,
event_date_time
  )
  values (:event_id,
:event_classifier_code,
:shipment_event_type_code,
:document_type_code,
:document_id,
:document_reference,
:reason,
:event_created_date_time,
:event_date_time);`

// selectShipmentEventsSQL - select ShipmentEventsSQL Query
const selectShipmentEventsSQL = `select 
  id,
  event_id,
  event_classifier_code,
  shipment_event_type_code,
  document_type_code,
  document_id,
  document_reference,
  reason,
  event_created_date_time,
  event_date_time from shipment_events`

// InsertShipmentEventTypeSQL - insert ShipmentEventTypeSQL query
const InsertShipmentEventTypeSQL = `insert into shipment_event_types
	  ( 
shipment_event_type_code,
shipment_event_type_name,
shipment_event_type_description,
shipment_event_id
  )
  values (:shipment_event_type_code,
:shipment_event_type_name,
:shipment_event_type_description,
:shipment_event_id);`

// CreateShipmentEvent - Create ShipmentEvent
func (es *EventService) CreateShipmentEvent(ctx context.Context, in *tntproto.CreateShipmentEventRequest) (*tntproto.CreateShipmentEventResponse, error) {
	shipmentEvent, err := es.ProcessShipmentEvent(ctx, in)
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = es.insertShipmentEvent(ctx, InsertShipmentEventSQL, shipmentEvent, InsertShipmentEventTypeSQL, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	shipmentEventResponse := tntproto.CreateShipmentEventResponse{}
	shipmentEventResponse.ShipmentEvent = shipmentEvent
	return &shipmentEventResponse, nil
}

// ProcessShipmentEvent - Process ShipmentEvent
func (es *EventService) ProcessShipmentEvent(ctx context.Context, in *tntproto.CreateShipmentEventRequest) (*tntproto.ShipmentEvent, error) {
	bkgServ := &bkgservice.BkgService{DBService: es.DBService, RedisService: es.RedisService, UserServiceClient: es.UserServiceClient}
	req := bkgproto.GetBookingByPkRequest{}
	getByIdRequest := commonproto.GetByIdRequest{}
	getByIdRequest.Id = in.BookingId
	getByIdRequest.UserEmail = in.UserEmail
	getByIdRequest.RequestId = in.RequestId
	req.GetByIdRequest = &getByIdRequest
	bookingResponse, err := bkgServ.GetBookingByPk(ctx, &req)
	booking := bookingResponse.Booking
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
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

	shipmentEventD := tntproto.ShipmentEventD{}
	shipmentEventD.EventIdS, err = common.GetUUIDBytes()
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	shipmentEventD.EventClassifierCode = in.EventClassifierCode
	shipmentEventD.ShipmentEventTypeCode = booking.BookingD.DocumentStatus
	shipmentEventD.DocumentTypeCode = in.DocumentTypeCode
	shipmentEventD.DocumentId = booking.BookingD.Id
	shipmentEventD.DocumentReference = booking.BookingD.CarrierBookingRequestReference
	shipmentEventD.Reason = in.Reason

	shipmentEventT := tntproto.ShipmentEventT{}
	shipmentEventT.EventCreatedDateTime = common.TimeToTimestamp(eventCreatedDateTime.UTC().Truncate(time.Second))
	shipmentEventT.EventDateTime = common.TimeToTimestamp(eventDateTime.UTC().Truncate(time.Second))

	shipmentEvent := tntproto.ShipmentEvent{ShipmentEventD: &shipmentEventD, ShipmentEventT: &shipmentEventT}

	inShipmentEventType := in.ShipmentEventType
	shipmentEventType := tntproto.ShipmentEventType{}
	shipmentEventType.ShipmentEventTypeCode = inShipmentEventType.ShipmentEventTypeCode
	shipmentEventType.ShipmentEventTypeName = inShipmentEventType.ShipmentEventTypeName
	shipmentEventType.ShipmentEventTypeDescription = inShipmentEventType.ShipmentEventTypeDescription
	shipmentEvent.ShipmentEventType = &shipmentEventType

	return &shipmentEvent, nil
}

// insertShipmentEvent - Insert ShipmentEvent
func (es *EventService) insertShipmentEvent(ctx context.Context, insertShipmentEventSQL string, shipmentEvent *tntproto.ShipmentEvent, insertShipmentEventTypeSQL string, userEmail string, requestID string) error {
	select {
	case <-ctx.Done():
		err := errors.New("Client closed connection")
		es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	default:
		shipmentEventTmp, err := es.CrShipmentEventStruct(ctx, shipmentEvent, userEmail, requestID)
		if err != nil {
			es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		err = es.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
			res, err := tx.NamedExecContext(ctx, insertShipmentEventSQL, shipmentEventTmp)
			if err != nil {
				es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
				return err
			}

			uID, err := res.LastInsertId()
			if err != nil {
				es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
				return err
			}
			shipmentEvent.ShipmentEventD.Id = uint32(uID)

			shipmentEvent.ShipmentEventType.ShipmentEventId = uint32(uID)

			_, err = tx.NamedExecContext(ctx, insertShipmentEventTypeSQL, shipmentEvent.ShipmentEventType)

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

// CrShipmentEventStruct - process ShipmentEvent details
func (es *EventService) CrShipmentEventStruct(ctx context.Context, shipmentEvent *tntproto.ShipmentEvent, userEmail string, requestID string) (*tntstruct.ShipmentEvent, error) {
	select {
	case <-ctx.Done():
		err := errors.New("Client closed connection")
		es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return nil, err
	default:
		shipmentEventT := new(tntstruct.ShipmentEventT)
		shipmentEventT.EventCreatedDateTime = common.TimestampToTime(shipmentEvent.ShipmentEventT.EventCreatedDateTime)
		shipmentEventT.EventDateTime = common.TimestampToTime(shipmentEvent.ShipmentEventT.EventDateTime)

		shipmentEventTmp := tntstruct.ShipmentEvent{ShipmentEventD: shipmentEvent.ShipmentEventD, ShipmentEventT: shipmentEventT}

		return &shipmentEventTmp, nil
	}
}

// LoadShipmentRelatedEntities - Get ShipmentEvents
func (es *EventService) LoadShipmentRelatedEntities(ctx context.Context, in *tntproto.LoadShipmentRelatedEntitiesRequest) (*tntproto.LoadShipmentRelatedEntitiesResponse, error) {
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

		shipmentEvents := []*tntproto.ShipmentEvent{}

		nselectShipmentEventsSQL := selectShipmentEventsSQL + query

		rows, err := es.DBService.DB.QueryxContext(ctx, nselectShipmentEventsSQL)
		if err != nil {
			es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		for rows.Next() {

			shipmentEventTmp := tntstruct.ShipmentEvent{}
			err = rows.StructScan(&shipmentEventTmp)
			if err != nil {
				es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
				return nil, err
			}

			getRequest := commonproto.GetRequest{}
			getRequest.UserEmail = in.UserEmail
			getRequest.RequestId = in.RequestId
			shipmentEvent, err := es.getShipmentEventStruct(ctx, &getRequest, shipmentEventTmp)
			if err != nil {
				es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
				return nil, err
			}
			shipmentEvents = append(shipmentEvents, shipmentEvent)

		}

		shipmentEventResponse := tntproto.LoadShipmentRelatedEntitiesResponse{}
		if len(shipmentEvents) != 0 {
			next := shipmentEvents[len(shipmentEvents)-1].ShipmentEventD.Id
			next--
			nextc := common.EncodeCursor(next)
			shipmentEventResponse = tntproto.LoadShipmentRelatedEntitiesResponse{ShipmentEvents: shipmentEvents, NextCursor: nextc}
		} else {
			shipmentEventResponse = tntproto.LoadShipmentRelatedEntitiesResponse{ShipmentEvents: shipmentEvents, NextCursor: "0"}
		}
		return &shipmentEventResponse, nil
	}
}

// getShipmentEventStruct - Get ShipmentEvent header
func (es *EventService) getShipmentEventStruct(ctx context.Context, in *commonproto.GetRequest, shipmentEventTmp tntstruct.ShipmentEvent) (*tntproto.ShipmentEvent, error) {
	shipmentEventT := new(tntproto.ShipmentEventT)
	shipmentEventT.EventCreatedDateTime = common.TimeToTimestamp(shipmentEventTmp.ShipmentEventT.EventCreatedDateTime)
	shipmentEventT.EventDateTime = common.TimeToTimestamp(shipmentEventTmp.ShipmentEventT.EventDateTime)

	shipmentEvent := tntproto.ShipmentEvent{ShipmentEventD: shipmentEventTmp.ShipmentEventD, ShipmentEventT: shipmentEventT}

	return &shipmentEvent, nil
}

// CreateShipmentEventFromShippingInstruction - CreateShipmentEventFromShippingInstruction
func (es *EventService) CreateShipmentEventFromShippingInstruction(ctx context.Context, inReq *tntproto.CreateShipmentEventFromShippingInstructionRequest) (*tntproto.CreateShipmentEventFromShippingInstructionResponse, error) {
	in := inReq.CreateShipmentEventRequest
	select {
	case <-ctx.Done():
		err := errors.New("Client closed connection")
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	default:

		eblServ := &eblservice.ShippingInstructionService{DBService: es.DBService, RedisService: es.RedisService, UserServiceClient: es.UserServiceClient}
		req := eblproto.GetShippingInstructionByPkRequest{}
		getByIdRequest := commonproto.GetByIdRequest{}
		getByIdRequest.Id = in.ShippingInstructionId
		getByIdRequest.UserEmail = in.UserEmail
		getByIdRequest.RequestId = in.RequestId
		req.GetByIdRequest = &getByIdRequest

		shippingInstructionResp, err := eblServ.GetShippingInstructionByPk(ctx, &req)
		shippingInstruction := shippingInstructionResp.ShippingInstruction
		if err != nil {
			es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
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

		shipmentEventD := tntproto.ShipmentEventD{}
		shipmentEventD.EventIdS, err = common.GetUUIDBytes()
		if err != nil {
			es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		shipmentEventD.EventClassifierCode = in.EventClassifierCode
		shipmentEventD.ShipmentEventTypeCode = shippingInstruction.ShippingInstructionD.DocumentStatus
		shipmentEventD.DocumentTypeCode = in.DocumentTypeCode
		shipmentEventD.DocumentId = shippingInstruction.ShippingInstructionD.Id
		shipmentEventD.DocumentReference = in.DocumentReference
		shipmentEventD.Reason = in.Reason

		shipmentEventT := tntproto.ShipmentEventT{}
		shipmentEventT.EventCreatedDateTime = common.TimeToTimestamp(eventCreatedDateTime.UTC().Truncate(time.Second))
		shipmentEventT.EventDateTime = common.TimeToTimestamp(eventDateTime.UTC().Truncate(time.Second))

		shipmentEvent := tntproto.ShipmentEvent{ShipmentEventD: &shipmentEventD, ShipmentEventT: &shipmentEventT}

		err = es.insertShipmentEvent(ctx, InsertShipmentEventSQL, &shipmentEvent, InsertShipmentEventTypeSQL, in.GetUserEmail(), in.GetRequestId())
		if err != nil {
			es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		shipmentEventResponse := tntproto.CreateShipmentEventFromShippingInstructionResponse{}
		shipmentEventResponse.ShipmentEvent = &shipmentEvent
		return &shipmentEventResponse, nil
	}
}
