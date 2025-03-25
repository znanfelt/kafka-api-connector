package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"kafka_api_enterprise/pkg/api"
)

func TestChaosLatencyInjection(t *testing.T) {
	apiServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(3 * time.Second) // simulate slowness
		w.WriteHeader(http.StatusOK)
	}))
	defer apiServer.Close()

	os.Setenv("API_AUTH_TOKEN", "")
	_ = api.InitAPIAuth()

	err := api.PostToAPI(apiServer.URL, []byte(`{"simulate":"slow"}`))
	if err == nil {
		t.Error("Expected timeout or error on slow API")
	}
}
