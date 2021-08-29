package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/yogamuris/sohappytocyou/entity"
	"github.com/yogamuris/sohappytocyou/entity/web"
	"github.com/yogamuris/sohappytocyou/helper"
	"github.com/yogamuris/sohappytocyou/repository"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthServiceImpl struct {
	UserRepository repository.UserRepository
	Db             *sql.DB
	Validate       *validator.Validate
}

func NewAuthServie(userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) AuthService {
	return &AuthServiceImpl{
		UserRepository: userRepository,
		Db:             DB,
		Validate:       validate,
	}
}

func (service AuthServiceImpl) Login(ctx context.Context, request web.LoginRequest) (entity.User, error) {
	tx, err := service.Db.Begin()
	if err != nil {
		return entity.User{}, err
	}
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindByUsername(ctx, tx, request.Username)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (service AuthServiceImpl) Register(ctx context.Context, request web.RegisterRequest) (web.AuthResponse, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return web.AuthResponse{}, err
	}

	tx, err := service.Db.Begin()
	if err != nil {
		return web.AuthResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(request.Password), 14)

	user := entity.User{
		Username:  request.Username,
		Email:     request.Email,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
	}

	user, err = service.UserRepository.Save(ctx, tx, user)

	if err != nil {
		return web.AuthResponse{}, err
	}

	return web.AuthResponse{
		Code:    200,
		Message: "Registrasi berhasil",
	}, nil
}

func (service AuthServiceImpl) Verify(ctx context.Context, request web.VerifyRequest) (web.AuthResponse, error) {
	panic("implement me")
}
