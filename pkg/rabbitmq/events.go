package rabbitmq

import "encoding/json"

const (
	// Exchanges
	ClientEventsExchange = "client_events"
	DeadLetterExchange   = "dead_letter_events" // This will be the final DLX for messages that fail all retries
	RetryExchange        = "retry_events"       // This exchange will route messages to the retry queue

	// Queues
	UserRegisteredQueue           = "user_registered"
	UserRegisteredDeadLetterQueue = "user_registered_dead_letter" // Final destination for failed messages
	UserRegisteredRetryQueue      = "user_registered_retry"       // Queue for messages waiting for a retry attempt

	// Exchange Types
	TopicExchange = "topic"

	// Routing Keys
	UserEventsRoutingKey = "user.*"

	// Retry Delays (in milliseconds)
	RetryDelay = 10000 // 10 seconds

	// Headers
	HeaderRetryCount = "x-retry-count"

	// Max Retries
	MaxRetries = 3
)

// Events
const (
	UserRegisteredEvent = "user.registered"
)

type Event struct {
	Type    string          `json:"type"`
	ID      string          `json:"id"`
	Source  string          `json:"source"`
	Payload json.RawMessage `json:"payload"`
}