package controller

import (
	"context"

	"github.com/Anhbman/microservice-server-cake/rpc/service"
)

func (c *Controller) CreateCake(ctx context.Context, req *service.CreateCakeRequest) (*service.Cake, error) {
	return c.cakeService.Create(req)
}

func (c *Controller) GetCakeById(ctx context.Context, req *service.GetCakeByIdRequest) (*service.GetCakeByIdResponse, error) {
	return c.cakeService.GetCakeById(req)
}

func (c *Controller) SearchCake(ctx context.Context, req *service.SearchCakeRequest) (*service.SearchCakeResponse, error) {
	return c.cakeService.SearchCake(req)
}

func (c *Controller) UpdateCake(ctx context.Context, req *service.Cake) (*service.Cake, error) {
	return c.cakeService.UpdateCake(req)
}
