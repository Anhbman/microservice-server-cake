//go:build wireinject
// +build wireinject

package main

import (
	"github.com/Anhbman/microservice-server-cake/internal/controller"
	"github.com/Anhbman/microservice-server-cake/internal/service/cake"
	"github.com/Anhbman/microservice-server-cake/internal/service/user"
	"github.com/Anhbman/microservice-server-cake/internal/storage"
	"github.com/google/wire"
)

func InitializeApp() *controller.ControllerServer {
	wire.Build(storage.InitDB, cake.NewProcessor, user.NewProcessor, controller.NewControllerServer)
	return &controller.ControllerServer{}
}
