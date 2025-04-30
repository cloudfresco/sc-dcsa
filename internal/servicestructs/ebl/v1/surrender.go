package v1

import (
	"time"

	commonproto "github.com/cloudfresco/sc-dcsa/internal/protogen/common/v1"
	eblproto "github.com/cloudfresco/sc-dcsa/internal/protogen/ebl/v1"
	commonstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/common/v1"
)

// TransactionParty - struct TransactionParty
type TransactionParty struct {
	*eblproto.TransactionPartyD
	*commonproto.CrUpdUser
	*commonstruct.CrUpdTime
}

// TransactionPartySupportingCode - struct TransactionPartySupportingCode
type TransactionPartySupportingCode struct {
	*eblproto.TransactionPartySupportingCodeD
	*commonproto.CrUpdUser
	*commonstruct.CrUpdTime
}

// SurrenderRequest - struct SurrenderRequest
type SurrenderRequest struct {
	*eblproto.SurrenderRequestD
	*SurrenderRequestT
	*commonproto.CrUpdUser
	*commonstruct.CrUpdTime
}

// SurrenderRequestT - struct SurrenderRequestT
type SurrenderRequestT struct {
	CreatedDateTime time.Time `protobuf:"bytes,6,opt,name=created_date_time,json=createdDateTime,proto3" json:"created_date_time,omitempty"`
}

// SurrenderRequestAnswer - struct SurrenderRequestAnswer
type SurrenderRequestAnswer struct {
	*eblproto.SurrenderRequestAnswerD
	*SurrenderRequestAnswerT
	*commonproto.CrUpdUser
	*commonstruct.CrUpdTime
}

// SurrenderRequestAnswerT - struct SurrenderRequestAnswerT
type SurrenderRequestAnswerT struct {
	CreatedDateTime time.Time `protobuf:"bytes,1,opt,name=created_date_time,json=createdDateTime,proto3" json:"created_date_time,omitempty"`
}

// EndorsementChainLink - struct EndorsementChainLink
type EndorsementChainLink struct {
	*eblproto.EndorsementChainLinkD
	*EndorsementChainLinkT
	*commonproto.CrUpdUser
	*commonstruct.CrUpdTime
}

// EndorsementChainLinkT - struct EndorsementChainLinkT
type EndorsementChainLinkT struct {
	ActionDateTime time.Time `protobuf:"bytes,1,opt,name=action_date_time,json=actionDateTime,proto3" json:"action_date_time,omitempty"`
}
