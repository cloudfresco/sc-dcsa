package v3

import (
	commonproto "github.com/cloudfresco/sc-dcsa/internal/protogen/common/v1"
	ovsproto "github.com/cloudfresco/sc-dcsa/internal/protogen/ovs/v3"
	commonstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/common/v1"
)

// Leg - struct Leg
type Leg struct {
	*ovsproto.LegD
	*commonproto.CrUpdUser
	*commonstruct.CrUpdTime
}

// PointToPointRouting - struct PointToPointRouting
type PointToPointRouting struct {
	*ovsproto.PointToPointRoutingD
	*commonproto.CrUpdUser
	*commonstruct.CrUpdTime
}
