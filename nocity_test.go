package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainHandlerWhenNoSuchCity(t *testing.T) {

	expected := "moscow"
	req := httptest.NewRequest("GET", "/cafe?count=6&city=tula", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	body := responseRecorder.Body.String()

	assert.NotEqual(t, expected, body)
	fmt.Println(http.StatusBadRequest, body)
}
