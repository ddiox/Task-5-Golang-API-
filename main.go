package main

import (
	"github.com/ddiox/evermos_api/config"
	"github.com/ddiox/evermos_api/routers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	config.ConnectDB()

	app.Use(logger.New())

	app.Use(cors.New())

	routers.SetupRoutes(app)

	app.Listen(":8080")
}
