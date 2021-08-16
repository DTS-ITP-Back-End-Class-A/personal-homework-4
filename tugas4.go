package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func main() {
	app := fiber.New()
	//basicAuth
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"eka":   "123",
			"rahma": "456",
		},
		Realm: "Forbidden",
		Authorizer: func(user, pass string) bool {
			if user == "eka" && pass == "123" {
				return true
			}
			if user == "rahma" && pass == "456" {
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

	//GetAll,GetById,Post
	app.Get("/cars", GetCars)
	app.Post("/create", CreateCar)
	app.Listen(":8081")

}

//make struct for car
type Car struct {
	CarName  string `json:"car_name"`
	CarColor string `json:"car_color"`
	CarType  string `json:"car_type"`
}

//list car
// var cars = []Car{
// 	{CarName: "Honda", CarColor: "red", CarType: "matic"},
// 	{CarName: "Honda", CarColor: "purple", CarType: "matic"},
// }
var cars = make([]Car, 0)

//getAllCar
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
			"message":     "car name, color and type are required",
			"status_code": 401,
		})
	}

	cars = append(cars, *body)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":     "success",
		"status_code": "OK",
		"data":        body,
	})

}
