package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sayedatif/tigerhall/types"
	"github.com/stretchr/testify/assert"
)

var tigerId int64

func TestCreateTiger(t *testing.T) {
	payload := []byte(`{"name": "Test Tiger", "dob": "2024-01-01", "last_seen_at": "2024-01-01 20:30:59.026", "last_seen_lat": 23.22, "last_seen_long": 80.854}`)

	req, err := http.NewRequest("POST", "/tigers", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", fmt.Sprintf(`Bearer %s`, authToken))

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)
	var createTigerResponse types.CreateTigerResponse
	err = json.NewDecoder(rr.Body).Decode(&createTigerResponse)
	if err != nil {
		t.Errorf("error decoding response body: %v", err)
	}

	expected := "Created new tiger successfully"
	if createTigerResponse.Message != expected {
		t.Errorf("expected message %q, got %q", expected, createTigerResponse.Message)
	}
	tigerId = createTigerResponse.Data.TigerId
}
