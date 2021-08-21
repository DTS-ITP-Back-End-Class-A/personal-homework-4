package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

/*Soal
buat 2 routing
1. get-car GET
pakek basic auth https://docs.gofiber.io/api/middleware/basicauth#examples
kalau ga username and password gak sama return
a. message : error, status_code : 401
kalau bener : return
car_name, car_color, car_type
2. create-car POST
body json,
{
	"cart_name" : "honda",
	"cart_color" : "red",
	"cart_type" : "matic",
}
return{
	""
}
*/

type Car struct {
	Car_name  string `json:"car_name"`
	Car_color string `json:"car_color"`
	Car_type  string `json:"car_type"`
}

type request struct {
	Car []string
}

var car = []Car{
	{
		Car_name:  "Honda",
		Car_color: "White",
		Car_type:  "Matic",
	},
	{
		Car_name:  "Suzuki",
		Car_color: "Black",
		Car_type:  "Manual",
	},
}

func main() {
	app := fiber.New()

	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"dyah": "luh",
		},
		Realm: "Forbidden",
		Authorizer: func(user, pass string) bool {
			if user == "dyah" && pass == "luh" {
				return true
			}
			return false
		},
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message":     "error",
				"status_code": 401,
			})
		},
		ContextUsername: "_user",
		ContextPassword: "_pass",
	}))

	app.Get("/get-car", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "succes",
			"status":  "OK",
			"data":    car,
		})

	})

	app.Post("/create-car", func(c *fiber.Ctx) error {
		// body := new(request)
		// err := c.BodyParser(&body)
		// response := []Car{}

		// if err != nil {
		// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		// 		"message":     "error",
		// 		"status_code": 401,
		// 	})
		// }
		// for _, req := range body.Car {
		// 	for _, v := range car {
		// 		if req == v.Car_name {
		// 			response = append(response, v)
		// 		}
		// 	}
		// }
		// return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		// 	"success": true,
		// 	"status":  "OK",
		// 	"data":    response,
		// })
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "succes",
			"status":  "OK",
			"data":    car,
		})
	})

	app.Listen(":3000")

}
