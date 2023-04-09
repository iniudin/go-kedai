package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
	"os"
)

// Protected protect routes
func Protected() fiber.Handler {

	readFile, err := os.ReadFile("cert/public.pem")
	if err != nil {
		panic(err)
	}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(readFile)
	if err != nil {
		panic(err)
	}

	return jwtware.New(jwtware.Config{
		SigningMethod: jwtware.RS256,
		SigningKey:    publicKey,
		ErrorHandler:  jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
}
