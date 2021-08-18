package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

type Car struct {
	Car_name  string
	Car_color string
	Car_type  string
}

func main() {
	app := fiber.New()

	// Provide a minimal config
	// app.Use(basicauth.New(basicauth.Config{
	// 	Users: map[string]string{
	// 		"darin": "123",
	// 	},
	// }))

	// Or extend your config for customization
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"darin": "123",
		},
		Realm: "Forbidden",
		Authorizer: func(user, pass string) bool {
			if user == "darin" && pass == "123" {
				return true
			}
			return false
		},
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Status(500).JSON(&fiber.Map{
				"message": "username/ password is incorrect",
				"status":  401,
			})
		},
		ContextUsername: "_user",
		ContextPassword: "_pass",
	}))

	app.Get("/get-car", func(c *fiber.Ctx) error {
		//create data struct
		data := Car{
			Car_name:  "Honda",
			Car_color: "Merah",
			Car_type:  "Matic",
		}

		return c.JSON(data)
	})

	app.Post("/create-car", func(c *fiber.Ctx) error { //jangan lupa tanda /(slash) di depan url
		//create data struct
		data1 := Car{
			Car_name:  "Honda",
			Car_color: "Gold",
			Car_type:  "Matic",
		}

		return c.JSON(fiber.Map{
			"message": "success",
			"status":  "ok",
			"data":    data1,
		})
	})

	app.Listen(":3000")

}
