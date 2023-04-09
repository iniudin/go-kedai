package auth

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"katalog/internal/pkg/middleware"
	"katalog/internal/pkg/validation"
)

func NewAuthRouter(router fiber.Router, db *gorm.DB, validator *validation.CustomValidator) {
	service := NewAuthServiceImpl(db)
	controller := NewAuthControllerImpl(validator, service)
	auth := router.Group("/auth")
	auth.Post("/register", controller.Register)
	auth.Post("/login", controller.Login)
	auth.Get("/refresh", middleware.Protected(), controller.Refresh)
}
