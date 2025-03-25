//go:build integration

package main

import (
    "context"
    "net/http"
    "net/http/httptest"
    "os"
    "testing"
    "time"

    tcKafka "github.com/testcontainers/testcontainers-go/modules/kafka"
)

func TestKafkaToAPI_Integration(t *testing.T) {
    ctx := context.Background()

    // Start mock API server
    apiServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
    }))
    defer apiServer.Close()

    // Start Kafka container
    kafkaContainer, err := tcKafka.RunContainer(ctx)
    if err != nil {
        t.Fatalf("Kafka container failed: %v", err)
    }
    defer kafkaContainer.Terminate(ctx)

    brokers, err := kafkaContainer.BootstrapServers(ctx)
    if err != nil {
        t.Fatalf("Kafka bootstrap error: %v", err)
    }

    // Inject env vars
    os.Setenv("KAFKA_BROKERS", brokers)
    os.Setenv("KAFKA_TOPIC", "test-topic")
    os.Setenv("KAFKA_GROUP", "test-consumer")
    os.Setenv("API_AUTH_TOKEN", "")
    os.Setenv("API_URL", apiServer.URL)

    // Create topic
    err = kafkaContainer.CreateTopic(ctx, "test-topic")
    if err != nil {
        t.Fatalf("Topic creation failed: %v", err)
    }

    // Start the consumer in background
    go func() {
        time.Sleep(2 * time.Second) // wait for Kafka to settle
        StartConsumer()
    }()

    // Allow time for processing (simplified)
    time.Sleep(5 * time.Second)
}
