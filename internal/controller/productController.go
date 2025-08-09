package controller

import (
	"context"

	"github.com/Anhbman/microservice-server-cake/internal/models"
	"github.com/Anhbman/microservice-server-cake/rpc/service"
	"github.com/labstack/gommon/log"
	"github.com/twitchtv/twirp"
)

func (c *Controller) CreateProduct(ctx context.Context, req *service.CreateProductRequest) (*service.Product, error) {
	if req.GetName() == "" {
		log.Errorf("Product name is required")
		return nil, twirp.InvalidArgumentError("Product name is required", "Name")
	}

	if req.GetPrice() <= 0 {
		log.Errorf("Product price must be greater than zero")
		return nil, twirp.InvalidArgumentError("Product price must be greater than zero", "Price")
	}

	product := &models.Product{
		Name:  req.GetName(),
		Price: req.GetPrice(),
	}

	resp, err := c.productService.Create(product)
	if err != nil {
		log.Errorf("Failed to create product: %v", err)
		return nil, err
	}

	return &service.Product{
		Id:    int64(resp.ID),
		Name:  resp.Name,
		Price: resp.Price,
	}, nil

}
