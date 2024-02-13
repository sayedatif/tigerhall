package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type LoginResponse struct {
	Data struct {
		Token string `json:"token"`
	} `json:"data"`
}

var authToken string

func TestLogin(t *testing.T) {
	payload := []byte(`{"email": "test@test.com", "password": "secret"}`)

	req, err := http.NewRequest("POST", "/users/login", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)

	var loginResponse LoginResponse
	err = json.NewDecoder(rr.Body).Decode(&loginResponse)
	if err != nil {
		t.Errorf("error decoding response body: %v", err)
	}
	authToken = loginResponse.Data.Token
}
