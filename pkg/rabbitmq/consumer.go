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

	// Declare the main exchange where events are originally published.
	err = ch.ExchangeDeclare(ClientEventsExchange, TopicExchange, true, false, false, false, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to declare main exchange: %w", err)
	}

	// Declare the retry exchange.
	err = ch.ExchangeDeclare(RetryExchange, TopicExchange, true, false, false, false, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to declare retry exchange: %w", err)
	}

	// Declare the dead-letter exchange for messages that fail all retries.
	err = ch.ExchangeDeclare(DeadLetterExchange, TopicExchange, true, false, false, false, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to declare dead-letter exchange: %w", err)
	}

	// Declare the main queue. Messages that fail processing will be sent to the RetryExchange.
	mainQueueArgs := amqp.Table{
		"x-dead-letter-exchange": RetryExchange,
	}
	mainQueue, err := ch.QueueDeclare(UserRegisteredQueue, true, false, false, false, mainQueueArgs)
	if err != nil {
		return nil, fmt.Errorf("failed to declare main queue: %w", err)
	}

	// Bind the main queue to the main exchange.
	err = ch.QueueBind(mainQueue.Name, UserEventsRoutingKey, ClientEventsExchange, false, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to bind main queue: %w", err)
	}

	// Declare the retry queue. It has a TTL and will send expired messages back to the main exchange.
	retryQueueArgs := amqp.Table{
		"x-dead-letter-exchange":    ClientEventsExchange,
		"x-message-ttl":             RetryDelay,
	}
	retryQueue, err := ch.QueueDeclare(UserRegisteredRetryQueue, true, false, false, false, retryQueueArgs)
	if err != nil {
		return nil, fmt.Errorf("failed to declare retry queue: %w", err)
	}

	// Bind the retry queue to the retry exchange.
	err = ch.QueueBind(retryQueue.Name, "#", RetryExchange, false, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to bind retry queue: %w", err)
	}

	// Declare the final dead-letter queue for messages that have exhausted all retries.
	deadLetterQueue, err := ch.QueueDeclare(UserRegisteredDeadLetterQueue, true, false, false, false, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to declare dead-letter queue: %w", err)
	}

	// Bind the dead-letter queue to the dead-letter exchange.
	err = ch.QueueBind(deadLetterQueue.Name, "#", DeadLetterExchange, false, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to bind dead-letter queue: %w", err)
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
				c.handleFailure(msg, err)
			} else {
				msg.Ack(false) // Acknowledge successful processing
			}
		}
	}()

	return nil
}

func (c *Consumer) processMessage(msg amqp.Delivery) error {
	// The actual business logic is in handleMessage.
	// This function just orchestrates the call and error handling.
	return c.handleMessage(context.Background(), msg)
}

func (c *Consumer) handleFailure(msg amqp.Delivery, err error) {
	log.Printf("Error processing message: %v", err)

	retryCount, _ := msg.Headers[HeaderRetryCount].(int32)

	if retryCount >= MaxRetries {
		log.Printf("Max retries reached for message. Sending to dead-letter exchange. Message ID: %s", msg.MessageId)
		// Publish to the final dead-letter exchange, preserving the original routing key
		err := c.ch.Publish(DeadLetterExchange, msg.RoutingKey, false, false, amqp.Publishing{
			ContentType: msg.ContentType,
			Body:        msg.Body,
		})
		if err != nil {
			log.Printf("Failed to publish to dead-letter exchange: %v", err)
		}
		msg.Ack(false) // Acknowledge the original message to remove it from the queue
		return
	}

	// Increment retry count and publish to the retry exchange, preserving the original routing key
	log.Printf("Retrying message. Count: %d. Message ID: %s", retryCount+1, msg.MessageId)
	retryCount++
	headers := amqp.Table{
		HeaderRetryCount: retryCount,
	}

	err = c.ch.Publish(RetryExchange, msg.RoutingKey, false, false, amqp.Publishing{
		ContentType: msg.ContentType,
		Body:        msg.Body,
		Headers:     headers,
	})

	if err != nil {
		log.Printf("Failed to publish to retry exchange: %v", err)
		// If we can't even publish to the retry exchange, we should probably Nack to keep it in the main queue for a bit.
		// This is a fallback and might need more sophisticated handling.
		msg.Nack(false, true)
		return
	}

	msg.Ack(false) // Acknowledge the original message to remove it from the main queue
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
