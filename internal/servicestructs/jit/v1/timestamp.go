package v1

import (
	"time"

	jitproto "github.com/cloudfresco/sc-dcsa/internal/protogen/jit/v1"
)

// Timestamp - struct Timestamp
type Timestamp struct {
	*jitproto.TimestampD
	*TimestampT
}

// TimestampT - struct TimestampT
type TimestampT struct {
	EventDateTime time.Time `protobuf:"bytes,1,opt,name=event_date_time,json=eventDateTime,proto3" json:"event_date_time,omitempty"`
}
