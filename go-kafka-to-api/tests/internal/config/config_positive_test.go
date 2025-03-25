package config

import (
    "os"
    "testing"
)

func TestLoadSecret_File(t *testing.T) {
    tmpFile := "/tmp/test-secret.txt"
    expected := "s3cr3t"
    os.WriteFile(tmpFile, []byte(expected), 0600)
    defer os.Remove(tmpFile)

    SetSecretDir("/tmp")
    val, err := LoadSecret("file://test-secret.txt")
    if err != nil {
        t.Fatalf("Unexpected error loading secret: %v", err)
    }
    if val != expected {
        t.Errorf("Expected %s, got %s", expected, val)
    }
}
