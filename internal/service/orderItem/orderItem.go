package orderitem

import (
	"github.com/Anhbman/microservice-server-cake/internal/models"
	"github.com/labstack/gommon/log"
	"github.com/twitchtv/twirp"
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}

func (s *Service) CreateOrderItem(orderItem []*models.OrderItem) ([]*models.OrderItem, error) {
	err := s.db.Create(&orderItem).Error
	if err != nil {
		log.Errorf("Cannot create order item: %s", err)
		return nil, twirp.Internal.Errorf("Cannot create order item: %w", err)
	}
	return orderItem, nil
}

func (s *Service) GetOrderItemsByOrderId(id uint64) ([]*models.OrderItem, error) {
	var orderItem []*models.OrderItem
	err := s.db.Where("order_id = ?", id).Find(&orderItem).Error
	if err != nil {
		log.Errorf("Cannot find order item: %s", err)
		return nil, err
	}
	return orderItem, nil
}
