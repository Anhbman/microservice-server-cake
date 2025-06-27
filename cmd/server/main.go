package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Anhbman/microservice-server-cake/internal/config"
	"github.com/Anhbman/microservice-server-cake/internal/hooks"
	"github.com/Anhbman/microservice-server-cake/pkg/rabbitmq"
	"github.com/Anhbman/microservice-server-cake/rpc/service"
)

func main() {
	// connect to database
	// storage.InitDB()
	// db := storage.GetDB()
	// cakeHandle := cake.NewProcessor(db)
	// handle := controller.NewControllerServer(cakeHandle)

	cfg, err := config.SetupEnv()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	handle := InitializeApp()

	rabbitConfig := rabbitmq.Config{
		Host:     "localhost",
		Port:     5672,
		Username: "guest",
		Password: "guest",
		Vhost:    "/",
	}

	conn, err := rabbitmq.NewConnection(rabbitConfig)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	server := service.NewServiceServer(handle, hooks.LoggingHooks(os.Stderr))

	consumer, err := rabbitmq.NewConsumer(conn)
	if err != nil {
		log.Fatalf("Failed to create RabbitMQ consumer: %v", err)
	}
	defer consumer.Close()

	consumer.StartConsuming()

	log.Printf("Server is running on port %s", cfg.ServerPort)
	err = http.ListenAndServe(":"+cfg.ServerPort, server)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
