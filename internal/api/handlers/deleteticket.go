package handlers

import (
	"net/http"
	"strconv"
	"github.com/gofiber/fiber/v2"
	"github.com/gibbyDev/go-crm-server/models"
	"gorm.io/gorm"
)

func DeleteTicketByID(c *fiber.Ctx) error {
	id := c.Params("id")
	ticketID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ticket ID"})
	}

	result := db.Unscoped().Delete(&models.Ticket{}, ticketID)
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Ticket not found"})
	}
	if result.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Ticket deleted successfully"})
}
