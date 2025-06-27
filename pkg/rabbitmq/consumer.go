package rabbitmq

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

type Consumer struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

type EventHandler interface {
	HandlerUserRegister(payload UserRegisterPayload) error
}

func NewConsumer(conn *amqp.Connection) (*Consumer, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open channel: %w", err)
	}

	err = ch.ExchangeDeclare(
		"client_events", // exchange name
		"topic",         // exchange type
		true,            // durable
		false,           // auto-deleted
		false,           // internal
		false,           // no-wait
		nil,             // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare exchange: %w", err)
	}

	queue, err := ch.QueueDeclare(
		"user_registered_queue", // queue name
		true,                    // durable
		false,                   // delete when unused
		false,                   // exclusive
		false,                   // no-wait
		nil,                     // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare queue: %w", err)
	}

	err = ch.QueueBind(
		queue.Name,               // queue name
		"user_events.registered", // routing key
		"client_events",          // exchange
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to bind queue: %w", err)
	}

	return &Consumer{
		conn: conn,
		ch:   ch,
	}, nil
}

func (c *Consumer) StartConsuming() error {
	msgs, err := c.ch.Consume(
		"user_registered_queue", // queue
		"",                      // consumer
		false,                   // auto-ack
		false,                   // exclusive
		false,                   // no-local
		false,                   // no-wait
		nil,                     // args
	)
	if err != nil {
		return fmt.Errorf("failed to register consumer: %w", err)
	}

	go func() {
		for msg := range msgs {
			if err := c.processMessage(msg); err != nil {
				log.Printf("Error processing message: %v", err)
				msg.Nack(false, true) // Reject and requeue
			} else {
				msg.Ack(false) // Acknowledge
			}
		}
	}()

	return nil
}

func (c *Consumer) processMessage(msg amqp.Delivery) error {
	// Parse event
	var event Event
	if err := json.Unmarshal(msg.Body, &event); err != nil {
		return fmt.Errorf("failed to unmarshal event: %w", err)
	}

	// Handle based on event type
	switch event.Type {
	case UserRegisterEvent:
		var payload UserRegisterPayload
		if err := json.Unmarshal(event.Payload, &payload); err != nil {
			return fmt.Errorf("failed to unmarshal payload: %w", err)
		}
		fmt.Printf("Received User Register Event: %+v\n", payload)
		return nil
	default:
		return fmt.Errorf("unknown event type: %s", event.Type)
	}
}

func (c *Consumer) Close() error {
	if c.ch != nil {
		c.ch.Close()
	}
	return nil
}
