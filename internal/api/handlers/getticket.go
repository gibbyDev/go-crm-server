package handlers

import (
	"net/http"
	"strconv"
	"github.com/gofiber/fiber/v2"
	"github.com/gibbyDev/go-crm-server/models"
	"gorm.io/gorm"
)

func ListTickets(c *fiber.Ctx) error {
	var tickets []models.Ticket
	if err := db.Find(&tickets).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(tickets)
}

func GetTicketByID(c *fiber.Ctx) error {
	id := c.Params("id") 
	var ticket models.Ticket

	ticketID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ticket ID"})
	}

	if err := db.First(&ticket, ticketID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Ticket not found"})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(ticket)
}
