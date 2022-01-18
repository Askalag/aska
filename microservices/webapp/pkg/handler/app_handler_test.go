package handler_test

import (
	"encoding/json"
	"github.com/go-playground/assert/v2"
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStatus(t *testing.T) {
	router := SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/app/status", nil)
	router.ServeHTTP(w, req)

	expected := map[string]interface{}{"status": "alive"}
	res := map[string]interface{}{}
	if err := json.NewDecoder(w.Body).Decode(&res); err != nil {
		log.Fatalln(err)
	}

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, expected, res)
}
