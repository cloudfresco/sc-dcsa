package v3

import (
	commonproto "github.com/cloudfresco/sc-dcsa/internal/protogen/common/v1"
	ovsproto "github.com/cloudfresco/sc-dcsa/internal/protogen/ovs/v3"
	commonstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/common/v1"
)

// ServiceSchedule - struct ServiceSchedule
type ServiceSchedule struct {
	*ovsproto.ServiceScheduleD
	*commonproto.CrUpdUser
	*commonstruct.CrUpdTime
}
