package main

import (
	"testing"
)

func TestRunConsumerCommand_LoadSchemaError(t *testing.T) {
	err := RunConsumerCommand("file://does-not-exist-schema.json")
	if err == nil {
		t.Error("Expected schema load error")
	}
}
