package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainHandlerWhenNoSuchCity(t *testing.T) {

	expected := "moscow"
	req := httptest.NewRequest("GET", "/cafe?count=6&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	body := responseRecorder.Body.String()

	assert.NotEmpty(t, body)
	assert.Equal(t, expected, body)
	//assert.Error()
	//fmt.Println(http.StatusBadRequest, body)
}
