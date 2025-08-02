package main

import (
	"log"
	"os"
	"time"

	"github.com/Lucashcr/openservicebroker/logger/internal/domain/services"
	"github.com/Lucashcr/openservicebroker/logger/internal/infra/rabbitmq"
)

func main() {
	addr := os.Getenv("RABBITMQ_URL")
	queueName := os.Getenv("RABBITMQ_QUEUE_NAME")
	if len(queueName) == 0 {
		log.Fatalln("Missing RabbitMQ credentials")
	}

	client := rabbitmq.MakeClient(queueName, addr)
	<-time.After(time.Second)

	deliveries, chClosedCh, err := client.MakeConsumer()
	if err != nil {
		log.Println("Could not start consuming: ", err)
	}

	lokiUrl := os.Getenv("LOKI_URL")
	if len(lokiUrl) == 0 {
		log.Fatalln("Missing LOKI_URL environment variable")
	}

	log.Println("Starting to consume messages from RabbitMQ...")

	for {
		select {
		case amqErr := <-chClosedCh:
			log.Println("AMQP Channel closed due to: ", amqErr)

			deliveries, chClosedCh, err = client.MakeConsumer()
			if err != nil {
				log.Println("Error trying to consume, will try again")
				continue
			}

		case delivery := <-deliveries:
			message, err := services.DecodeMessage(delivery.Body)
			if err != nil {
				log.Println("Error to decode message: ", err)
				err = delivery.Nack(false, true)
				if err != nil {
					log.Println("Error non acknowledging message: ", err)
				}
				continue
			}

			err = services.ProcessMessage(message, lokiUrl)
			if err != nil {
				log.Println("Error to process message: ", err)
				err = delivery.Nack(false, false)
				if err != nil {
					log.Println("Error non acknowledging message: ", err)
				}
				continue
			}

			err = delivery.Ack(false)
			if err != nil {
				log.Println("Error acknowledging message:", err)
			}
		}
	}
}
