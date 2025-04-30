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

func TestCreateEventSubscription(t *testing.T) {
	err := test.LoadSQL(log, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	tokenString, email, backendServerAddr := LoginUser()

	w := httptest.NewRecorder()

	eventSubscription := tntproto.CreateEventSubscriptionRequest{}

	eventSubscription.CallbackUrl = ""
	eventSubscription.DocumentReference = ""
	eventSubscription.EquipmentReference = ""
	eventSubscription.TransportCallReference = ""
	eventSubscription.VesselImoNumber = ""
	eventSubscription.CarrierExportVoyageNumber = ""
	eventSubscription.UniversalExportVoyageReference = ""
	eventSubscription.CarrierServiceCode = ""
	eventSubscription.UniversalServiceReference = ""
	eventSubscription.UnLocationCode = ""
	// eventSubscription.Secret = in.Secret

	data, _ := json.Marshal(&eventSubscription)

	req, err := http.NewRequest("POST", backendServerAddr+"/v3/event-subscriptions", bytes.NewBuffer(data))
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
