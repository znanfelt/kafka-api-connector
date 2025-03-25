package api

import (
	"encoding/json"
	"errors"
	"kafka_api_enterprise/internal/config"
	"testing"
)

func init() {
	_ = LoadSchema("../../internal/config/schema.json")
	config.SetSecretDir("../../test-secrets")
}

func TestTransform_ValidInput(t *testing.T) {
	input := []byte(`{"id":"123","old_key":"hello","password":"hide"}`)
	out, err := Transform(input)
	if err != nil {
		t.Fatalf("Expected success, got err: %v", err)
	}
	var result map[string]interface{}
	json.Unmarshal(out, &result)
	if result["new_key"] != "hello" {
		t.Error("Expected renamed field new_key")
	}
	if _, exists := result["password"]; exists {
		t.Error("Expected password to be removed")
	}
	if result["source"] != "redpanda-franz-pipeline" {
		t.Error("Expected static tag 'source'")
	}
}

func TestTransform_InvalidJSON(t *testing.T) {
	_, err := Transform([]byte(`{this is invalid}`))
	if err == nil {
		t.Error("Expected JSON parse error")
	}
}

func TestTransform_MissingRequiredField(t *testing.T) {
	input := []byte(`{"old_key":"hello"}`)
	_, err := Transform(input)
	if err == nil || !errors.Is(err, err) {
		t.Error("Expected missing required field error")
	}
}
