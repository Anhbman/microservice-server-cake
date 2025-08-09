package product

import (
	"github.com/Anhbman/microservice-server-cake/internal/models"
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}

func (p *Service) Create(product *models.Product) (*models.Product, error) {
	res := p.db.Create(product)
	if res.Error != nil {
		return nil, res.Error
	}
	return product, nil
}
