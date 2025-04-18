package handlers

import (
    "log"
	"net/http"
	"github.com/gofiber/fiber/v2"
	"github.com/gibbyDev/OpsMastery/models"
	"golang.org/x/crypto/bcrypt"
	"github.com/gibbyDev/OpsMastery/utils"
	"time"
	"fmt"
	// "os"
)

func VerifyEmail(c *fiber.Ctx) error {
    token := c.Params("token")

    var user models.User
    if err := db.Where("verification_token = ?", token).First(&user).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Invalid verification token",
        })
    }

    user.Active = true
    user.VerificationToken = "" // Clear the token after verification

    if err := db.Save(&user).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to verify email",
        })
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "Email verified successfully",
    })
}