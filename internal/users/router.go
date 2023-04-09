package users

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"katalog/internal/pkg/middleware"
	"katalog/internal/pkg/validation"
)

func NewUserRouter(router fiber.Router, db *gorm.DB, validator *validation.CustomValidator) {
	service := NewServiceImpl(db)
	controller := NewUserControllerImpl(validator, service)

	users := router.Group("/users", middleware.Protected())
	users.Get("", controller.FindAll)
	users.Get("/:id", controller.FindById)
	users.Put("/:id", controller.Update)
	users.Delete("/:id", controller.Delete)
}
