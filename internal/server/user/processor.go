package user

import (
	"context"
	"fmt"
	"strings"

	"github.com/Anhbman/microservice-server-cake/internal/models"
	"github.com/Anhbman/microservice-server-cake/rpc/service"
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

func (p *Processor) Register(ctx context.Context, user *service.RegisterUserRequest) (*service.RegisterUserResponse, error) {
	var u models.User
	fmt.Println("12: ", user)
	if user.GetName() == "" {
		log.Errorf("Name is required")
		return nil, twirp.InvalidArgumentError("Name is required", "Name")
	}

	if user.GetEmail() == "" {
		log.Errorf("Email is required")
		return nil, twirp.InvalidArgumentError("Email is required", "Email")
	}

	if user.GetPassword() == "" {
		log.Errorf("Password is required")
		return nil, twirp.InvalidArgumentError("Password is required", "Password")
	}
	u.Name = user.Name
	u.Email = user.Email
	hashedPassword, err := u.HashPassword(user.Password)
	if err != nil {
		log.Errorf("Cannot hash password: %s", err)
		return nil, twirp.Internal.Errorf("Cannot hash password: %w", err)
	}
	u.Password = hashedPassword
	err = p.db.Create(&u).Error
	fmt.Println("123: ", err)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			log.Errorf("Email is already taken")
			return nil, twirp.InvalidArgumentError("Email is already taken", "Email")
		}
		log.Errorf("Cannot create user: %s", err)
		return nil, twirp.Internal.Errorf("Cannot create user: %w", err)
	}

	return &service.RegisterUserResponse{
		Id:    uint64(u.ID),
		Name:  u.Name,
		Email: u.Email,
	}, nil
}
