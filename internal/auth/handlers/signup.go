package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func SignUp(c *fiber.Ctx) error {
	// Simulate a registration process
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "" || password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Username and password are required",
		})
	}

	// Here you would typically save the user to your database

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User registered successfully",
	})
}