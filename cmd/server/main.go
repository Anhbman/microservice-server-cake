package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Anhbman/microservice-server-cake/internal/controller"
	"github.com/Anhbman/microservice-server-cake/internal/hooks"
	"github.com/Anhbman/microservice-server-cake/internal/server/cake"
	"github.com/Anhbman/microservice-server-cake/internal/storage"
	"github.com/Anhbman/microservice-server-cake/rpc/service"
)

func main() {
	// connect to database
	storage.InitDB()
	db := storage.GetDB()
	cakeHandle := cake.NewProcessor(db)
	handle := controller.NewServiceServer(cakeHandle)
	server := service.NewServiceServer(handle, hooks.LoggingHooks(os.Stderr))

	log.Printf("Server is running on port %s", ":8081")
	err := http.ListenAndServe(":8081", server)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
