package config

import (
    "github.com/spf13/viper"
    "log"
)

type AppConfig struct {
    KafkaBroker      string `mapstructure:"kafka_broker"`
    KafkaTopic       string `mapstructure:"kafka_topic"`
    KafkaGroupID     string `mapstructure:"kafka_group_id"`
    APIEndpoint      string `mapstructure:"api_endpoint"`
    DLQTopic         string `mapstructure:"dlq_topic"`
    RequiredFields   []string `mapstructure:"required_fields"`
    SchemaPath       string `mapstructure:"schema_path"`
}

var Config AppConfig

func LoadConfig() {
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath(".")
    viper.AutomaticEnv()

    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("Failed to read config: %v", err)
    }

    if err := viper.Unmarshal(&Config); err != nil {
        log.Fatalf("Unable to decode config: %v", err)
    }
}
