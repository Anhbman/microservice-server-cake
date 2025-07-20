package eventHandler

import (
	"context"

	"github.com/Anhbman/microservice-server-cake/internal/service/cake"
	"github.com/Anhbman/microservice-server-cake/internal/service/user"
	"github.com/Anhbman/microservice-server-cake/rpc/service"
	"github.com/labstack/gommon/log"
	"github.com/twitchtv/twirp"
)

type EventHandler struct {
	cakeService *cake.Service
	userService *user.Service
}

func NewEventHandler(cakeService *cake.Service, userService *user.Service) *EventHandler {
	return &EventHandler{cakeService: cakeService, userService: userService}
}

func (c *EventHandler) RegisterUser(ctx context.Context, req *service.RegisterUserRequest) error {
	if req.GetName() == "" {
		log.Errorf("Name is required")
		return twirp.InvalidArgumentError("Name is required", "Name")
	}

	if req.GetEmail() == "" {
		log.Errorf("Email is required")
		return twirp.InvalidArgumentError("Email is required", "Email")
	}

	if req.GetPassword() == "" {
		log.Errorf("Password is required")
		return twirp.InvalidArgumentError("Password is required", "Password")
	}
	_, err := c.userService.Register(req)
	if err != nil {
		log.Errorf("Failed to register user: %s", err)
		return twirp.InternalErrorWith(err)
	}
	return nil
}
