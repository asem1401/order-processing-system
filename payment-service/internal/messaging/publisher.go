package messaging

import (
	"encoding/json"
	"log"
	"time"

	"payment-service/internal/domain"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Publisher struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}


func NewPublisher() *Publisher {
	var conn *amqp.Connection
	var err error

	for i := 0; i < 10; i++ {
		conn, err = amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
		if err == nil {
			break
		}
		log.Println("Waiting for RabbitMQ...")
		time.Sleep(3 * time.Second)
	}

	if err != nil {
		log.Fatal(err)
	}

	
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	return &Publisher{
		conn: conn,
		ch:   ch,
	}
}


func (p *Publisher) Publish(event domain.PaymentEvent) {
	body, err := json.Marshal(event)
	if err != nil {
		log.Println(err)
		return
	}

	err = p.ch.Publish(
		"",
		"payment.completed", 
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         body,
			DeliveryMode: amqp.Persistent,
		},
	)

	if err != nil {
		log.Println("publish error:", err)
		return
	}

	log.Println("Event published:", event.EventID)
}