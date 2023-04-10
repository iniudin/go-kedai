package user

import (
	"github.com/gofiber/fiber/v2"
	"katalog/internal/pkg/response"
	"katalog/internal/pkg/validation"
	"net/http"
	"strconv"
)

type ControllerImpl struct {
	validator *validation.CustomValidator
	service   Service
}

func (controller *ControllerImpl) FindAll(ctx *fiber.Ctx) error {
	page, _ := strconv.Atoi(ctx.Query("page", "1"))
	size, _ := strconv.Atoi(ctx.Query("size", "10"))

	users, err := controller.service.FindAll(ctx.Context(), page, size)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(response.WebResponse{
			Status:  http.StatusText(http.StatusInternalServerError),
			Message: err.Error(),
			Data:    nil,
		})

	}

	return ctx.Status(http.StatusOK).JSON(response.WebResponse{
		Status:  http.StatusText(http.StatusOK),
		Message: "Success get all user",
		Data:    users,
	})
}

func (controller *ControllerImpl) FindById(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id")
	user, err := controller.service.FindById(ctx.Context(), uint(id))
	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(response.WebResponse{
			Status:  http.StatusText(http.StatusNotFound),
			Message: "User not found",
			Data:    nil,
		})
	}

	return ctx.Status(http.StatusOK).JSON(response.WebResponse{
		Status:  http.StatusText(http.StatusOK),
		Message: "Success get a user",
		Data:    user,
	})
}

func (controller *ControllerImpl) Update(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id")
	body := UserUpdateRequest{
		Id: uint(id),
	}

	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(response.WebResponse{
			Status:  http.StatusText(http.StatusBadRequest),
			Message: err.Error(),
			Data:    nil,
		})
	}
	if err := controller.validator.Validate(&body); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(response.WebResponse{
			Status:  http.StatusText(http.StatusBadRequest),
			Message: "Some field should be provided",
			Data:    err,
		})
	}
	user, err := controller.service.Update(ctx.Context(), body)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(response.WebResponse{
			Status:  http.StatusText(http.StatusInternalServerError),
			Message: err.Error(),
			Data:    nil,
		})
	}
	return ctx.Status(http.StatusOK).JSON(response.WebResponse{
		Status:  http.StatusText(http.StatusOK),
		Message: "Success update user",
		Data:    user,
	})
}

func (controller *ControllerImpl) Delete(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id")
	err := controller.service.Delete(ctx.Context(), uint(id))
	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(response.WebResponse{
			Status:  http.StatusText(http.StatusNotFound),
			Message: "User not found",
			Data:    nil,
		})
	}

	return ctx.Status(http.StatusOK).JSON(response.WebResponse{
		Status:  http.StatusText(http.StatusOK),
		Message: "Success delete a user",
		Data:    nil,
	})
}

func NewUserControllerImpl(validator *validation.CustomValidator, service Service) *ControllerImpl {
	return &ControllerImpl{
		validator: validator,
		service:   service,
	}
}
