package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

type Car struct {
	carName, carColor, carType string
}

func main() {
	app := fiber.New()

	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"ahmad": "12345",
		},
		Realm: "Restricted",
		Authorizer: func(user, pass string) bool {
			if user == "ancas" || pass == "12345" {
				return true
			}
			return false
		},
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"message": "unkown",
				"status":  401,
			})
		},
		ContextUsername: "_user",
		ContextPassword: "_pass",
	}))

	app.Get("/getcar", func(c *fiber.Ctx) error {
		mobil := Car{
			carName: "Avanza", carColor: "Silver", carType: "Manual",
		}
		return c.JSON(mobil)
	})
	app.Post("/createcar", func(c *fiber.Ctx) error {
		tambah := Car{
			carName: "Xenia", carColor: "Putih", carType: "Matic",
		}
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "success", "status": "ok", "input baru": tambah,
		})
	})
	app.Listen(":3000")
}
