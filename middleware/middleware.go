package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func BasicAuthMiddleware() func(*fiber.Ctx) error {
	basicAuthHandler := basicauth.New(basicauth.Config{
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
				"success": false,
				"message": "User Doesn't have Access",
				"error":   "Unauthorized",
			})
		},
		ContextUsername: "_user",
		ContextPassword: "_pass",
	})

	return basicAuthHandler
}
