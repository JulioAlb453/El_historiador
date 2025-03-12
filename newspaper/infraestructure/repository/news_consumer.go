package repository

import (
	"encoding/json"
	"log"
	"main/newspaper/domain"

	"github.com/streadway/amqp"
)

type NewsConsumer struct {
    channel *amqp.Channel
}

func NewNewsConsumer(broker *amqp.Connection) *NewsConsumer {
    ch, err := broker.Channel()
    if err != nil {
        log.Fatalf("Error al abrir el canal: %s", err)
    }

    return &NewsConsumer{
        channel: ch,
    }
}

func (c *NewsConsumer) ConsumeNewsEvents(newsHandler func(news domain.News) error) error {
    q, err := c.channel.QueueDeclare(
        "NoticiaPublicada", 
        true,        
        false,       
        false,       
        false,       
        nil,          
    )
    if err != nil {
        return err
    }

    msgs, err := c.channel.Consume(
        q.Name, 
        "",    
        true,   
        false,  
        false, 
        false,  
        nil,    
    )
    if err != nil {
        return err
    }

    forever := make(chan bool)

    go func() {
        for d := range msgs {
            var news domain.News
            err := json.Unmarshal(d.Body, &news)
            if err != nil {
                log.Printf("Error al decodificar el mensaje: %s", err)
                continue
            }

            err = newsHandler(news)
            if err != nil {
                log.Printf("Error al procesar la noticia: %s", err)
            }
        }
    }()

    log.Printf(" [*] Esperando mensajes. Para salir presiona CTRL+C")
    <-forever

	return nil
}

