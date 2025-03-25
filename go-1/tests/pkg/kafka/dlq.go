package kafka

import (
	"context"
	"log"
	"os"

	"github.com/twmb/franz-go/pkg/kgo"
)

type DLQWriter interface {
	SendToDLQ(value []byte)
}

type KafkaDLQ struct {
	Client KafkaClient
	Topic  string
}

type KafkaClient interface {
	Produce(ctx context.Context, record *kgo.Record, cb func(*kgo.Record, error))
}

func (d *KafkaDLQ) SendToDLQ(value []byte) {
	d.Client.Produce(context.Background(), &kgo.Record{
		Topic: d.Topic,
		Value: value,
	}, func(_ *kgo.Record, err error) {
		if err != nil {
			log.Printf("Failed to send to DLQ: %v", err)
		}
	})
}

// Default entry point
func NewDLQWriter(client *kgo.Client) DLQWriter {
	return &KafkaDLQ{
		Client: client,
		Topic:  os.Getenv("KAFKA_DLQ_TOPIC"),
	}
}
