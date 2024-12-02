package v1

import (
	commonproto "github.com/cloudfresco/sc-dcsa/internal/protogen/common/v1"
	eventcoreproto "github.com/cloudfresco/sc-dcsa/internal/protogen/eventcore/v1"
	commonstruct "github.com/cloudfresco/sc-dcsa/internal/servicestructs/common/v1"
)

// UtilizedTransportEquipment - struct UtilizedTransportEquipment
type UtilizedTransportEquipment struct {
	*eventcoreproto.UtilizedTransportEquipmentD
	*commonproto.CrUpdUser
	*commonstruct.CrUpdTime
}

// Equipment - struct Equipment
type Equipment struct {
	*eventcoreproto.EquipmentD
	*commonproto.CrUpdUser
	*commonstruct.CrUpdTime
}
