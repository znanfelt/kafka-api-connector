package config

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
)

var SecretDir = "/etc/kafka-api/secrets" // default path

// SetSecretDir allows dynamic configuration of the base directory for file:// secrets
func SetSecretDir(dir string) {
    SecretDir = dir
}

// LoadSecret loads a secret string or resolves file://path relative to the SecretDir
func LoadSecret(valueOrFile string) (string, error) {
    if strings.HasPrefix(valueOrFile, "file://") {
        relativePath := strings.TrimPrefix(valueOrFile, "file://")
        fullPath := filepath.Join(SecretDir, relativePath)
        data, err := os.ReadFile(fullPath)
        if err != nil {
            return "", fmt.Errorf("failed to read secret file %s: %w", fullPath, err)
        }
        return strings.TrimSpace(string(data)), nil
    }
    return valueOrFile, nil
}
