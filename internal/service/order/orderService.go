package order

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

func (p *Service) CreateOrder(order *models.Order) (*models.Order, error) {
	err := p.db.Create(&order).Error
	if err != nil {
		log.Errorf("Cannot create order: %s", err)
		return nil, twirp.Internal.Errorf("Cannot create order: %w", err)
	}
	return order, nil
}

func (p *Service) GetOrderById(id uint64) (*models.Order, error) {
	var order models.Order
	err := p.db.Where("id = ?", id).First(&order).Error
	if err != nil {
		log.Errorf("Cannot find order: %s", err)
		return nil, twirp.NotFoundError("Order not found")
	}
	return &order, nil
}
