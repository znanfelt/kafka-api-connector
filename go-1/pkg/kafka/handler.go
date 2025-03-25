package kafka

import (
    "log"
    "os"

    "github.com/twmb/franz-go/pkg/kgo"
    "kafka_api_enterprise/pkg/api"
    "kafka_api_enterprise/pkg/metrics"
)

var apiPost = api.PostToAPI
var apiTransform = api.Transform

func HandleRecord(record *kgo.Record, dlq DLQWriter) {
    apiURL := os.Getenv("API_URL")

    transformed, err := apiTransform(record.Value)
    if err != nil {
        log.Printf("transform error: %v", err)
        return
    }

    err = apiPost(apiURL, transformed)
    if err != nil {
        log.Printf("API post error: %v", err)
        dlq.SendToDLQ(record.Value)
        metrics.MessagesFailed.Inc()
    } else {
        metrics.MessagesPosted.Inc()
    }

    metrics.MessagesConsumed.Inc()
}
