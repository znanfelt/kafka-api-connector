package consumer

import (
	"context"
	"log"
	"go_kafka_consumer/internal/config"
	"go_kafka_consumer/internal/processor"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func StartConsumer(cfg *config.Config) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": cfg.KafkaBrokers,
		"group.id":          cfg.GroupID,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		log.Fatalf("Failed to create Kafka consumer: %v", err)
	}

	defer c.Close()

	c.SubscribeTopics([]string{cfg.KafkaTopic}, nil)
	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			processor.ProcessMessage(msg, cfg.APIEndpoint)
		} else {
			log.Printf("Consumer error: %v", err)
		}
	}
}
