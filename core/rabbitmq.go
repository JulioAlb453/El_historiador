// core/rabbitmq.go
package core

import (
        "log"
        "os"

        "github.com/streadway/amqp"
)

func ConnectRabbitMQ() *amqp.Connection {
        rabbitMQURL := getenv("RABBITMQ_URL", "amqp://guest:guest@l52.20.122.112:5672/") 
        conn, err := amqp.Dial(rabbitMQURL)
        if err != nil {
                log.Fatalf("Error al conectar a RabbitMQ: %v", err)
        }

        return conn
}

func getenv(key, defaultValue string) string {
        value := os.Getenv(key)
        if value == "" {
                return defaultValue
        }
        return value
}