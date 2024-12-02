package tntcontrollers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	tntproto "github.com/cloudfresco/sc-dcsa/internal/proto-gen/tnt/v3"
	"github.com/cloudfresco/sc-dcsa/test"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func TestCreateEventSubscription(t *testing.T) {
	fmt.Println("TestCreateEventSubscription started")
	err := test.LoadSQL(log, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	tokenString := LoginUser()
	fmt.Println("TestCreateEventSubscription tokenString", tokenString)

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

	req, err := http.NewRequest("POST", "https://localhost:9061/v3/event-subscriptions", bytes.NewBuffer(data))
	if err != nil {
		t.Error(err)
		return
	}

	req.Header.Set("authorization", "Bearer "+tokenString)
	mux.ServeHTTP(w, req)

	resp := w.Result()
	// Check the status code is what we expect.
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Unexpected status code %d", resp.StatusCode)
		return
	}
	assert.NotNil(t, w.Body.String())
}
