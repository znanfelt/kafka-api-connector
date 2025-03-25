package consumer

import (
	"testing"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// Mock Kafka Consumer
type MockConsumer struct {
	Messages chan *kafka.Message
}

func (mc *MockConsumer) ReadMessage(timeoutMs int) (*kafka.Message, error) {
	msg, ok := <-mc.Messages
	if !ok {
		return nil, kafka.NewError(kafka.ErrUnknown, "No more messages", false)
	}
	return msg, nil
}

func TestStartConsumer(t *testing.T) {
	mockConsumer := &MockConsumer{Messages: make(chan *kafka.Message, 1)}
	mockConsumer.Messages <- &kafka.Message{
		Key:   []byte("test-key"),
		Value: []byte(`{"type": "important", "data": "test data"}`),
	}

	close(mockConsumer.Messages)

	// Simulating message processing
	go func() {
		for msg := range mockConsumer.Messages {
			if msg == nil {
				t.Errorf("Expected message but got nil")
			}
		}
	}()
}
