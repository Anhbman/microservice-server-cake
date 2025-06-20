package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Anhbman/microservice-server-cake/internal/config"
	"github.com/Anhbman/microservice-server-cake/internal/hooks"
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
	server := service.NewServiceServer(handle, hooks.LoggingHooks(os.Stderr))

	log.Printf("Server is running on port %s", cfg.ServerPort)
	err = http.ListenAndServe(":"+cfg.ServerPort, server)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
