package controllers

import (
	"github.com/DTS-ITP-Back-End-Class-A/personal-homework-4/model"
	"github.com/gofiber/fiber/v2"
)

var (
	Cars = make([]model.Car, 0)
)

func CreateNewCar(c *fiber.Ctx) error {
	body := new(model.Car)
	err := c.BodyParser(&body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	if body.CarColor == "" || body.CarName == "" || body.CarType == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Car Color, Name and Type are required",
		})
	}
	Cars = append(Cars, *body)

	if len(Cars) == 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Fail Add New Car",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"status":  "OK",
		"data":    body,
	})
}

func GetAllCarsc(c *fiber.Ctx) error {
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"status":  "OK",
		"data":    Cars,
	})
}
