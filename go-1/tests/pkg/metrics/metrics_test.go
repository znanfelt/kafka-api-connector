package metrics

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestMetricsHandler(t *testing.T) {
    req, err := http.NewRequest("GET", "/metrics", nil)
    if err != nil {
        t.Fatalf("Failed to create request: %v", err)
    }

    rr := httptest.NewRecorder()
    handler := Handler()
    handler.ServeHTTP(rr, req)

    if rr.Code != http.StatusOK {
        t.Errorf("Expected status 200, got %d", rr.Code)
    }

    if len(rr.Body.String()) == 0 {
        t.Error("Expected metrics output, got empty response")
    }
}
