package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rsbh/ecom/router"

	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	expected := "{\"message\": \"pong\"}"
	r := router.InitRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/ping", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.JSONEq(t, expected, w.Body.String())
}
