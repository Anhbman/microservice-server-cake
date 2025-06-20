package cake

import (
	"context"
	"fmt"
	"strings"

	"github.com/Anhbman/microservice-server-cake/internal/models"
	"github.com/Anhbman/microservice-server-cake/internal/utils"
	pb "github.com/Anhbman/microservice-server-cake/rpc/service"
	"github.com/labstack/gommon/log"
	"github.com/twitchtv/twirp"
	"gorm.io/gorm"
)

type Processor struct {
	db *gorm.DB
}

func NewProcessor(db *gorm.DB) *Processor {
	return &Processor{db: db}
}

func (p *Processor) Create(ctx context.Context, cake *pb.CreateCakeRequest) (*pb.Cake, error) {
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
		return nil, twirp.Internal.Errorf("Cannot create cake: %w", err)
	}
	return &pb.Cake{
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
		return nil, twirp.NotFoundError("Cannot find cake: " + fmt.Sprint(id.Id))
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

func (p *Processor) SearchCake(ctx context.Context, search *pb.SearchCakeRequest) (*pb.SearchCakeResponse, error) {

	conditions := make(map[string]interface{})

	if search.UserId != 0 {
		conditions["user_id"] = search.UserId
	}

	if search.Page < 1 {
		search.Page = 1
	}

	if search.PageSize < 0 {
		search.PageSize = 10
	}

	var cakes []models.Cake
	paginate := utils.NewPaginate(int(search.PageSize), int(search.Page))
	err := p.db.Where("name LIKE ?", "%"+strings.ToLower(search.Name)+"%").Where(conditions).Scopes(paginate.PaginatedResult).Find(&cakes).Error
	if err != nil {
		log.Errorf("Cannot search cake: %s", err)
		return nil, twirp.Internal.Errorf("Cannot search cake: %w", err)
	}

	resp := make([]*pb.Cake, len(cakes))

	for i, cake := range cakes {
		resp[i] = &pb.Cake{
			Id:          int64(cake.ID),
			Name:        cake.Name,
			Description: cake.Description,
			Price:       cake.Price,
			ImageUrl:    cake.ImageUrl,
			UserId:      uint64(cake.UserID),
		}
	}
	return &pb.SearchCakeResponse{
		Cakes: resp,
	}, nil
}

func (p *Processor) UpdateCake(ctx context.Context, cake *pb.Cake) (*pb.Cake, error) {
	cakeUpdate := models.Cake{}
	err := p.db.First(&cakeUpdate, cake.Id).Error
	if err != nil {
		log.Errorf("Cannot find cake: %s", err)
		return nil, twirp.NotFoundError("Cannot find cake: " + fmt.Sprint(cake.Id))
	}
	cakeUpdate.Name = cake.Name
	cakeUpdate.Description = cake.Description
	cakeUpdate.Price = cake.Price
	cakeUpdate.ImageUrl = cake.ImageUrl
	err = p.db.Save(&cakeUpdate).Error
	if err != nil {
		log.Errorf("Cannot update cake: %s", err)
		return nil, twirp.Internal.Errorf("Cannot update cake: %w", err)
	}
	return &pb.Cake{
		Id:          int64(cakeUpdate.ID),
		Name:        cakeUpdate.Name,
		Description: cakeUpdate.Description,
		Price:       cakeUpdate.Price,
		ImageUrl:    cakeUpdate.ImageUrl,
		UserId:      uint64(cakeUpdate.UserID),
	}, nil
}
