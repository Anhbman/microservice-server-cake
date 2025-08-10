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

func (c *Controller) GetAllCakes(ctx context.Context, req *service.GetAllCakesRequest) (*service.GetAllCakesResponse, error) {
	cakes, err := c.cakeService.GetAllCakes()
	if err != nil {
		return nil, err
	}

	response := &service.GetAllCakesResponse{
		Cakes: make([]*service.Cake, len(cakes)),
	}

	for i, cake := range cakes {
		response.Cakes[i] = &service.Cake{
			Id:          int64(cake.ID),
			Name:        cake.Name,
			Description: cake.Description,
			Price:       cake.Price,
			ImageUrl:    cake.ImageUrl,
			UserId:      uint64(cake.UserID),
		}
	}

	return &service.GetAllCakesResponse{
		Cakes: response.Cakes,
	}, nil
}
