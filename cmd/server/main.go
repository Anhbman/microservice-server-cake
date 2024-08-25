package main

import (
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
	http.ListenAndServe(":8081", server)
}
