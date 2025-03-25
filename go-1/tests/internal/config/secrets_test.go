package config

import (
    "os"
    "testing"
)

func TestLoadSecretFromValue(t *testing.T) {
    val, err := LoadSecret("my-secret")
    if err != nil || val != "my-secret" {
        t.Errorf("Expected 'my-secret', got %s, err: %v", val, err)
    }
}

func TestLoadSecretFromFile(t *testing.T) {
    os.MkdirAll("./test-secrets", 0755)
    os.WriteFile("./test-secrets/test.txt", []byte("supersecret"), 0644)
    defer os.RemoveAll("./test-secrets")

    SetSecretDir("./test-secrets")
    val, err := LoadSecret("file://test.txt")
    if err != nil || val != "supersecret" {
        t.Errorf("Expected 'supersecret', got %s, err: %v", val, err)
    }
}
