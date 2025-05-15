package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

func RBACCheckMiddleware(logger zerolog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// TODO: Implement RBAC check logic here

		return c.Next()
	}
}
