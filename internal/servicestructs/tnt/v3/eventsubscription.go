package v3

import (
	commonproto "github.com/cloudfresco/sc-dcsa/internal/proto-gen/common/v1"
	tntproto "github.com/cloudfresco/sc-dcsa/internal/proto-gen/tnt/v3"
	commonstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/common/v1"
)

// EventSubscription - struct EventSubscription
type EventSubscription struct {
	*tntproto.EventSubscriptionD
	*commonproto.CrUpdUser
	*commonstruct.CrUpdTime
}
