package controller

import (
	"github.com/Anhbman/microservice-server-cake/internal/service/cake"
	"github.com/Anhbman/microservice-server-cake/internal/service/order"
	orderitem "github.com/Anhbman/microservice-server-cake/internal/service/orderItem"
	"github.com/Anhbman/microservice-server-cake/internal/service/user"
	"github.com/Anhbman/microservice-server-cake/rpc/service"
)

type Controller struct {
	cakeService      *cake.Service
	userService      *user.Service
	orderService     *order.Service
	orderItemService *orderitem.Service
}

var _ service.Service = (*Controller)(nil)

func NewController(cakeService *cake.Service, userService *user.Service, orderService *order.Service, orderItemService *orderitem.Service) *Controller {
	return &Controller{cakeService: cakeService, userService: userService, orderService: orderService, orderItemService: orderItemService}
}
