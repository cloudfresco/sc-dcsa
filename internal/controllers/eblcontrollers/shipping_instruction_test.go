package eblcontrollers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	bkgproto "github.com/cloudfresco/sc-dcsa/internal/protogen/bkg/v2"
	eblproto "github.com/cloudfresco/sc-dcsa/internal/protogen/ebl/v1"
	eventcoreproto "github.com/cloudfresco/sc-dcsa/internal/protogen/eventcore/v1"
	"github.com/cloudfresco/sc-dcsa/test"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func TestCreateShippingInstruction(t *testing.T) {
	err := test.LoadSQL(log, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	tokenString, email, backendServerAddr := LoginUser()

	w := httptest.NewRecorder()

	shippingInstruction := eblproto.CreateShippingInstructionRequest{}
	shippingInstruction.ShippingInstructionReference = "SI_REF_10"
	shippingInstruction.DocumentStatus = "DRFT"
	shippingInstruction.IsShippedOnboardType = true
	shippingInstruction.NumberOfCopies = uint32(2)
	shippingInstruction.NumberOfOriginals = uint32(4)
	shippingInstruction.IsElectronic = true
	shippingInstruction.IsToOrder = true
	shippingInstruction.AreChargesDisplayedOnOriginals = true
	shippingInstruction.AreChargesDisplayedOnCopies = false
	shippingInstruction.LocationId = uint32(0)
	shippingInstruction.TransportDocumentTypeCode = ""
	shippingInstruction.DisplayedNameForPlaceOfReceipt = ""
	shippingInstruction.DisplayedNameForPortOfLoad = ""
	shippingInstruction.DisplayedNameForPortOfDischarge = ""
	shippingInstruction.DisplayedNameForPlaceOfDelivery = ""
	shippingInstruction.AmendToTransportDocument = ""
	shippingInstruction.CreatedDateTime = "02/16/2023"
	shippingInstruction.UpdatedDateTime = "02/16/2023"

	utilizedTransportEquipment := eventcoreproto.CreateUtilizedTransportEquipmentRequest{}
	utilizedTransportEquipment.EquipmentReference = "BMOU2149612"
	utilizedTransportEquipment.CargoGrossWeight = float64(3000)
	utilizedTransportEquipment.CargoGrossWeightUnit = "KGM"
	utilizedTransportEquipment.IsShipperOwned = false

	equipment := eventcoreproto.CreateEquipmentRequest{}
	equipment.EquipmentReference = "BMOU2149612"
	equipment.IsoEquipmentCode = "22G1"
	equipment.TareWeight = float64(2000)
	equipment.WeightUnit = "KGM"

	utilizedTransportEquipment.Equipment = &equipment

	utilizedTransportEquipments := []*eventcoreproto.CreateUtilizedTransportEquipmentRequest{}
	utilizedTransportEquipments = append(utilizedTransportEquipments, &utilizedTransportEquipment)
	shippingInstruction.UtilizedTransportEquipments = utilizedTransportEquipments

	consignmentItem := eblproto.CreateConsignmentItemRequest{}
	consignmentItem.DescriptionOfGoods = "Leather Jackets"
	consignmentItem.HsCode = "411510"
	consignmentItem.Weight = float64(4000)
	consignmentItem.Volume = float64(20)
	consignmentItem.WeightUnit = "KGM"
	consignmentItem.VolumeUnit = "m"

	consignmentItems := []*eblproto.CreateConsignmentItemRequest{}
	consignmentItems = append(consignmentItems, &consignmentItem)
	shippingInstruction.ConsignmentItems = consignmentItems

	reference := bkgproto.CreateReferenceRequest{}
	reference.ReferenceTypeCode = "FF"
	reference.ReferenceValue = "test"

	references := []*bkgproto.CreateReferenceRequest{}
	references = append(references, &reference)
	shippingInstruction.References = references

	data, _ := json.Marshal(&shippingInstruction)

	req, err := http.NewRequest("POST", backendServerAddr+"/v2/shipping-instructions", bytes.NewBuffer(data))
	if err != nil {
		t.Error(err)
		return
	}

	req = common.SetEmailToken(req, tokenString, email)
	mux.ServeHTTP(w, req)

	resp := w.Result()
	// Check the status code is what we expect.
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Unexpected status code %d", resp.StatusCode)
		return
	}
	assert.NotNil(t, w.Body.String())
}

func TestUpdateShippingInstruction(t *testing.T) {
	err := test.LoadSQL(log, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	tokenString, email, backendServerAddr := LoginUser()

	w := httptest.NewRecorder()

	form := eblproto.UpdateShippingInstructionByShippingInstructionReferenceRequest{}
	form.ShippingInstructionReference = "SI_REF_4"
	form.DocumentStatus = "APPR"
	form.TransportDocumentTypeCode = ""
	form.DisplayedNameForPlaceOfReceipt = ""

	data, _ := json.Marshal(&form)

	req, err := http.NewRequest("PUT", backendServerAddr+"/v2/shipping-instructions/SI_REF_4", bytes.NewBuffer(data))
	if err != nil {
		t.Error(err)
		return
	}

	req = common.SetEmailToken(req, tokenString, email)
	mux.ServeHTTP(w, req)

	resp := w.Result()
	// Check the status code is what we expect.
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Unexpected status code %d", resp.StatusCode)
		return
	}

	expected := string(`"Updated Successfully"` + "\n")

	assert.Equal(t, w.Body.String(), expected, "they should be equal")
}
