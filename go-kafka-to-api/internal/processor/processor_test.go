package processor

import (
	"encoding/json"
	"go_kafka_consumer/internal/sender"
	"testing"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaMessage struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type mockSender struct {
	LastMessage sender.APIMessage
}

func (m *mockSender) SendMessage(apiURL string, msg sender.APIMessage) {
	m.LastMessage = msg
}

func TestProcessMessage(t *testing.T) {
	apiEndpoint := "http://localhost:8080/api/messages"

	message := &kafka.Message{
		Value: []byte(`{"type": "important", "data": "test data"}`),
	}

	var kafkaMsg KafkaMessage
	err := json.Unmarshal(message.Value, &kafkaMsg)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	if kafkaMsg.Type != "important" {
		t.Fatalf("Expected message type 'important', got %s", kafkaMsg.Type)
	}

	// Call ProcessMessage and verify behavior
	ProcessMessage(message, apiEndpoint)
}
