package main

import (
	"fmt"
	"kage/config"
	"kage/utils/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
)

var (
	log = logger.NewLogger().WithPrefix("Main Process")
)

func main() {
	if config.Config.Debug {
		log.WithTimestamp().WithTimeFormat("02/01/2006 03:04:05 PM")
	}

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:  "*",
		AllowMethods:  "GET, HEAD, PUT, PATCH, POST, DELETE, OPTIONS",
		AllowHeaders:  "Origin, Content-Type, Accept, Authorization, X-Requested-With, X-API-Key, X-CSRF-Token",
		ExposeHeaders: "Content-Length, Content-Type, Content-Disposition, X-Pagination, X-Total-Count",
		MaxAge:        86400,
	}))
	app.Use(helmet.New())

	log.Infof("Attempting to start server on Port %d...", config.Config.Port)

	if err := app.Listen(fmt.Sprintf(":%d", config.Config.Port)); err != nil {
		log.Fatalf("Failed to start server on port %d: %v", config.Config.Port, err)
	} else {
		log.Successf("Server started on port %d", config.Config.Port)
	}
}
