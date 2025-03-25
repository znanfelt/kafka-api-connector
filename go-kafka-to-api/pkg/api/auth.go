package api

import (
    "os"
    "strings"
)

var authHeader string

func InitAPIAuth() error {
    token := os.Getenv("API_AUTH_TOKEN")
    if strings.HasPrefix(token, "file://") {
        content, err := os.ReadFile(strings.TrimPrefix(token, "file://"))
        if err != nil {
            return err
        }
        token = string(content)
    }

    if token != "" {
        authHeader = "Bearer " + strings.TrimSpace(token)
    }
    return nil
}

func getAuthHeader() string {
    return authHeader
}
