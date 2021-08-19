package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

type data struct {
	car_name  string `json:"car"`
	car_color string `json:"color"`
	car_type  string `json:"type"`
}

func main() {
	// referensi https://dev.to/devsmranjan/golang-build-your-first-rest-api-with-fiber-24eh
	app := fiber.New()

	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"henry": "123456",
		},
		Realm: "Forbidden",
		Authorizer: func(user, pass string) bool {
			if user == "henry" && pass == "123456" {
				return true
			}
			return false
		},
		Unauthorized: func(c *fiber.Ctx) error {
			return c.JSON(fiber.Map{
				"message":     "username/ password is incorrect",
				"status_code": 401,
			})
		},

		ContextUsername: "_user",
		ContextPassword: "_pass",
	}))

	app.Get("/get-car", GetCar)
	app.Post("/create-car", CreateCar)
	app.Delete("/delete-car/:id", DeleteCar)
	app.Listen(":3000")
}

//make struct for car
type Car struct {
	Id       int    `json:"id"`
	CarName  string `json:"car_name"`
	CarColor string `json:"car_color"`
	CarType  string `json:"car_type"`
}

var cars = make([]Car, 0)

func DeleteCar(c *fiber.Ctx) error {
	// get param
	paramId := c.Params("id")

	// convert param string to int
	id, err := strconv.Atoi(paramId)

	// if parameter cannot parse
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot parse id",
			"status":  "gagal",
		})
	}

	// find and delete todo
	for i, todo := range cars {
		if todo.Id == id {
			cars = append(cars[:i], cars[i+1:]...)
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"message": "Deleted Succesfully",
				"status":  "OK",
			})
		}
	}

	// if todo not found
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"message": "Car not found",
		"status":  "gagal",
	})
}

func GetCar(c *fiber.Ctx) error {
	//get car
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"status":  "OK",
		"data":    cars,
	})
}

func CreateCar(c *fiber.Ctx) error {
	//declare struct car ke variable data
	data := new(Car)
	err := c.BodyParser(&data)
	//cek error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":     "error",
			"status_code": 401,
		})
	}
	//validasi
	if data.CarName == "" || data.CarColor == "" || data.CarType == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":     "car_name , car_color , car_type is required",
			"status_code": 401,
		})
	}
	// create a todo variable
	todo := &Car{
		Id:       len(cars) + 1,
		CarName:  data.CarName,
		CarColor: data.CarColor,
		CarType:  data.CarType,
	}
	// append in cars
	cars = append(cars, *todo)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":     "success",
		"status_code": "OK",
		"data":        data,
	})

}
