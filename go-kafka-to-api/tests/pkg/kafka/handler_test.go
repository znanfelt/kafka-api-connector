package kafka

import (
	"errors"
	"os"
	"testing"

	"github.com/twmb/franz-go/pkg/kgo"
)

func TestHandleRecord_Success(t *testing.T) {
	apiTransform = func(data []byte) ([]byte, error) {
		return []byte(`{"ok":true}`), nil
	}
	apiPost = func(url string, data []byte) error {
		return nil
	}

	os.Setenv("API_URL", "http://example.com")
	r := &kgo.Record{Value: []byte(`{}`)}
	HandleRecord(r, &mockDLQ{})
}

func TestHandleRecord_TransformError(t *testing.T) {
	apiTransform = func(data []byte) ([]byte, error) {
		return nil, errors.New("transform failed")
	}

	r := &kgo.Record{Value: []byte(`bad-data`)}
	HandleRecord(r, &mockDLQ{})
}

func TestHandleRecord_PostError(t *testing.T) {
	apiTransform = func(data []byte) ([]byte, error) {
		return []byte(`ok`), nil
	}
	apiPost = func(url string, data []byte) error {
		return errors.New("api error")
	}

	m := &mockDLQ{}
	r := &kgo.Record{Value: []byte(`ok`)}
	HandleRecord(r, m)

	if !m.called {
		t.Error("Expected DLQ to be called on post error")
	}
}
