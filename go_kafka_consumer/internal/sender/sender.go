package sender

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

func SendMessage(m Message, apiEndpoint string) {
	jsonData, err := json.Marshal(m)
	if err != nil {
		log.Printf("Failed to serialize message: %v", err)
		return
	}

	resp, err := http.Post(apiEndpoint, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("Failed to send message: %v", err)
		return
	}
	defer resp.Body.Close()

	log.Printf("Message sent successfully: %s", resp.Status)
}
