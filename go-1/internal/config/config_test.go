package config

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	os.WriteFile("test_config.json", []byte(`{"kafka_brokers": "localhost:9092", "kafka_topic": "test-topic", "api_endpoint": "http://localhost:8080/api", "group_id": "test-group"}`), 0644)
	defer os.Remove("test_config.json")

	cfg, err := LoadConfig("test_config.json")
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	if cfg.KafkaBrokers != "localhost:9092" {
		t.Errorf("Expected localhost:9092, got %s", cfg.KafkaBrokers)
	}
}
