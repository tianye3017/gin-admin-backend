package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tianye3017/gin-admin-backend/engine"
	_ "github.com/tianye3017/gin-admin-backend/router"
)

func TestPingRoute(t *testing.T) {
	router := engine.Router

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
