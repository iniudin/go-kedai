package user

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"gokedai/internal/pkg/middleware"
	"gokedai/internal/pkg/validation"
)

func NewUserRouter(router fiber.Router, db *sql.DB, validator *validation.CustomValidator) {
	service := NewServiceImpl(db)
	controller := NewUserControllerImpl(validator, service)

	users := router.Group("/user", middleware.Protected())
	users.Get("", controller.FindAll)
	users.Get("/:id", controller.FindById)
	users.Put("/:id", controller.Update)
	users.Delete("/:id", controller.Delete)
}
