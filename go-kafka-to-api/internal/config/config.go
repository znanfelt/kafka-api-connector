package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	KafkaBrokers string `json:"kafka_brokers"`
	KafkaTopic   string `json:"kafka_topic"`
	APIEndpoint  string `json:"api_endpoint"`
	GroupID      string `json:"group_id"`
}

func LoadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	cfg := &Config{}
	err = decoder.Decode(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
