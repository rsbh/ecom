package api_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rsbh/ecom/router"

	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	logger := log.Default()
	expected := "{\"message\": \"pong\"}"
	r := router.InitRouter(logger)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/ping", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.JSONEq(t, expected, w.Body.String())
}
