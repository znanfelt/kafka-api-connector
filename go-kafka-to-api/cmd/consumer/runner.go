package main

import (
    "log"
    "net/http"

    "kafka_api_enterprise/pkg/api"
    "kafka_api_enterprise/pkg/kafka"
    "kafka_api_enterprise/pkg/metrics"
)

func RunConsumerCommand(schemaPath string) error {
    if err := api.LoadSchema(schemaPath); err != nil {
        return err
    }

    go func() {
        log.Println("Prometheus metrics at :2112/metrics")
        http.Handle("/metrics", metrics.Handler())
        http.ListenAndServe(":2112", nil)
    }()

    kafka.StartConsumer()
    return nil
}
