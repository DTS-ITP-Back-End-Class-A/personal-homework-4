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

	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"admin": "123456",
		},
		Realm: "Forbidden",
		Authorizer: func(user, pass string) bool {
			if user == "admin" && pass == "123456" {
				return true
			}
			return false
		},
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message":     "error",
				"status_code": 401,
			})
		},
		ContextUsername: "_user",
		ContextPassword: "_pass",
	}))

	app.Get("/cars", func(c *fiber.Ctx) error {
		data := Car{
			CarName:  "Honda",
			CarColor: "Putih",
			CarType:  "Matic",
		}
		return c.JSON(data)
	})

	var cars = make([]Car, 0)

	app.Post("/create", func(c *fiber.Ctx) error {
		body := new(Car)
		err := c.BodyParser(&body)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message":     "error",
				"status_code": err,
			})
		}

		if body.CarName == "" || body.CarColor == "" || body.CarType == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message":     "car name, color and type are required",
				"status_code": 401,
			})
		}

		cars = append(cars, *body)
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message":     "succes",
			"status_code": "Ok",
			"data":        body,
		})
	})

	// app.Post("/create", func(c *fiber.Ctx) error {
	// 	body := Car{
	// 		CarName:  "Honda",
	// 		CarColor: "Putih",
	// 		CarType:  "Matic",
	// 	}

	// 	return c.JSON(fiber.Map{
	// 		"message": "success",
	// 		"status":  "ok",
	// 		"data":    body,
	// 	})
	// })

	app.Listen(":3000")
}
