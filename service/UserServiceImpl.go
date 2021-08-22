package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/yogamuris/sohappytocyou/entity"
	"github.com/yogamuris/sohappytocyou/entity/web"
	"github.com/yogamuris/sohappytocyou/repository"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	Db	*sql.DB
	Validate *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		Db:             DB,
		Validate:       validate,
	}
}

func (service *UserServiceImpl) Create(ctx context.Context, request web.UserCreateRequest) (web.UserResponse, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return web.UserResponse{}, err
	}

	tx, err := service.Db.Begin()
	if err != nil {
		return web.UserResponse{}, err
	}
	defer commitOrRollback(tx)

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(request.Password), 14)
	
	user := entity.User{
		Username: request.Username,
		Email: request.Email,
		Password: string(hashedPassword),
		CreatedAt: time.Now(),
	}
	
	user = service.UserRepository.Save(ctx, tx, user)
	
	return web.UserResponse{
		Id:        user.Id,
		Username:  user.Username,
		Email:     user.Email,
	}, nil
}

func (service *UserServiceImpl) ChangePassword(ctx context.Context, request web.UserChangePasswordRequest) (web.UserResponse, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return web.UserResponse{}, err
	}

	tx, err := service.Db.Begin()
	if err != nil {
		return web.UserResponse{}, err
	}

	defer commitOrRollback(tx)

	user, err := service.UserRepository.FindByUsername(ctx, tx, request.Username)
	if err != nil {
		return web.UserResponse{}, err
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(request.Password), 14)

	user.Password = string(hashedPassword)
	user, err = service.UserRepository.ChangePassword(ctx, tx, user)
	if err != nil {
		return web.UserResponse{}, err
	}

	return web.UserResponse{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

func (service *UserServiceImpl) FindByUsername(ctx context.Context, username string) (web.UserResponse, error) {
	tx, err := service.Db.Begin()
	if err != nil {
		return web.UserResponse{}, err
	}
	defer commitOrRollback(tx)

	user, err := service.UserRepository.FindByUsername(ctx, tx, username)
	if err != nil {
		return web.UserResponse{}, err
	}

	return web.UserResponse{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

func (service *UserServiceImpl) Delete(ctx context.Context, id int) error {
	tx, err := service.Db.Begin()
	if err != nil {
		return err
	}

	defer commitOrRollback(tx)

	err = service.UserRepository.Delete(ctx, tx, id)
	if err != nil {
		return err
	}

	return nil
}

func commitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errorRollback := tx.Rollback()
		if errorRollback != nil {
			panic(errorRollback)
		}
		panic(err)
	} else {
		errorCommit := tx.Commit()
		if errorCommit != nil {
			panic(errorCommit)
		}
	}
}


