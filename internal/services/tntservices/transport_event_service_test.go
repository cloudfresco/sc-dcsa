package tntservices

import (
	"context"
	"testing"

	tntproto "github.com/cloudfresco/sc-dcsa/internal/protogen/tnt/v3"
	"github.com/cloudfresco/sc-dcsa/test"
)

func TestEventService_CreateTransportEvent(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	eventService := NewEventService(log, dbService, redisService, userServiceClient)

	transportEvent1 := tntproto.TransportEventRequest{}
	transportEvent1.EventClassifierCode = "ACT"
	transportEvent1.TransportEventTypeCode = "ARRI"
	transportEvent1.DelayReasonCode = "WEA"
	transportEvent1.ChangeRemark = "Bad weather"
	transportEvent1.TransportCallId = uint32(7)
	transportEvent1.EventCreatedDateTime = "11/15/2023"
	transportEvent1.EventDateTime = "12/11/2023"
	transportEvent1.UserId = "auth0|66fd06d0bfea78a82bb42459"
	transportEvent1.UserEmail = "sprov300@gmail.com"
	transportEvent1.RequestId = "bks1m1g91jau4nkks2f0"

	transportEvent2 := tntproto.TransportEventRequest{}
	transportEvent2.EventClassifierCode = "EST"
	transportEvent2.TransportEventTypeCode = "DEPA"
	transportEvent2.DelayReasonCode = "ANA"
	transportEvent2.ChangeRemark = "Authorities not available"
	transportEvent2.TransportCallId = uint32(8)
	transportEvent2.EventCreatedDateTime = "09/12/2024"
	transportEvent2.EventDateTime = "11/11/2024"
	transportEvent2.UserId = "auth0|66fd06d0bfea78a82bb42459"
	transportEvent2.UserEmail = "sprov300@gmail.com"
	transportEvent2.RequestId = "bks1m1g91jau4nkks2f0"

	transportEvents := []*tntproto.TransportEventRequest{}
	transportEvents = append(transportEvents, &transportEvent1)
	transportEvents = append(transportEvents, &transportEvent2)

	equipEvents := tntproto.CreateTransportEventRequest{}
	equipEvents.TransportEventRequests = transportEvents
	equipEvents.UserId = "auth0|66fd06d0bfea78a82bb42459"
	equipEvents.UserEmail = "sprov300@gmail.com"
	equipEvents.RequestId = "bks1m1g91jau4nkks2f0"

	type args struct {
		ctx context.Context
		in  *tntproto.CreateTransportEventRequest
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
		_, err := tt.es.CreateTransportEvent(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("EventService.CreateTransportEvent() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
	}
}
