package handlers

import (
	"net/http"
	"strconv"
	"github.com/gofiber/fiber/v2"
	"github.com/gibbyDev/go-crm-server/models"
	"gorm.io/gorm"
)

func CreateTicket(c *fiber.Ctx) error {
	var ticket models.Ticket
	if err := c.BodyParser(&ticket); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := db.Create(&ticket).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusCreated).JSON(ticket)
}
