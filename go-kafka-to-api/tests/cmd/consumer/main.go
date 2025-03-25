package main

import (
    "github.com/spf13/cobra"
    "kafka_api_enterprise/internal/config"
    "kafka_api_enterprise/pkg/kafka"
)

var rootCmd = &cobra.Command{
    Use:   "consumer",
    Short: "Kafka to API consumer",
    Run: func(cmd *cobra.Command, args []string) {
        config.LoadConfig()
        kafka.StartConsumer()
    },
}

func Execute() error {
    return rootCmd.Execute()
}

func NewRootCommand() *cobra.Command {
    return rootCmd
}

func main() {
    _ = Execute()
}
