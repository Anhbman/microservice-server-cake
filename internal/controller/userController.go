package controller

import (
	"context"

	"github.com/Anhbman/microservice-server-cake/rpc/service"
	"github.com/labstack/gommon/log"
	"github.com/twitchtv/twirp"
)

func (c *Controller) RegisterUser(ctx context.Context, req *service.RegisterUserRequest) (*service.RegisterUserResponse, error) {
	if req.GetName() == "" {
		log.Errorf("Name is required")
		return nil, twirp.InvalidArgumentError("Name is required", "Name")
	}

	if req.GetEmail() == "" {
		log.Errorf("Email is required")
		return nil, twirp.InvalidArgumentError("Email is required", "Email")
	}

	if req.GetPassword() == "" {
		log.Errorf("Password is required")
		return nil, twirp.InvalidArgumentError("Password is required", "Password")
	}
	return c.userService.Register(req)
}

func (c *Controller) LoginUser(ctx context.Context, req *service.LoginUserRequest) (*service.LoginUserResponse, error) {
	if conditions := req.GetEmail() == "" || req.GetPassword() == ""; conditions {
		log.Errorf("Email and password are required")
		return nil, twirp.InvalidArgumentError("Email and password are required", "Email, Password")
	}
	return c.userService.Login(req)
}

func (c *Controller) GetUserById(ctx context.Context, req *service.GetUserByIdRequest) (*service.GetUserByIdResponse, error) {
	if req.GetId() == 0 {
		log.Errorf("ID is required")
		return nil, twirp.InvalidArgumentError("ID is required", "ID")
	}
	return c.userService.GetUserById(req)
}
