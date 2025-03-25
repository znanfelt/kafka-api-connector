package main

import (
    "testing"
)

func TestNewRootCommandHelp(t *testing.T) {
    cmd := NewRootCommand()
    cmd.SetArgs([]string{"--help"})
    if err := cmd.Execute(); err != nil {
        t.Errorf("Expected help command to execute successfully, got error: %v", err)
    }
}
