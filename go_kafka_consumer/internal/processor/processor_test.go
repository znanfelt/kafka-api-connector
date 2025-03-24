package processor

import (
	"testing"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func TestProcessMessage(t *testing.T) {
	message := &kafka.Message{
		Value: []byte(`{"type": "important", "data": "test data"}`),
	}

	apiEndpoint := "http://localhost:8080/api/messages"
	ProcessMessage(message, apiEndpoint)
}
