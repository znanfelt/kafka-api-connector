package sender

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSendMessage(t *testing.T) {
	type Message struct {
		Type string `json:"type"`
		Data string `json:"data"`
	}

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer testServer.Close()

	message := Message{Type: "important", Data: "test data"}
	SendMessage(message, testServer.URL)
}
