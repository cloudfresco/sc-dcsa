package jitcontrollers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	jitproto "github.com/cloudfresco/sc-dcsa/internal/protogen/jit/v1"
	"github.com/cloudfresco/sc-dcsa/test"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func TestCreateTimestamp(t *testing.T) {
	fmt.Println("TestCreateTimestamp started")
	err := test.LoadSQL(log, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	tokenString := LoginUser()
	fmt.Println("TestCreateTimestamp tokenString", tokenString)

	w := httptest.NewRecorder()

	tstamp := jitproto.CreateTimestampRequest{}
	tstamp.EventTypeCode = "ARRI"
	tstamp.EventClassifierCode = "ACT"
	tstamp.DelayReasonCode = "ANA"
	tstamp.ChangeRemark = "Authorities not available"
	tstamp.EventDateTime = "11/09/2023"

	data, _ := json.Marshal(&tstamp)

	req, err := http.NewRequest("POST", "https://localhost:9061/v0.1/timestamps", bytes.NewBuffer(data))
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
