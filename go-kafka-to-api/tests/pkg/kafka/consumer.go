package kafka

import (
    "context"
    "log"
    "os"
    "time"

    "github.com/twmb/franz-go/pkg/kgo"
    "kafka_api_enterprise/pkg/api"
    "kafka_api_enterprise/pkg/dedup"
    "kafka_api_enterprise/pkg/metrics"
)

func StartConsumer() {
    client, err := NewSecureKafkaClient()
    if err != nil {
        log.Fatalf("failed to create Kafka client: %v", err)
    }
    defer client.Close()

    dlq := NewDLQWriter(client)

    filter := dedup.NewRotatingBloomFilter(2*time.Minute, 100000, 0.01)
    ticker := time.NewTicker(30 * time.Second)
    defer ticker.Stop()

    ctx := context.Background()
    apiURL := os.Getenv("API_URL")

    go func() {
        for range ticker.C {
            filter.Rotate()
        }
    }()

    for {
        fetches := client.PollRecords(ctx, 100)
        fetches.EachPartition(func(p kgo.FetchTopicPartition) {
            for _, record := range p.Records {
                if filter.Exists(record.Key) {
                    continue // skip duplicates
                }
                filter.Add(record.Key)

                transformed, err := api.Transform(record.Value)
                if err != nil {
                    log.Printf("transform error: %v", err)
                    continue
                }

                err = api.PostToAPI(apiURL, transformed)
                if err != nil {
                    log.Printf("API post error: %v", err)
                    dlq.SendToDLQ(record.Value)
                    metrics.MessagesFailed.Inc()
                } else {
                    metrics.MessagesPosted.Inc()
                }

                metrics.MessagesConsumed.Inc()
            }
        })
    }
}
