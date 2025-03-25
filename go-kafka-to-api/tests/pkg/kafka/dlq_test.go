package kafka

import (
    "context"
    "errors"
    "testing"

    "github.com/twmb/franz-go/pkg/kgo"
)

type mockKafkaClient struct {
    produced bool
    err      error
}

func (m *mockKafkaClient) Produce(_ context.Context, record *kgo.Record, cb func(*kgo.Record, error)) {
    m.produced = true
    cb(record, m.err)
}

func TestKafkaDLQ_SendToDLQ_Success(t *testing.T) {
    client := &mockKafkaClient{err: nil}
    dlq := &KafkaDLQ{
        Client: client,
        Topic:  "dlq-topic",
    }

    dlq.SendToDLQ([]byte("test-message"))
    if !client.produced {
        t.Error("Expected message to be produced")
    }
}

func TestKafkaDLQ_SendToDLQ_Error(t *testing.T) {
    client := &mockKafkaClient{err: errors.New("produce failed")}
    dlq := &KafkaDLQ{
        Client: client,
        Topic:  "dlq-topic",
    }

    dlq.SendToDLQ([]byte("test-error"))
    if !client.produced {
        t.Error("Expected message to be produced even on error")
    }
}
