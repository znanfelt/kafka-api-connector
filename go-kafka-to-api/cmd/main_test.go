package main

import (
	"go_kafka_consumer/internal/config"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	cfg, err := config.LoadConfig("../config.json")
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	if len(cfg.KafkaBrokers) == 0 {
		t.Errorf("Expected KafkaBrokers to be set, but got empty list")
	}

	if cfg.KafkaTopic == "" {
		t.Errorf("Expected KafkaTopic to be set, but got empty string")
	}

	if cfg.APIEndpoint == "" {
		t.Errorf("Expected APIEndpoint to be set, but got empty string")
	}

	if cfg.GroupID == "" {
		t.Errorf("Expected GroupID to be set, but got empty string")
	}
}
