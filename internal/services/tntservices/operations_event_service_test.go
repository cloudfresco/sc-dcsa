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

func TestEventService_CreateOperationsEvent(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	eventService := NewEventService(log, dbService, redisService, userServiceClient)

	operationsEvent1 := tntproto.OperationsEventRequest{}
	operationsEvent1.EventClassifierCode = "ACT"
	operationsEvent1.Publisher = "be5bc290-7bac-48bb-a211-f3fa5a3ab3ae"
	operationsEvent1.PublisherRole = "TR"
	operationsEvent1.OperationsEventTypeCode = "DEPA"
	operationsEvent1.EventLocation = "06aca2f6-f1d0-48f8-ba46-9a3480adfd23"
	operationsEvent1.TransportCallId = uint32(10)
	operationsEvent1.PortCallServiceTypeCode = "BUNK"
	operationsEvent1.FacilityTypeCode = "BRTH"
	operationsEvent1.DelayReasonCode = "ANA"
	operationsEvent1.VesselPosition = ""
	operationsEvent1.Remark = ""
	operationsEvent1.PortCallPhaseTypeCode = ""
	operationsEvent1.VesselDraft = float64(0)
	operationsEvent1.VesselDraftUnit = ""
	operationsEvent1.MilesRemainingToDestination = float64(0)
	operationsEvent1.EventCreatedDateTime = "11/15/2023"
	operationsEvent1.EventDateTime = "12/11/2023"
	operationsEvent1.UserId = "auth0|66fd06d0bfea78a82bb42459"
	operationsEvent1.UserEmail = "sprov300@gmail.com"
	operationsEvent1.RequestId = "bks1m1g91jau4nkks2f0"

	operationsEvent2 := tntproto.OperationsEventRequest{}
	operationsEvent2.EventClassifierCode = "EST"
	operationsEvent2.Publisher = "be5bc290-7bac-48bb-a211-f3fa5a3ab3ae"
	operationsEvent2.PublisherRole = "CA"
	operationsEvent2.OperationsEventTypeCode = "ARRI"
	operationsEvent2.EventLocation = "6748a259-fb7e-4f27-9a88-3669e8b9c5f8"
	operationsEvent2.TransportCallId = uint32(9)
	operationsEvent2.PortCallServiceTypeCode = "SAFE"
	operationsEvent2.FacilityTypeCode = "BRTH"
	operationsEvent2.DelayReasonCode = "ANA"
	operationsEvent2.VesselPosition = ""
	operationsEvent2.Remark = ""
	operationsEvent2.PortCallPhaseTypeCode = ""
	operationsEvent2.VesselDraft = float64(0)
	operationsEvent2.VesselDraftUnit = ""
	operationsEvent2.MilesRemainingToDestination = float64(0)
	operationsEvent2.EventCreatedDateTime = "09/12/2024"
	operationsEvent2.EventDateTime = "11/11/2024"
	operationsEvent2.UserId = "auth0|66fd06d0bfea78a82bb42459"
	operationsEvent2.UserEmail = "sprov300@gmail.com"
	operationsEvent2.RequestId = "bks1m1g91jau4nkks2f0"

	operationsEvents := []*tntproto.OperationsEventRequest{}
	operationsEvents = append(operationsEvents, &operationsEvent1)
	operationsEvents = append(operationsEvents, &operationsEvent2)

	equipEvents := tntproto.CreateOperationsEventRequest{}
	equipEvents.OperationsEventRequests = operationsEvents
	equipEvents.UserId = "auth0|66fd06d0bfea78a82bb42459"
	equipEvents.UserEmail = "sprov300@gmail.com"
	equipEvents.RequestId = "bks1m1g91jau4nkks2f0"

	type args struct {
		ctx context.Context
		in  *tntproto.CreateOperationsEventRequest
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
		_, err := tt.es.CreateOperationsEvent(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("EventService.CreateOperationsEvent() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
	}
}

func TestEventService_LoadOperationsRelatedEntities(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	eventService := NewEventService(log, dbService, redisService, userServiceClient)

	operationsEvent1, err := GetOperationsEvent(uint32(4), []byte{83, 131, 18, 218, 103, 76, 66, 120, 191, 159, 16, 226, 167, 192, 24, 227}, "PLN", "7bf6f428-58f0-4347-9ce8-d6be2f5d5745", "PLT", "STRT", "06aca2f6-f1d0-48f8-ba46-9a3480adfd23", uint32(10), "PILO", "", "ANA", "", "", "INBD", float64(0), "", float64(0), "2022-05-03T21:02:44Z", "2022-03-07T00:00:00Z")
	if err != nil {
		t.Error(err)
		return
	}

	operationsEvent2, err := GetOperationsEvent(uint32(3), []byte{211, 48, 182, 245, 237, 203, 78, 158, 160, 159, 233, 142, 145, 222, 186, 149}, "REQ", "c49ea2d6-3806-46c8-8490-294affc71286", "TR", "ARRI", "b4454ae5-dcd4-4955-8080-1f986aa5c6c3", uint32(1), "", "BRTH", "", "1d09e9e9-dba3-4de1-8ef8-3ab6d32dbb40", "", "INBD", float64(0), "", float64(0), "2022-05-03T21:02:44Z", "2022-03-07T00:00:00Z")
	if err != nil {
		t.Error(err)
		return
	}

	operationsEvents := []*tntproto.OperationsEvent{}
	operationsEvents = append(operationsEvents, operationsEvent1, operationsEvent2)

	form := tntproto.LoadOperationsRelatedEntitiesRequest{}
	form.Limit = "2"
	form.NextCursor = ""
	form.UserEmail = "sprov300@gmail.com"
	form.RequestId = "bks1m1g91jau4nkks2f0"

	nextc := "Mg=="
	operationsEventResp := tntproto.LoadOperationsRelatedEntitiesResponse{OperationsEvents: operationsEvents, NextCursor: nextc}

	type args struct {
		ctx context.Context
		in  *tntproto.LoadOperationsRelatedEntitiesRequest
	}
	tests := []struct {
		es      *EventService
		args    args
		want    *tntproto.LoadOperationsRelatedEntitiesResponse
		wantErr bool
	}{
		{
			es: eventService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &operationsEventResp,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		operationsEventResponse, err := tt.es.LoadOperationsRelatedEntities(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("EventService.LoadOperationsRelatedEntities() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(operationsEventResponse, tt.want) {
			t.Errorf("EventService.LoadOperationsRelatedEntities() = %v, want %v", operationsEventResponse, tt.want)
		}
		operationsEventResult := operationsEventResponse.OperationsEvents[0]
		assert.NotNil(t, operationsEventResult)
		assert.Equal(t, operationsEventResult.OperationsEventD.EventClassifierCode, "PLN", "they should be equal")
		assert.Equal(t, operationsEventResult.OperationsEventD.PublisherRole, "PLT", "they should be equal")
		assert.Equal(t, operationsEventResult.OperationsEventD.OperationsEventTypeCode, "STRT", "they should be equal")
		assert.Equal(t, operationsEventResult.OperationsEventD.PortCallServiceTypeCode, "PILO", "they should be equal")
		assert.Equal(t, operationsEventResult.OperationsEventD.DelayReasonCode, "ANA", "they should be equal")
		assert.Equal(t, operationsEventResult.OperationsEventD.PortCallPhaseTypeCode, "INBD", "they should be equal")
		assert.NotNil(t, operationsEventResult.OperationsEventT.EventDateTime)
	}
}

func GetOperationsEvent(id uint32, eventIdS []byte, eventClassifierCode string, publisher string, publisherRole string, operationsEventTypeCode string, eventLocation string, transportCallId uint32, portCallServiceTypeCode string, facilityTypeCode string, delayReasonCode string, vesselPosition string, remark string, portCallPhaseTypeCode string, vesselDraft float64, vesselDraftUnit string, milesRemainingToDestination float64, eventCreatedDateTime string, eventDateTime string) (*tntproto.OperationsEvent, error) {
	eventCreatedDateTime1, err := common.ConvertTimeToTimestamp(Layout, eventCreatedDateTime)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	eventDateTime1, err := common.ConvertTimeToTimestamp(Layout, eventDateTime)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	operationsEventD := new(tntproto.OperationsEventD)
	operationsEventD.Id = id
	operationsEventD.EventIdS = eventIdS
	operationsEventD.EventClassifierCode = eventClassifierCode
	operationsEventD.Publisher = publisher
	operationsEventD.PublisherRole = publisherRole
	operationsEventD.OperationsEventTypeCode = operationsEventTypeCode
	operationsEventD.EventLocation = eventLocation
	operationsEventD.TransportCallId = transportCallId
	operationsEventD.PortCallServiceTypeCode = portCallServiceTypeCode
	operationsEventD.FacilityTypeCode = facilityTypeCode
	operationsEventD.DelayReasonCode = delayReasonCode
	operationsEventD.VesselPosition = vesselPosition
	operationsEventD.Remark = remark
	operationsEventD.PortCallPhaseTypeCode = portCallPhaseTypeCode
	operationsEventD.VesselDraft = vesselDraft
	operationsEventD.VesselDraftUnit = vesselDraftUnit
	operationsEventD.MilesRemainingToDestination = milesRemainingToDestination

	operationsEventT := new(tntproto.OperationsEventT)
	operationsEventT.EventCreatedDateTime = eventCreatedDateTime1
	operationsEventT.EventDateTime = eventDateTime1

	operationsEvent := tntproto.OperationsEvent{OperationsEventD: operationsEventD, OperationsEventT: operationsEventT}

	return &operationsEvent, nil
}
