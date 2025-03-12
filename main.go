package main

import (
	"log"
	"main/newspaper/domain"
	newsInfra "main/newspaper/infraestructure"
	"main/routes"

	"github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error al cargar el archivo .env: %v", err)
    }

    router := routes.SetupRouter()

    deps := newsInfra.Init()

    go func() {
        err := deps.ProcessNewsConsumer.ConsumeNewsEvents(func(news domain.News) error {
            log.Printf("Noticia recibida de RabbitMQ: %+v", news)
            return nil
        })
        if err != nil {
            log.Fatalf("Error al consumir mensajes de RabbitMQ: %s", err)
        }
    }()

    router.Run(":8080")
}