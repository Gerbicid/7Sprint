package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainHandlerWhenNoSuchCity(t *testing.T) {

	req := httptest.NewRequest("GET", "/cafe?count=2&city=tula", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	body := responseRecorder.Body.String()

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

	assert.Equal(t, "wrong city value", body)
}
