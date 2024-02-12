package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignup(t *testing.T) {
	payload := []byte(`{"email": "test@test.com", "password": "secret", "first_name": "test", "last_name": "test"}`)

	req, err := http.NewRequest("POST", "/users/signup", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)
}
