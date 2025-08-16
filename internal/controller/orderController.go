package controller

import (
	"context"

	"github.com/Anhbman/microservice-server-cake/internal/models"
	"github.com/Anhbman/microservice-server-cake/rpc/service"
	"github.com/labstack/gommon/log"
	"github.com/twitchtv/twirp"
)

func (c *Controller) CreateOrder(ctx context.Context, req *service.CreateOrderRequest) (*service.Order, error) {
	if req.GetUserId() == 0 {
		log.Errorf("User ID is required")
		return nil, twirp.InvalidArgumentError("User ID is required", "UserId")
	}
	if len(req.GetItems()) == 0 {
		log.Errorf("Order items are required")
		return nil, twirp.InvalidArgumentError("Order items are required", "OrderItems")
	}

	order := &models.Order{
		UserID: int64(req.GetUserId()),
	}

	orderItems := make([]*models.OrderItem, len(req.GetItems()))
	for i, item := range req.GetItems() {
		orderItems[i] = &models.OrderItem{
			CakeID:   item.GetCakeId(),
			Quantity: item.GetQuantity(),
		}
	}

	order, err := c.orderService.CreateOrderWithItems(order, orderItems)
	if err != nil {
		log.Errorf("Cannot create order: %+v", err)
		return nil, twirp.Internal.Errorf("Cannot create order: %w", err)
	}

	OIs := make([]*service.OrderItem, len(orderItems))
	for i, item := range orderItems {
		OIs[i] = &service.OrderItem{
			Id:       uint64(item.ID),
			CakeId:   item.CakeID,
			Quantity: item.Quantity,
		}
	}

	return &service.Order{
		Id:     uint64(order.ID),
		UserId: uint64(order.UserID),
		Items:  OIs,
	}, nil
}

func (c *Controller) GetOrderById(ctx context.Context, req *service.GetOrderByIdRequest) (*service.Order, error) {
	if req.GetId() == 0 {
		log.Errorf("Order ID is required")
		return nil, twirp.InvalidArgumentError("Order ID is required", "OrderId")
	}

	order, err := c.orderService.GetOrderById(req.GetId())
	if err != nil {
		log.Errorf("Cannot find order: %s", err)
		return nil, twirp.Internal.Errorf("Cannot find order: %w", err)
	}

	orderItems, err := c.orderItemService.GetOrderItemsByOrderId(req.GetId())
	if err != nil {
		log.Errorf("Cannot find order items: %s", err)
		return nil, twirp.Internal.Errorf("Cannot find order items: %w", err)
	}

	OIs := make([]*service.OrderItem, len(orderItems))
	for i, item := range orderItems {
		OIs[i] = &service.OrderItem{
			Id:       uint64(item.ID),
			CakeId:   item.CakeID,
			Quantity: item.Quantity,
		}
	}
	return &service.Order{
		Id:     uint64(order.ID),
		UserId: uint64(order.UserID),
		Items:  OIs,
	}, nil
}
