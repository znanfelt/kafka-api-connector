package api

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

func PostToAPI(url string, payload []byte) error {
	client := &http.Client{Timeout: 2 * time.Second}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	if getAuthHeader() != "" {
		req.Header.Set("Authorization", getAuthHeader())
	}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		return fmt.Errorf("API call failed: %v", err)
	}
	return nil
}
