package tntservices

import (
	"context"
	"testing"

	tntproto "github.com/cloudfresco/sc-dcsa/internal/protogen/tnt/v3"
	"github.com/cloudfresco/sc-dcsa/test"
)

func TestEventService_CreateShipmentEvent(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	eventService := NewEventService(log, dbService, redisService, userServiceClient)

	shipmentEvent1 := tntproto.ShipmentEventRequest{}
	shipmentEvent1.EventClassifierCode = "ACT"
	shipmentEvent1.ShipmentEventTypeCode = "APPR"
	shipmentEvent1.DocumentTypeCode = "BKG"
	shipmentEvent1.DocumentId = uint32(2)
	shipmentEvent1.DocumentReference = "BR1239719971"
	shipmentEvent1.Reason = ""
	shipmentEvent1.EventCreatedDateTime = "11/15/2023"
	shipmentEvent1.EventDateTime = "12/11/2023"
	shipmentEvent1.UserId = "auth0|66fd06d0bfea78a82bb42459"
	shipmentEvent1.UserEmail = "sprov300@gmail.com"
	shipmentEvent1.RequestId = "bks1m1g91jau4nkks2f0"

	shipmentEvent2 := tntproto.ShipmentEventRequest{}
	shipmentEvent2.EventClassifierCode = "ACT"
	shipmentEvent2.ShipmentEventTypeCode = "CONF"
	shipmentEvent2.DocumentTypeCode = "BKG"
	shipmentEvent2.DocumentId = uint32(2)
	shipmentEvent2.DocumentReference = "ABC123123123"
	shipmentEvent2.Reason = ""
	shipmentEvent2.EventCreatedDateTime = "09/12/2024"
	shipmentEvent2.EventDateTime = "11/11/2024"
	shipmentEvent2.UserId = "auth0|66fd06d0bfea78a82bb42459"
	shipmentEvent2.UserEmail = "sprov300@gmail.com"
	shipmentEvent2.RequestId = "bks1m1g91jau4nkks2f0"

	shipmentEvents := []*tntproto.ShipmentEventRequest{}
	shipmentEvents = append(shipmentEvents, &shipmentEvent1)
	shipmentEvents = append(shipmentEvents, &shipmentEvent2)

	equipEvents := tntproto.CreateShipmentEventRequest{}
	equipEvents.ShipmentEventRequests = shipmentEvents
	equipEvents.UserId = "auth0|66fd06d0bfea78a82bb42459"
	equipEvents.UserEmail = "sprov300@gmail.com"
	equipEvents.RequestId = "bks1m1g91jau4nkks2f0"

	type args struct {
		ctx context.Context
		in  *tntproto.CreateShipmentEventRequest
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
		_, err := tt.es.CreateShipmentEvent(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("EventService.CreateShipmentEvent() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
	}
}
