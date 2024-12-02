package v1

import (
	commonproto "github.com/cloudfresco/sc-dcsa/internal/protogen/common/v1"
	jitproto "github.com/cloudfresco/sc-dcsa/internal/protogen/jit/v1"
	commonstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/common/v1"
)

// Service - struct Service
type Service struct {
	*jitproto.ServiceD
	*commonproto.CrUpdUser
	*commonstruct.CrUpdTime
}
