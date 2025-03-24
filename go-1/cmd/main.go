package main

import (
	"go_kafka_consumer/internal/config"
	"go_kafka_consumer/internal/consumer"
	"log"
)

func main() {
	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	consumer.StartConsumer(cfg)
}
