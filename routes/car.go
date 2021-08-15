package routes

import (
	"github.com/DTS-ITP-Back-End-Class-A/personal-homework-4/controllers"
	"github.com/DTS-ITP-Back-End-Class-A/personal-homework-4/middleware"
	"github.com/gofiber/fiber/v2"
)

func CarRoutes(route fiber.Router) {
	route.Post("/", controllers.CreateNewCar)
	route.Use("/", middleware.BasicAuthMiddleware()).Get("/", controllers.GetAllCarsc)
}
