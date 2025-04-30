package v3

import (
	commonproto "github.com/cloudfresco/sc-dcsa/internal/protogen/common/v1"
	ovsproto "github.com/cloudfresco/sc-dcsa/internal/protogen/ovs/v3"
	commonstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/common/v1"
)

// VesselSchedule - struct VesselSchedule
type VesselSchedule struct {
	*ovsproto.VesselScheduleD
	*commonproto.CrUpdUser
	*commonstruct.CrUpdTime
}

// Vessel - struct Vessel
type Vessel struct {
	*ovsproto.VesselD
	*commonproto.CrUpdUser
	*commonstruct.CrUpdTime
}
