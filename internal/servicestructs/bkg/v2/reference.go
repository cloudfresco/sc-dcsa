package v2

import (
	bkgproto "github.com/cloudfresco/sc-dcsa/internal/proto-gen/bkg/v2"
	commonproto "github.com/cloudfresco/sc-dcsa/internal/proto-gen/common/v1"
	commonstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/common/v1"
)

// Reference1 - struct Reference1
type Reference1 struct {
	*bkgproto.Reference1D
	*commonproto.CrUpdUser
	*commonstruct.CrUpdTime
}