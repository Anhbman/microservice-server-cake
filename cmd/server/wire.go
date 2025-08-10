//go:build wireinject
// +build wireinject

package main

import (
	"github.com/Anhbman/microservice-server-cake/internal/controller"
	"github.com/Anhbman/microservice-server-cake/internal/eventHandler"
	"github.com/Anhbman/microservice-server-cake/internal/service/cake"
	"github.com/Anhbman/microservice-server-cake/internal/service/order"
	"github.com/Anhbman/microservice-server-cake/internal/service/user"
	"github.com/Anhbman/microservice-server-cake/internal/storage"
	"github.com/google/wire"
)

func InitializeApp() *controller.Controller {
	wire.Build(storage.InitDB, cake.NewService, user.NewService, order.NewService, controller.NewController)
	return &controller.Controller{}
}

func InitializeEventHandler() *eventHandler.EventHandler {
	wire.Build(storage.InitDB, cake.NewService, user.NewService, eventHandler.NewEventHandler)
	return &eventHandler.EventHandler{}
}
