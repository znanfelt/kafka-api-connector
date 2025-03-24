package processor

import (
	"encoding/json"
	"go_kafka_consumer/internal/sender"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

func ProcessMessage(msg *kafka.Message, apiEndpoint string) {
	var m Message
	err := json.Unmarshal(msg.Value, &m)
	if err != nil {
		log.Printf("Failed to parse message: %v", err)
		return
	}

	if m.Type == "important" {
		sender.SendMessage(sender.APIMessage{Type: m.Type, Data: m.Data}, apiEndpoint)
	}
}
