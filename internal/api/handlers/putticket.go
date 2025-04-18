package handlers

import (
	"net/http"
	"strconv"
	"github.com/gofiber/fiber/v2"
	"github.com/gibbyDev/go-crm-server/models"
	"gorm.io/gorm"
)

func UpdateTicketByID(c *fiber.Ctx) error {
	id := c.Params("id") 
	ticketID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ticket ID"})
	}

	var ticket models.Ticket
	if err := c.BodyParser(&ticket); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := db.Model(&models.Ticket{}).Where("id = ?", ticketID).Updates(ticket).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Ticket updated successfully"})
}