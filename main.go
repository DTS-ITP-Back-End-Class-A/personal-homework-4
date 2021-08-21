package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

type Car struct {
	CarName  string `json:"car_name"`
	CarColor string `json:"car_color"`
	CarType  string `json:"car_type"`
}

func main() {
	app := fiber.New()

	//  pake basic auth
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"zazhil": "zzzzzz",
			"admin": "12345",
		},
		Realm: "Forbidden",
		Authorizer: func(user, pass string) bool {
			if user == "zazhil" && pass == "zzzzzz" {
				return true
			}
			if user == "admin" && pass == "12345" {
				return true
			}
			return false
		},
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Status(401).JSON(&fiber.Map{
				"message": "username / password is incorrect",
			})
		},
		ContextUsername: "_user",
		ContextPassword: "_pass",
	}))

	car := Car{
		CarName:  "Toyota",
		CarColor: "Black",
		CarType:  "Manual",
	}

	// routing 1 get-car
	app.Get("/get-car", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "success",
            "status": "ok",
            "data" : car,
		})
	})

	// routing 2 create-car
	app.Post("/create-car", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"message": "success",
			"status": "ok",
			"data": car,
		})
	})

	app.Listen(":3000")
}