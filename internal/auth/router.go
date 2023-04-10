package auth

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"gokedai/internal/pkg/middleware"
	"gokedai/internal/pkg/validation"
)

func NewAuthRouter(router fiber.Router, db *sql.DB, validator *validation.CustomValidator) {
	service := NewAuthServiceImpl(db)
	controller := NewAuthControllerImpl(validator, service)
	auth := router.Group("/auth")
	auth.Post("/register", controller.Register)
	auth.Post("/login", controller.Login)
	auth.Get("/refresh", middleware.Protected(), controller.Refresh)
}
