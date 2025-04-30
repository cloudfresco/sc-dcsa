package tntservices

import (
	"context"
	"reflect"
	"testing"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	tntproto "github.com/cloudfresco/sc-dcsa/internal/protogen/tnt/v3"
	"github.com/cloudfresco/sc-dcsa/test"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestEventService_CreateEquipmentEvent(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	eventService := NewEventService(log, dbService, redisService, userServiceClient)

	equipmentEvent1 := tntproto.EquipmentEventRequest{}
	equipmentEvent1.EventClassifierCode = "ACT"
	equipmentEvent1.EquipmentEventTypeCode = "LOAD"
	equipmentEvent1.EquipmentReference = "equipref3453"
	equipmentEvent1.EmptyIndicatorCode = "EMPTY"
	equipmentEvent1.TransportCallId = uint32(6)
	equipmentEvent1.EventLocation = ""
	equipmentEvent1.EventCreatedDateTime = "11/15/2023"
	equipmentEvent1.EventDateTime = "12/11/2023"
	equipmentEvent1.UserId = "auth0|66fd06d0bfea78a82bb42459"
	equipmentEvent1.UserEmail = "sprov300@gmail.com"
	equipmentEvent1.RequestId = "bks1m1g91jau4nkks2f0"

	equipmentEvent2 := tntproto.EquipmentEventRequest{}
	equipmentEvent2.EventClassifierCode = "EST"
	equipmentEvent2.EquipmentEventTypeCode = "LOAD"
	equipmentEvent2.EquipmentReference = "APZU4812090"
	equipmentEvent2.EmptyIndicatorCode = "EMPTY"
	equipmentEvent2.TransportCallId = uint32(6)
	equipmentEvent2.EventLocation = ""
	equipmentEvent2.EventCreatedDateTime = "09/12/2024"
	equipmentEvent2.EventDateTime = "11/11/2024"
	equipmentEvent2.UserId = "auth0|66fd06d0bfea78a82bb42459"
	equipmentEvent2.UserEmail = "sprov300@gmail.com"
	equipmentEvent2.RequestId = "bks1m1g91jau4nkks2f0"

	equipmentEvents := []*tntproto.EquipmentEventRequest{}
	equipmentEvents = append(equipmentEvents, &equipmentEvent1)
	equipmentEvents = append(equipmentEvents, &equipmentEvent2)

	equipEvents := tntproto.CreateEquipmentEventRequest{}
	equipEvents.EquipmentEventRequests = equipmentEvents
	equipEvents.UserId = "auth0|66fd06d0bfea78a82bb42459"
	equipEvents.UserEmail = "sprov300@gmail.com"
	equipEvents.RequestId = "bks1m1g91jau4nkks2f0"

	type args struct {
		ctx context.Context
		in  *tntproto.CreateEquipmentEventRequest
	}
	tests := []struct {
		es      *EventService
		args    args
		wantErr bool
	}{
		{
			es: eventService,
			args: args{
				ctx: ctx,
				in:  &equipEvents,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		_, err := tt.es.CreateEquipmentEvent(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("EventService.CreateEquipmentEvent() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
	}
}

func TestEventService_LoadEquipmentRelatedEntities(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	eventService := NewEventService(log, dbService, redisService, userServiceClient)

	equipmentEvent1, err := GetEquipmentEvent(uint32(2), []byte{132, 219, 146, 61, 42, 25, 78, 176, 190, 181, 68, 108, 30, 197, 125, 52}, "EST", "LOAD", "APZU4812090", "EMPTY", uint32(7), "", "2021-01-09T14:12:56Z", "2019-11-12T07:41:00Z")
	if err != nil {
		t.Error(err)
		return
	}

	equipmentEvent2, err := GetEquipmentEvent(uint32(1), []byte{94, 81, 231, 44, 216, 114, 17, 234, 129, 28, 15, 143, 16, 163, 46, 162}, "ACT", "LOAD", "equipref3453", "EMPTY", uint32(7), "", "2003-05-02T21:02:44Z", "2003-05-03T21:02:44Z")
	if err != nil {
		t.Error(err)
		return
	}

	equipmentEvents := []*tntproto.EquipmentEvent{}
	equipmentEvents = append(equipmentEvents, equipmentEvent1, equipmentEvent2)

	form := tntproto.LoadEquipmentRelatedEntitiesRequest{}
	form.Limit = "2"
	form.NextCursor = ""
	form.UserEmail = "sprov300@gmail.com"
	form.RequestId = "bks1m1g91jau4nkks2f0"

	nextc := "MA=="
	equipmentEventResp := tntproto.LoadEquipmentRelatedEntitiesResponse{EquipmentEvents: equipmentEvents, NextCursor: nextc}

	type args struct {
		ctx context.Context
		in  *tntproto.LoadEquipmentRelatedEntitiesRequest
	}
	tests := []struct {
		es      *EventService
		args    args
		want    *tntproto.LoadEquipmentRelatedEntitiesResponse
		wantErr bool
	}{
		{
			es: eventService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &equipmentEventResp,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		equipmentEventResponse, err := tt.es.LoadEquipmentRelatedEntities(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("EventService.LoadEquipmentRelatedEntities() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(equipmentEventResponse, tt.want) {
			t.Errorf("EventService.LoadEquipmentRelatedEntities() = %v, want %v", equipmentEventResponse, tt.want)
		}
		equipmentEventResult := equipmentEventResponse.EquipmentEvents[0]
		assert.NotNil(t, equipmentEventResult)
		assert.Equal(t, equipmentEventResult.EquipmentEventD.EventClassifierCode, "EST", "they should be equal")
		assert.Equal(t, equipmentEventResult.EquipmentEventD.EquipmentEventTypeCode, "LOAD", "they should be equal")
		assert.Equal(t, equipmentEventResult.EquipmentEventD.EquipmentReference, "APZU4812090", "they should be equal")
		assert.Equal(t, equipmentEventResult.EquipmentEventD.EmptyIndicatorCode, "EMPTY", "they should be equal")
		assert.NotNil(t, equipmentEventResult.EquipmentEventT.EventCreatedDateTime)
	}
}

func GetEquipmentEvent(id uint32, eventIdS []byte, eventClassifierCode string, equipmentEventTypeCode string, equipmentReference string, emptyIndicatorCode string, transportCallId uint32, eventLocation string, eventCreatedDateTime string, eventDateTime string) (*tntproto.EquipmentEvent, error) {
	eventCreatedDateTime1, err := common.ConvertTimeToTimestamp(Layout, eventCreatedDateTime)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	eventDateTime1, err := common.ConvertTimeToTimestamp(Layout, eventDateTime)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	equipmentEventD := new(tntproto.EquipmentEventD)
	equipmentEventD.Id = id
	equipmentEventD.EventIdS = eventIdS
	equipmentEventD.EventClassifierCode = eventClassifierCode
	equipmentEventD.EquipmentEventTypeCode = equipmentEventTypeCode
	equipmentEventD.EquipmentReference = equipmentReference
	equipmentEventD.EmptyIndicatorCode = emptyIndicatorCode
	equipmentEventD.TransportCallId = transportCallId
	equipmentEventD.EventLocation = eventLocation

	equipmentEventT := new(tntproto.EquipmentEventT)
	equipmentEventT.EventCreatedDateTime = eventCreatedDateTime1
	equipmentEventT.EventDateTime = eventDateTime1

	equipmentEvent := tntproto.EquipmentEvent{EquipmentEventD: equipmentEventD, EquipmentEventT: equipmentEventT}

	return &equipmentEvent, nil
}
