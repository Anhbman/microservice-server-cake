package product

import (
	"github.com/Anhbman/microservice-server-cake/internal/models"
	"github.com/twitchtv/twirp"
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}

func (p *Service) exitsProductForNameAndPrice(name string, price int64) (bool, error) {
	resp := p.db.Where("name = ? AND price = ?", name, price).Find(&models.Product{})
	if err := resp.Error; err != nil {
		return false, err
	}
	return resp.RowsAffected > 0, nil
}

func (p *Service) Create(product *models.Product) (*models.Product, error) {
	isExits, err := p.exitsProductForNameAndPrice(product.Name, product.Price)
	if err != nil {
		return nil, err
	}
	if isExits {
		return nil, twirp.InvalidArgumentError("Product already exists", "Name, Price")
	}
	res := p.db.Create(product)
	if res.Error != nil {
		return nil, res.Error
	}
	return product, nil
}

func (p *Service) GetProductById(id int64) (*models.Product, error) {
	var product models.Product
	res := p.db.First(&product, id)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return nil, twirp.NotFoundError("Product not found")
		}
		return nil, res.Error
	}
	return &product, nil
}

func (p *Service) GetAll() ([]*models.Product, error) {
	var products []*models.Product
	res := p.db.Find(&products)
	if res.Error != nil {
		return nil, res.Error
	}
	return products, nil
}
