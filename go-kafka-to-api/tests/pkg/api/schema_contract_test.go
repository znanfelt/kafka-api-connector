package api

import (
    "encoding/json"
    "testing"
)

func TestSchemaContractSnapshot(t *testing.T) {
    sample := map[string]interface{}{
        "id":      "abc123",
        "old_key": "legacy",
    }

    doc, _ := json.Marshal(sample)
    transformed, err := Transform(doc)
    if err != nil {
        t.Fatalf("Schema contract failed: %v", err)
    }

    var msg map[string]interface{}
    json.Unmarshal(transformed, &msg)

    expected := "redpanda-franz-pipeline"
    if msg["source"] != expected {
        t.Errorf("Expected source '%s', got '%s'", expected, msg["source"])
    }
}
