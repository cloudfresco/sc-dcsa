package v1

import (
	commonproto "github.com/cloudfresco/sc-dcsa/internal/protogen/common/v1"
	eventcoreproto "github.com/cloudfresco/sc-dcsa/internal/protogen/eventcore/v1"
	commonstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/common/v1"
)

// Charge - struct Charge
type Charge struct {
	*eventcoreproto.ChargeD
	*commonproto.CrUpdUser
	*commonstruct.CrUpdTime
}
