package middleware

import (
    "github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
    // Example: Check for a valid token
    token := c.Get("Authorization")
    if token == "" {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
    }
    return c.Next()
}