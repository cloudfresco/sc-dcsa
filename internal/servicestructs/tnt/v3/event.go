package v3

import (
	"time"

	tntproto "github.com/cloudfresco/sc-dcsa/internal/protogen/tnt/v3"
)

// Event - struct Event
type Event struct {
	*tntproto.EventD
	*EventT
}

// EventT - struct EventT
type EventT struct {
	EventCreatedDateTime time.Time `protobuf:"bytes,1,opt,name=event_created_date_time,json=eventCreatedDateTime,proto3" json:"event_created_date_time,omitempty"`
	EventDateTime        time.Time `protobuf:"bytes,2,opt,name=event_date_time,json=eventDateTime,proto3" json:"event_date_time,omitempty"`
}
