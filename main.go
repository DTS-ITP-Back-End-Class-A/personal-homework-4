package main

import (
	"github.com/DTS-ITP-Back-End-Class-A/personal-homework-4/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func setupRoutes(app *fiber.App) {

	api := app.Group("/")
	carApi := app.Group("/car")

	routes.GeneralRoutes(api)
	routes.CarRoutes(carApi)
}

func main() {
	app := fiber.New()

	app.Use(logger.New())

	setupRoutes(app)

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
