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

func SignUp(c *fiber.Ctx) error {
    var user models.User
    var input struct {
        Email    string `json:"email"`
        Password string `json:"password"`
        Name     string `json:"name"`
        Role     string `json:"role"`
    }

    if err := c.BodyParser(&input); err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }

    user.Email = input.Email
    user.Password = input.Password
    user.Name = input.Name
    user.Role = input.Role
    user.Active = false
    user.VerificationToken = utils.GenerateRandomToken()

    storedHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to hash password"})
    }
    user.Password = string(storedHash)

    log.Printf("Hashed password for user %s: %s\n", user.Email, user.Password)

    if err := db.Create(&user).Error; err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    if err := utils.SendVerificationEmail(user.Email, user.VerificationToken); err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to send verification email"})
    }

    return c.Status(http.StatusCreated).JSON(fiber.Map{
        "message": "Registration successful. Please check your email to verify your account.",
    })
}
