package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"gokedai/internal/pkg/response"
	"gokedai/internal/pkg/validation"
	"net/http"
)

type ControllerImpl struct {
	validator *validation.CustomValidator
	service   Service
}

func (controller *ControllerImpl) Login(ctx *fiber.Ctx) error {
	body := new(LoginRequest)
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
	token, err := controller.service.Login(ctx.Context(), body)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(response.WebResponse{
			Status:  http.StatusText(http.StatusInternalServerError),
			Message: err.Error(),
			Data:    nil,
		})
	}
	return ctx.Status(http.StatusOK).JSON(response.WebResponse{
		Status:  http.StatusText(http.StatusOK),
		Message: "Success login user",
		Data:    token,
	})
}

func (controller *ControllerImpl) Register(ctx *fiber.Ctx) error {
	body := RegisterRequest{}
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
	user, err := controller.service.Register(ctx.Context(), body)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(response.WebResponse{
			Status:  http.StatusText(http.StatusInternalServerError),
			Message: err.Error(),
			Data:    nil,
		})
	}
	return ctx.Status(http.StatusOK).JSON(response.WebResponse{
		Status:  http.StatusText(http.StatusOK),
		Message: "Success register new user",
		Data:    user,
	})
}

func (controller *ControllerImpl) Refresh(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["email"].(int)
	email := claims["email"].(string)

	token, err := controller.service.Refresh(ctx.Context(), AccessTokenRequest{Id: uint(id), Email: email})
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(response.WebResponse{
			Status:  http.StatusText(http.StatusInternalServerError),
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.Status(http.StatusOK).JSON(response.WebResponse{
		Status:  http.StatusText(http.StatusOK),
		Message: "Success get access token user",
		Data:    token,
	})
}

func NewAuthControllerImpl(validator *validation.CustomValidator, service Service) *ControllerImpl {
	return &ControllerImpl{
		validator: validator,
		service:   service,
	}
}
