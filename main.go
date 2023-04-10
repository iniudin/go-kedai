package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gokedai/internal/app"
	"gokedai/internal/database"
	"gokedai/internal/pkg/validation"
	"log"
)

// @title 			Kedai API
// @version 		1.0
// @description 	API documentation for Kedai API
// @termsOfService 	http://swagger.io/terms/
// @contact.name 	API Support
// @contact.email 	fiber@swagger.io
// @license.name 	Apache 2.0
// @license.url 	http://www.apache.org/licenses/LICENSE-2.0.html
// @host 			localhost:3000
// @BasePath 		/api
func main() {
	// Load Configuration
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Set Validator and translate Configuration
	validatorConf, translatorConf := validation.InitializeValidator()
	validate := validation.NewValidator(validatorConf, translatorConf)

	server := fiber.New()
	db := database.NewConnect()

	httpServer := app.NewServerImpl(server, db, validate)
	httpServer.InitializeMiddleware()
	httpServer.InitializeRouter()
	httpServer.Start()

}
