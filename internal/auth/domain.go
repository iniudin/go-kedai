package auth

import (
	"context"
	"github.com/gofiber/fiber/v2"
)

type Service interface {
	Login(ctx context.Context, request *LoginRequest) (*LoginResponse, error)
	Register(ctx context.Context, request RegisterRequest) (*RegisterResponse, error)
	Refresh(ctx context.Context, request AccessTokenRequest) (*AccessTokenResponse, error)
}

type Controller interface {
	Login(ctx *fiber.Ctx) error
	Register(ctx *fiber.Ctx) error
	Refresh(ctx *fiber.Ctx) error
}
