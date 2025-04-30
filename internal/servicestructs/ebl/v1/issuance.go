package v1

import (
	"time"

	commonproto "github.com/cloudfresco/sc-dcsa/internal/protogen/common/v1"
	eblproto "github.com/cloudfresco/sc-dcsa/internal/protogen/ebl/v1"
	commonstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/common/v1"
)

// IssueParty - struct IssueParty
type IssueParty struct {
	*eblproto.IssuePartyD
	*commonproto.CrUpdUser
	*commonstruct.CrUpdTime
}

// IssuePartySupportingCode - struct IssuePartySupportingCode
type IssuePartySupportingCode struct {
	*eblproto.IssuePartySupportingCodeD
	*commonproto.CrUpdUser
	*commonstruct.CrUpdTime
}

// IssuanceRequest - struct IssuanceRequest
type IssuanceRequest struct {
	*eblproto.IssuanceRequestD
	*commonproto.CrUpdUser
	*commonstruct.CrUpdTime
}

// IssuanceRequestResponse - struct IssuanceRequestResponse
type IssuanceRequestResponse struct {
	*eblproto.IssuanceRequestResponseD
	*IssuanceRequestResponseT
	*commonproto.CrUpdUser
	*commonstruct.CrUpdTime
}

// IssuanceRequestResponseT - struct IssuanceRequestResponseT
type IssuanceRequestResponseT struct {
	CreatedDateTime time.Time `protobuf:"bytes,1,opt,name=created_date_time,json=createdDateTime,proto3" json:"created_date_time,omitempty"`
}

// EblVisualization - struct EblVisualization
type EblVisualization struct {
	*eblproto.EblVisualizationD
	*commonproto.CrUpdUser
	*commonstruct.CrUpdTime
}
