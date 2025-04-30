package jitcontrollers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	jitproto "github.com/cloudfresco/sc-dcsa/internal/protogen/jit/v1"
	"github.com/cloudfresco/sc-dcsa/test"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func TestCreateTimestamp(t *testing.T) {
	err := test.LoadSQL(log, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	tokenString, email, backendServerAddr := LoginUser()

	w := httptest.NewRecorder()

	tstamp := jitproto.CreateTimestampRequest{}
	tstamp.EventTypeCode = "ARRI"
	tstamp.EventClassifierCode = "ACT"
	tstamp.DelayReasonCode = "ANA"
	tstamp.ChangeRemark = "Authorities not available"
	tstamp.EventDateTime = "11/09/2023"

	data, _ := json.Marshal(&tstamp)

	req, err := http.NewRequest("POST", backendServerAddr+"/v1/timestamps", bytes.NewBuffer(data))
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
