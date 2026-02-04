package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
)

func LoggingMiddleware() fiber.Handler {
	return func(c fiber.Ctx) error {
		// Log the incoming request
		fmt.Printf("Incoming request: %s %s\n", c.Method(), c.OriginalURL())
		return c.Next()
	}
}
