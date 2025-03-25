package kafka

import (
    "os"
    "testing"
)

func TestNewSecureKafkaClient_InvalidSecrets(t *testing.T) {
    os.Setenv("KAFKA_USERNAME", "file://does-not-exist-user")
    os.Setenv("KAFKA_PASSWORD", "file://does-not-exist-pass")
    _, err := NewSecureKafkaClient()
    if err == nil {
        t.Error("Expected error for missing secret files")
    }
}
