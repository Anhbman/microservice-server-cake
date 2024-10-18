package controller

import (
	"context"

	"github.com/Anhbman/microservice-server-cake/internal/server/cake"
	"github.com/Anhbman/microservice-server-cake/internal/server/user"
	"github.com/Anhbman/microservice-server-cake/rpc/service"
)

type ControllerServer struct {
	cakeProcessor *cake.Processor
	userProcessor *user.Processor
}

var _ service.Service = (*ControllerServer)(nil)

func NewControllerServer(cakeProcessor *cake.Processor, userProcessor *user.Processor) *ControllerServer {
	return &ControllerServer{cakeProcessor: cakeProcessor, userProcessor: userProcessor}
}

func (c *ControllerServer) CreateCake(ctx context.Context, req *service.CreateCakeRequest) (*service.Cake, error) {
	return c.cakeProcessor.Create(ctx, req)
}

func (c *ControllerServer) GetCakeById(ctx context.Context, req *service.GetCakeByIdRequest) (*service.GetCakeByIdResponse, error) {
	return c.cakeProcessor.GetCakeById(ctx, req)
}

func (c *ControllerServer) SearchCake(ctx context.Context, req *service.SearchCakeRequest) (*service.SearchCakeResponse, error) {
	return c.cakeProcessor.SearchCake(ctx, req)
}

func (c *ControllerServer) UpdateCake(ctx context.Context, req *service.Cake) (*service.Cake, error) {
	return c.cakeProcessor.UpdateCake(ctx, req)
}

func (c *ControllerServer) RegisterUser(ctx context.Context, req *service.RegisterUserRequest) (*service.RegisterUserResponse, error) {
	return c.userProcessor.Register(ctx, req)
}

func (c *ControllerServer) LoginUser(ctx context.Context, req *service.LoginUserRequest) (*service.LoginUserResponse, error) {
	return c.userProcessor.Login(ctx, req)
}
