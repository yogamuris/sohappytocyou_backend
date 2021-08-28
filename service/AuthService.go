package service

import (
	"context"
	"github.com/yogamuris/sohappytocyou/entity/web"
)

type AuthService interface {
	Login(ctx context.Context, request web.LoginRequest) (web.AuthResponse, error)
	Register(ctx context.Context, request web.RegisterRequest) (web.AuthResponse, error)
	Verify(ctx context.Context, request web.VerifyRequest) (web.AuthResponse, error)
}
