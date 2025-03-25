package api

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestPostToAPI_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "Bearer test-token" {
			t.Error("Missing or incorrect auth header")
		}
		io.Copy(io.Discard, r.Body)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	os.Setenv("API_AUTH_TOKEN", "test-token")

	_ = InitAPIAuth()
	err := PostToAPI(server.URL, []byte(`{"test":"data"}`))
	if err != nil {
		t.Errorf("Expected successful POST, got: %v", err)
	}
}

func TestPostToAPI_Failure(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "fail", http.StatusInternalServerError)
	}))
	defer server.Close()

	err := PostToAPI(server.URL, []byte(`{"fail":"now"}`))
	if err == nil {
		t.Error("Expected error from failed API call")
	}
}
