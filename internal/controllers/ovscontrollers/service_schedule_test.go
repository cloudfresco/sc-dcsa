package ovscontrollers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	ovsproto "github.com/cloudfresco/sc-dcsa/internal/proto-gen/ovs/v3"
	"github.com/cloudfresco/sc-dcsa/test"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func TestCreateServiceSchedule(t *testing.T) {
	fmt.Println("TestCreateServiceSchedule started")
	err := test.LoadSQL(log, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	tokenString := LoginUser()

	w := httptest.NewRecorder()

	serviceSchedule := ovsproto.CreateServiceScheduleRequest{}
	serviceSchedule.CarrierServiceName = "B_carrier_service_name_1"
	serviceSchedule.CarrierServiceCode = "B_HLC"
	serviceSchedule.UniversalServiceReference = "SR00003H"

	data, _ := json.Marshal(&serviceSchedule)

	req, err := http.NewRequest("POST", "https://localhost:9061/v3/service-schedules", bytes.NewBuffer(data))
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

func TestUpdateServiceSchedule(t *testing.T) {
	err := test.LoadSQL(log, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	tokenString := LoginUser()

	w := httptest.NewRecorder()

	form := ovsproto.UpdateServiceScheduleByUniversalServiceReferenceRequest{}
	form.CarrierServiceName = "C_carrier_service_name"
	form.CarrierServiceCode = "C_HLC"

	data, _ := json.Marshal(&form)

	req, err := http.NewRequest("PUT", "https://localhost:9061/v3/service-schedules/SR00002B", bytes.NewBuffer(data))
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

	expected := string(`"Updated Successfully"` + "\n")
	assert.Equal(t, w.Body.String(), expected, "they should be equal")
}