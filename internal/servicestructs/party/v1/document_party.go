package v1

import (
	commonproto "github.com/cloudfresco/sc-dcsa/internal/proto-gen/common/v1"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/proto-gen/party/v1"
	commonstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/common/v1"
)

// DocumentParty - struct DocumentParty
type DocumentParty struct {
	*partyproto.DocumentPartyD
	*commonproto.CrUpdUser
	*commonstruct.CrUpdTime
}