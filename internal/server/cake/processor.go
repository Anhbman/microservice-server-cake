package cake

import (
	"context"

	"github.com/Anhbman/microservice-server-cake/internal/models"
	pb "github.com/Anhbman/microservice-server-cake/rpc/service"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type Processor struct {
	db *gorm.DB
}

func NewProcessor(db *gorm.DB) *Processor {
	return &Processor{db: db}
}

func (p *Processor) Create(ctx context.Context, cake *pb.CreateCakeRequest) (*pb.SearchCake, error) {
	var cakeInsert = models.Cake{
		Name:        cake.Name,
		Price:       cake.Price,
		Description: cake.Description,
		ImageUrl:    cake.ImageUrl,
		UserID:      int64(cake.UserId),
	}
	err := p.db.Create(&cakeInsert).Error
	if err != nil {
		log.Errorf("Cannot create cake: %s", err)
		return nil, err
	}
	return &pb.SearchCake{
		Id:          int64(cakeInsert.ID),
		Name:        cakeInsert.Name,
		Description: cakeInsert.Description,
		Price:       cakeInsert.Price,
		ImageUrl:    cakeInsert.ImageUrl,
	}, nil
}

func (p *Processor) GetCakeById(ctx context.Context, id *pb.GetCakeByIdRequest) (*pb.GetCakeByIdResponse, error) {
	var cake models.Cake
	err := p.db.First(&cake, id.Id).Error
	if err != nil {
		log.Errorf("cake %s", err)
		return nil, err
	}
	return &pb.GetCakeByIdResponse{
		Id:          int64(cake.ID),
		Name:        cake.Name,
		Description: cake.Description,
		Price:       cake.Price,
		ImageUrl:    cake.ImageUrl,
		UserId:      uint64(cake.UserID),
	}, nil
}
