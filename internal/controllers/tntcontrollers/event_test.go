package tntcontrollers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	tntproto "github.com/cloudfresco/sc-dcsa/internal/protogen/tnt/v3"
	"github.com/cloudfresco/sc-dcsa/test"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func TestCreateEquipmentEvent(t *testing.T) {
	err := test.LoadSQL(log, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	tokenString, email, backendServerAddr := LoginUser()

	w := httptest.NewRecorder()

	equipmentEvent1 := tntproto.EquipmentEventRequest{}
	equipmentEvent1.EventClassifierCode = "ACT"
	equipmentEvent1.EquipmentEventTypeCode = "LOAD"
	equipmentEvent1.EquipmentReference = "equipref3453"
	equipmentEvent1.EmptyIndicatorCode = "EMPTY"
	equipmentEvent1.TransportCallId = uint32(6)
	equipmentEvent1.EventLocation = ""
	equipmentEvent1.EventCreatedDateTime = "11/15/2023"
	equipmentEvent1.EventDateTime = "12/11/2023"

	equipmentEvent2 := tntproto.EquipmentEventRequest{}
	equipmentEvent2.EventClassifierCode = "EST"
	equipmentEvent2.EquipmentEventTypeCode = "LOAD"
	equipmentEvent2.EquipmentReference = "APZU4812090"
	equipmentEvent2.EmptyIndicatorCode = "EMPTY"
	equipmentEvent2.TransportCallId = uint32(6)
	equipmentEvent2.EventLocation = ""
	equipmentEvent2.EventCreatedDateTime = "09/12/2024"
	equipmentEvent2.EventDateTime = "11/11/2024"

	equipmentEvents := []*tntproto.EquipmentEventRequest{}
	equipmentEvents = append(equipmentEvents, &equipmentEvent1)
	equipmentEvents = append(equipmentEvents, &equipmentEvent2)

	equipEvents := tntproto.CreateEquipmentEventRequest{}
	equipEvents.EquipmentEventRequests = equipmentEvents

	data, _ := json.Marshal(&equipEvents)

	req, err := http.NewRequest("POST", backendServerAddr+"/v3/events/equipment-event", bytes.NewBuffer(data))
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
