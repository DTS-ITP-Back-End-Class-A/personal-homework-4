package routes

import (
	"github.com/DTS-ITP-Back-End-Class-A/personal-homework-4/controllers"
	"github.com/gofiber/fiber/v2"
)

func GeneralRoutes(route fiber.Router) {
	route.Get("/health", controllers.HealthCheck)
}
