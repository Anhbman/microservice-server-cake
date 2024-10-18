// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/Anhbman/microservice-server-cake/internal/controller"
	"github.com/Anhbman/microservice-server-cake/internal/server/cake"
	"github.com/Anhbman/microservice-server-cake/internal/server/user"
	"github.com/Anhbman/microservice-server-cake/internal/storage"
)

// Injectors from wire.go:

func InitializeApp() *controller.ControllerServer {
	db := storage.InitDB()
	processor := cake.NewProcessor(db)
	userProcessor := user.NewProcessor(db)
	controllerServer := controller.NewControllerServer(processor, userProcessor)
	return controllerServer
}
