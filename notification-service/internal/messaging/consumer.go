package messaging
/
import (
	"encoding/json"
	"log"
	"time"

	"notification-service/internal/domain"

	amqp "github.com/rabbitmq/amqp091-go"
)

func StartConsumer() {
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

	
	err = ch.ExchangeDeclare(
		"dlx.exchange",
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	
	_, err = ch.QueueDeclare(
		"payment.dlq",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	
	err = ch.QueueBind(
		"payment.dlq",
		"payment.dlq",
		"dlx.exchange",
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	
	args := amqp.Table{
		"x-dead-letter-exchange":    "dlx.exchange",
		"x-dead-letter-routing-key": "payment.dlq",
	}

	_, err = ch.QueueDeclare(
		"payment.completed",
		true,
		false,
		false,
		false,
		args,
	)
	if err != nil {
		log.Fatal(err)
	}

	
	msgs, err := ch.Consume(
		"payment.completed",
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Notification service is listening...")

	
	processed := make(map[string]bool)

	for msg := range msgs {
		var event domain.PaymentEvent

		err := json.Unmarshal(msg.Body, &event)
		if err != nil {
			log.Println("unmarshal error:", err)
			msg.Nack(false, false) 
			continue
		}

	
		if processed[event.EventID] {
			msg.Ack(false)
			continue
		}

		
		if event.Email == "fail@test.com" {
			headers := msg.Headers
			retries := 0

			if headers != nil && headers["x-retry"] != nil {
	switch v := headers["x-retry"].(type) {
	case int32:
		retries = int(v)
	case int64:
		retries = int(v)
	case int:
		retries = v
	default:
		retries = 0
	}
}

			if retries >= 3 {
				log.Println("💀 Sent to DLQ:", event.Email)
				msg.Nack(false, false) 
				continue
			}

			log.Println("🔁 retry:", retries+1)

			newHeaders := amqp.Table{
				"x-retry": retries + 1,
			}

			
			err = ch.Publish(
				"",
				"payment.completed",
				false,
				false,
				amqp.Publishing{
					ContentType: "application/json",
					Body:        msg.Body,
					Headers:     newHeaders,
				},
			)

			if err != nil {
				log.Println("republish error:", err)
			}

			msg.Ack(false)
			continue
		}

		
		log.Println("📧 Email sent to:", event.Email)

		processed[event.EventID] = true

		msg.Ack(false)
	}
}