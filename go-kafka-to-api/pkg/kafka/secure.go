package kafka

import (
	"context"
	"crypto/tls"
	"os"

	"kafka_api_enterprise/internal/config"

	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/pkg/sasl/plain"
)

func NewSecureKafkaClient() (*kgo.Client, error) {
	username, err := config.LoadSecret(os.Getenv("KAFKA_USERNAME"))
	if err != nil {
		return nil, err
	}

	password, err := config.LoadSecret(os.Getenv("KAFKA_PASSWORD"))
	if err != nil {
		return nil, err
	}

	brokers := os.Getenv("KAFKA_BROKERS")
	topic := os.Getenv("KAFKA_TOPIC")
	group := os.Getenv("KAFKA_GROUP")

	mech := plain.Plain(func(ctx context.Context) (plain.Auth, error) {
		return plain.Auth{
			User: username,
			Pass: password,
		}, nil
	})

	kgo.SASL(mech)

	opts := []kgo.Opt{
		kgo.SeedBrokers(brokers),
		kgo.SASL(mech),
		kgo.DialTLSConfig(&tls.Config{InsecureSkipVerify: false}),
		kgo.ConsumeTopics(topic),
		kgo.ConsumerGroup(group),
	}

	client, err := kgo.NewClient(opts...)
	if err != nil {
		return nil, err
	}

	return client, nil
}
