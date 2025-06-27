package rabbitmq

import (
	"fmt"

	"github.com/streadway/amqp"
)

type Config struct {
	Host     string
	Port     int
	Username string
	Password string
	Vhost    string
}

func NewConnection(cfg Config) (*amqp.Connection, error) {
	dsn := fmt.Sprintf("amqp://%s:%s@%s:%d/%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Vhost,
	)

	conn, err := amqp.Dial(dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}
	fmt.Println("Connected to RabbitMQ successfully")
	return conn, nil
}
