package service

import (
    "bytes"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os"
)

func SendEventToPublisher(payload map[string]interface{}) error {
    publisherURL := getenv("PUBLISHER_URL", "http://localhost:8081/publish") // Valor predeterminado

    log.Println("Enviando payload al publisher en URL:", publisherURL)

    payloadBytes, err := json.Marshal(payload)
    if err != nil {
        log.Printf("Error al serializar el payload: %v", err)
        return fmt.Errorf("Error al serializar el payload: %v", err)
    }

    log.Println("Payload:", string(payloadBytes))

    resp, err := http.Post(publisherURL, "application/json", bytes.NewReader(payloadBytes))
    if err != nil {
        log.Printf("Error al enviar el payload al publisher: %v", err)
        return fmt.Errorf("Error al enviar el payload al publisher: %v", err)
    }

    defer resp.Body.Close()

    log.Println("Respuesta del publisher:", resp.StatusCode)

    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("Error en el publisher: Status code %d", resp.StatusCode)
    }

    log.Println("Payload enviado al publisher exitosamente")

    return nil
}

func getenv(key, defaultValue string) string {
    value := os.Getenv(key)
    if value == "" {
        return defaultValue
    }
    return value
}