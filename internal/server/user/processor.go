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
		User: &service.User{
			Id:    uint64(u.ID),
			Name:  u.Name,
			Email: u.Email,
		},
	}, nil
}

func (p *Processor) Login(ctx context.Context, user *service.LoginUserRequest) (*service.LoginUserResponse, error) {
	if conditions := user.GetEmail() == "" || user.GetPassword() == ""; conditions {
		log.Errorf("Email and password are required")
		return nil, twirp.InvalidArgumentError("Email and password are required", "Email, Password")
	}

	var u models.User
	err := p.db.Where("email = ?", user.Email).First(&u).Error
	if err != nil {
		log.Errorf("Cannot find user: %s", err)
		return nil, twirp.InvalidArgumentError("Invalid email or password", "Email orPassword")
	}
	if !u.CheckPassword(user.Password) {
		log.Errorf("Invalid password")
		return nil, twirp.InvalidArgumentError("Invalid email or password", "Email orPassword")
	}

	return &service.LoginUserResponse{
		User: &service.User{
			Id:    uint64(u.ID),
			Name:  u.Name,
			Email: u.Email,
		},
	}, nil
}

func (p *Processor) GetUserById(ctx context.Context, req *service.GetUserByIdRequest) (*service.GetUserByIdResponse, error) {
	if req.GetId() == 0 {
		log.Errorf("ID is required")
		return nil, twirp.InvalidArgumentError("ID is required", "ID")
	}

	var u models.User
	err := p.db.Where("id = ?", req.GetId()).First(&u).Error
	if err != nil {
		log.Errorf("Cannot find user: %s", err)
		return nil, twirp.NotFoundError("User not found")
	}

	return &service.GetUserByIdResponse{
		User: &service.User{
			Id:    uint64(u.ID),
			Name:  u.Name,
			Email: u.Email,
		},
	}, nil
}
