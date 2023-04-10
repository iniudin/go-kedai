package app

import (
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"github.com/spf13/viper"
	_ "gokedai/docs"
	"gokedai/internal/auth"
	"gokedai/internal/pkg/validation"
	"gokedai/internal/user"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type Server interface {
	InitializeMiddleware()
	InitializeRouter()
	Start()
}

type ServerImpl struct {
	app       *fiber.App
	db        *sql.DB
	validator *validation.CustomValidator
}

func (server *ServerImpl) InitializeMiddleware() {
	server.app.Use(cors.New())
	server.app.Use(recover.New())
	server.app.Use(logger.New())
}

func (server *ServerImpl) InitializeRouter() {
	server.app.Get("/swagger/*", swagger.HandlerDefault)
	server.app.Get("/metrics", monitor.New(monitor.Config{Title: "Katabe Metrics Page"}))
	api := server.app.Group("/api")
	auth.NewAuthRouter(api, server.db, server.validator)
	user.NewUserRouter(api, server.db, server.validator)
}

func (server *ServerImpl) Start() {
	// Listen from a different goroutine
	go func() {
		if err := server.app.Listen(":" + viper.GetString("APP_PORT")); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	_ = <-c // This blocks the main thread until an interrupt is received
	fmt.Println("Gracefully shutting down...")
	_ = server.app.Shutdown()

	fmt.Println("Running cleanup tasks...")
	// Your cleanup tasks go here
	// redisConn.Close()
	fmt.Println("Fiber was successful shutdown.")
}

func NewServerImpl(app *fiber.App, db *sql.DB, validator *validation.CustomValidator) *ServerImpl {
	return &ServerImpl{
		app:       app,
		db:        db,
		validator: validator,
	}
}
