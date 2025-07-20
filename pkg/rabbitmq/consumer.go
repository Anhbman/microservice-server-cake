package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/Anhbman/microservice-server-cake/internal/eventHandler"
	"github.com/Anhbman/microservice-server-cake/rpc/service"
	"github.com/streadway/amqp"
)

type Consumer struct {
	conn         *amqp.Connection
	ch           *amqp.Channel
	eventHandler *eventHandler.EventHandler
}

type ConsumerConfig struct {
	Exchange    string
	Queue       string
	RoutingKeys []string
	Durable     bool
	AutoDelete  bool
	Exclusive   bool
}

// type EventHandler interface {
// 	Handle(ctx context.Context, event Event) error
// 	HandleUserRegistered(ctx context.Context, payload UserRegisterPayload) error
// }

func NewConsumer(conn *amqp.Connection, eventHandler *eventHandler.EventHandler) (*Consumer, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open channel: %w", err)
	}

	err = ch.ExchangeDeclare(
		ClientEventsExchange, // exchange name
		TopicExchange,        // exchange type
		true,                 // durable
		false,                // auto-deleted
		false,                // internal
		false,                // no-wait
		nil,                  // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare exchange: %w", err)
	}

	queue, err := ch.QueueDeclare(
		UserRegisteredQueue, // queue name
		true,                // durable
		false,               // delete when unused
		false,               // exclusive
		false,               // no-wait
		nil,                 // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare queue: %w", err)
	}

	err = ch.QueueBind(
		queue.Name,           // queue name
		UserEventsRoutingKey, // routing key
		ClientEventsExchange, // exchange
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to bind queue: %w", err)
	}

	return &Consumer{
		conn:         conn,
		ch:           ch,
		eventHandler: eventHandler,
	}, nil
}

func (c *Consumer) StartConsuming() error {
	msgs, err := c.ch.Consume(
		UserRegisteredQueue, // queue
		"",                  // consumer
		false,               // auto-ack
		false,               // exclusive
		false,               // no-local
		false,               // no-wait
		nil,                 // args
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

	if err := c.handleMessage(context.Background(), msg); err != nil {
		log.Printf("Error processing message: %v", err)
	}
	return nil
}

func (c *Consumer) handleMessage(ctx context.Context, msg amqp.Delivery) error {
	var event Event
	if err := json.Unmarshal(msg.Body, &event); err != nil {
		return fmt.Errorf("failed to unmarshal event: %w", err)
	}

	log.Printf("Received event: Type=%s, ID=%s, Source=%s", event.Type, event.ID, event.Source)

	switch event.Type {
	case UserRegisteredEvent:
		var payload service.RegisterUserRequest
		if err := json.Unmarshal(event.Payload, &payload); err != nil {
			return fmt.Errorf("failed to unmarshal UserRegisterPayload: %w", err)
		}
		return c.eventHandler.RegisterUser(ctx, &payload)
	default:
		// Use generic handler for unknown event types
		return fmt.Errorf("unknown event type: %s", event.Type)
	}
}

func (c *Consumer) Close() error {
	if c.ch != nil {
		c.ch.Close()
	}
	return nil
}
