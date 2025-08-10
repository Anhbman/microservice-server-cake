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

func (c *Controller) GetProductById(ctx context.Context, req *service.GetProductByIdRequest) (*service.Product, error) {
	if req.GetId() == 0 {
		log.Errorf("Product ID is required")
		return nil, twirp.InvalidArgumentError("Product ID is required", "ID")
	}

	product, err := c.productService.GetProductById(req.GetId())
	if err != nil {
		log.Errorf("Failed to get product by ID: %v", err)
		return nil, err
	}

	return &service.Product{
		Id:          int64(product.ID),
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
		ImageUrl:    product.ImageURL,
	}, nil
}

func (c *Controller) GetAllProducts(ctx context.Context, req *service.GetAllProductsRequest) (*service.GetAllProductsResponse, error) {
	products, err := c.productService.GetAll()
	if err != nil {
		log.Errorf("Failed to get all products: %v", err)
		return nil, err
	}

	productResponses := make([]*service.Product, len(products))
	for _, product := range products {
		productResponses = append(productResponses, &service.Product{
			Id:          int64(product.ID),
			Name:        product.Name,
			Price:       product.Price,
			Description: product.Description,
			ImageUrl:    product.ImageURL,
		})
	}

	return &service.GetAllProductsResponse{Products: productResponses}, nil
}
