package service

import (
	"context"
	"github.com/yogamuris/sohappytocyou/entity/web"
)

type UserService interface {
	Create(ctx context.Context, request web.UserCreateRequest) (web.UserResponse, error)
	ChangePassword(ctx context.Context, request web.UserChangePasswordRequest) (web.UserResponse, error)
	FindById(ctx context.Context, id int) (web.UserResponse, error)
	Delete(ctx context.Context, id int) error
}