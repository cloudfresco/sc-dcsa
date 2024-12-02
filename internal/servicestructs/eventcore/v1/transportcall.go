package v1

import (
	commonproto "github.com/cloudfresco/sc-dcsa/internal/proto-gen/common/v1"
	eventcoreproto "github.com/cloudfresco/sc-dcsa/internal/proto-gen/eventcore/v1"
	commonstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/common/v1"
)

// TransportCall - struct TransportCall
type TransportCall struct {
	*eventcoreproto.TransportCallD
	*commonproto.CrUpdUser
	*commonstruct.CrUpdTime
}
