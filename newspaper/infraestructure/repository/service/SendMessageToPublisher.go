package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func SendEventToPublisher(eventMessage string) error {

	publisherURL := "http://localhost:8081/publish"

	payload := map[string]string{
		"message": eventMessage,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("Error al serializar el mensaje: %v", err)
	}

	resp, err := http.Post(publisherURL, "application/json", bytes.NewReader(payloadBytes))

	if err != nil {
		return fmt.Errorf("Error al enviar el mensaje al publisher: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Error en el publisher: Status code %d", resp.StatusCode)
	}
	return nil
}
