//go:build wireinject
// +build wireinject

package main

import (
	"github.com/Anhbman/microservice-server-cake/internal/controller"
	"github.com/Anhbman/microservice-server-cake/internal/server/cake"
	"github.com/Anhbman/microservice-server-cake/internal/storage"
	"github.com/google/wire"
)

func InitializeApp() *controller.ControllerServer {
	wire.Build(storage.InitDB, cake.NewProcessor, controller.NewControllerServer)
	return &controller.ControllerServer{}
}
