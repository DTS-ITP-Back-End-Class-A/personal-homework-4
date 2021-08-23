package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func main() {
	app := fiber.New()

	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"rifky": "admin123",
			"tedi":  "admin789",
		},
		Realm: "Forbidden",
		Authorizer: func(userName, password string) bool {
			if userName == "rifky" && password == "admin123" {
				return true
			}
			if userName == "tedi" && password == "admin789" {
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

	app.Get("/cars", GetCars)
	app.Post("/create", CreateCar)

	app.Listen(":3001")
}

type Car struct {
	CarName  string `"json: "car_name"`
	CarColor string `"json: "car_color"`
	CarType  string `"json: "car_type"`
}

//list cars
var cars = []Car{
	{CarName: "Ferrari", CarColor: "Black", CarType: "Manual"},
	{CarName: "Ferrari", CarColor: "Yellow", CarType: "Manual"},
}

func GetCars(c *fiber.Ctx) error {
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success",
		"status":  "OK",
		"data":    cars,
	})
}

func CreateCar(c *fiber.Ctx) error {
	body := new(Car)
	err := c.BodyParser(&body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":     "error",
			"status_code": 401,
		})
	}

	if body.CarName == "" || body.CarColor == "" || body.CarType == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":     "Car name, color and type are required",
			"status_code": 401,
		})
	}

	cars = append(cars, *body)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":     "Success",
		"status_code": "Ok",
		"data":        body,
	})

}
