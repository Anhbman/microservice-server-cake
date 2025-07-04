package controller

import (
	"github.com/Anhbman/microservice-server-cake/internal/service/cake"
	"github.com/Anhbman/microservice-server-cake/internal/service/user"
	"github.com/Anhbman/microservice-server-cake/rpc/service"
)

type Controller struct {
	cakeService *cake.Service
	userService *user.Service
}

var _ service.Service = (*Controller)(nil)

func NewController(cakeService *cake.Service, userService *user.Service) *Controller {
	return &Controller{cakeService: cakeService, userService: userService}
}
