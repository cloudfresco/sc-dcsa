package v3

import (
	commonproto "github.com/cloudfresco/sc-dcsa/internal/proto-gen/common/v1"
	tntproto "github.com/cloudfresco/sc-dcsa/internal/proto-gen/tnt/v3"
	commonstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/common/v1"
)

// Seal - struct Seal
type Seal struct {
	*tntproto.SealD
	*commonproto.CrUpdUser
	*commonstruct.CrUpdTime
}
