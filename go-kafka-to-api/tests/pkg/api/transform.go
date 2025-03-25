package api

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/xeipuuv/gojsonschema"
)

var schemaValidator *gojsonschema.Schema

func LoadSchema(path string) error {
    schemaLoader := gojsonschema.NewReferenceLoader("file://" + path)
    var err error
    schemaValidator, err = gojsonschema.NewSchema(schemaLoader)
    return err
}

type Config struct {
    RenameFields   map[string]string
    RequiredFields []string
    RemoveFields   []string
    StaticTags     map[string]interface{}
}

var DefaultConfig = Config{
    RenameFields: map[string]string{
        "old_key": "new_key",
    },
    RequiredFields: []string{"id"},
    RemoveFields:   []string{"password"},
    StaticTags: map[string]interface{}{
        "source": "redpanda-franz-pipeline",
    },
}

func Transform(input []byte) ([]byte, error) {
    var msg map[string]interface{}
    if err := json.Unmarshal(input, &msg); err != nil {
        return nil, errors.New("invalid JSON input")
    }

    if schemaValidator != nil {
        docLoader := gojsonschema.NewGoLoader(msg)
        result, err := schemaValidator.Validate(docLoader)
        if err != nil {
            return nil, fmt.Errorf("schema validation error: %w", err)
        }
        if !result.Valid() {
            return nil, fmt.Errorf("schema validation failed: %v", result.Errors())
        }
    }

    for oldKey, newKey := range DefaultConfig.RenameFields {
        if val, exists := msg[oldKey]; exists {
            msg[newKey] = val
            delete(msg, oldKey)
        }
    }

    for _, field := range DefaultConfig.RemoveFields {
        delete(msg, field)
    }

    for k, v := range DefaultConfig.StaticTags {
        msg[k] = v
    }

    return json.Marshal(msg)
}
