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
)

func ResetPassword(c *fiber.Ctx) error {
    var body struct {
        ResetToken  string `json:"reset_token"`
        NewPassword string `json:"new_password"`
    }
    
    // Log the raw request body
    rawBody := string(c.Body())
    log.Printf("Received reset password request body: %s", rawBody)
    
    if err := c.BodyParser(&body); err != nil {
        log.Printf("Error parsing request body: %v", err)
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request body",
        })
    }

    log.Printf("Reset token: %s", body.ResetToken)
    
    // Validate the token
    var user models.User
    if err := db.Where("reset_token = ?", body.ResetToken).First(&user).Error; err != nil {
        log.Printf("No user found with reset token: %s, error: %v", body.ResetToken, err)
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid reset token",
        })
    }

    // Check if token has expired
    if user.ResetTokenExpiry.Before(time.Now()) {
        log.Printf("Token expired. Expiry: %v, Current time: %v", user.ResetTokenExpiry, time.Now())
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Reset token has expired",
        })
    }

    // Hash the new password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.NewPassword), bcrypt.DefaultCost)
    if err != nil {
        log.Printf("Error hashing password: %v", err)
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to hash password",
        })
    }

    // Update user's password and clear reset token
    user.Password = string(hashedPassword)
    user.ResetToken = ""
    user.ResetTokenExpiry = time.Time{}

    if err := db.Save(&user).Error; err != nil {
        log.Printf("Error saving user: %v", err)
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to update password",
        })
    }

    log.Printf("Password successfully reset for user: %s", user.Email)
    return c.JSON(fiber.Map{
        "message": "Password successfully reset",
    })
}

func validateResetToken(token string) (bool, error) {
    var user models.User
    if err := db.Where("reset_token = ?", token).First(&user).Error; err != nil {
        log.Printf("No user found with reset token: %s, error: %v", token, err)
        return false, err
    }

    // Check if token has expired
    if user.ResetTokenExpiry.Before(time.Now()) {
        log.Printf("Token expired. Expiry: %v, Current time: %v", user.ResetTokenExpiry, time.Now())
        return false, fmt.Errorf("reset token has expired")
    }

    log.Printf("Token validated successfully for user: %s", user.Email)
    return true, nil
}