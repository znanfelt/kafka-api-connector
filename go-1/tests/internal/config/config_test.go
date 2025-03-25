package config

import (
	"testing"
)

func TestLoadSecret_MissingFile(t *testing.T) {
	path := "file://nonexistent.txt"
	_, err := LoadSecret(path)
	if err == nil {
		t.Error("Expected error for missing secret file")
	}
}

func TestLoadSecret_Inline(t *testing.T) {
	val, err := LoadSecret("plainvalue")
	if err != nil {
		t.Errorf("Expected inline value to succeed, got error: %v", err)
	}
	if val != "plainvalue" {
		t.Errorf("Expected plainvalue, got %s", val)
	}
}
