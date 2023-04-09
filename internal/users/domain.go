package users

import (
	"context"
	"github.com/gofiber/fiber/v2"
)

type Repository interface {
	Save(ctx context.Context, user User)
	FindAll(ctx context.Context, page int, size int) ([]UserResponse, error)
	FindById(ctx context.Context, id uint) (*UserResponse, error)
	Update(ctx context.Context, user User) (*UserResponse, error)
	Delete(ctx context.Context, id uint) error
}

type Service interface {
	FindAll(ctx context.Context, page int, size int) ([]UserResponse, error)
	FindById(ctx context.Context, id uint) (*UserResponse, error)
	Update(ctx context.Context, request UserUpdateRequest) (*UserResponse, error)
	Delete(ctx context.Context, id uint) error
}

type Controller interface {
	FindAll(ctx *fiber.Ctx) error
	FindById(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}
