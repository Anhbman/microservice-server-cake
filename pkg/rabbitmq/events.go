package rabbitmq

import (
	"encoding/json"
	"time"
)

type Event struct {
	ID        string          `json:"id"`
	Type      string          `json:"type"`
	Payload   json.RawMessage `json:"payload"`
	Timestamp time.Time       `json:"timestamp"`
	Source    string          `json:"source"`
}

type UserRegisterPayload struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Event Types
const (
	UserRegisteredEvent = "user.registered"
	CakeCreatedEvent    = "cake.created"
	CakeUpdatedEvent    = "cake.updated"
)

// Exchange Names
const (
	ClientEventsExchange = "client_events"
	CakeEventsExchange   = "cake_events"
	OrderEventsExchange  = "order_events"
)

// Queue Names
const (
	UserRegisteredQueue = "user_registered_queue"
	CakeCreatedQueue    = "cake_created_queue"
	OrderProcessQueue   = "order_process_queue"
)

// Routing Keys
const (
	UserEventsRoutingKey  = "user_events.*"
	CakeEventsRoutingKey  = "cake_events.*"
	OrderEventsRoutingKey = "order_events.*"
)

// Exchange Types
const (
	TopicExchange  = "topic"
	DirectExchange = "direct"
	FanoutExchange = "fanout"
)
