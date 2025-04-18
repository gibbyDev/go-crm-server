package handlers

import (
    "log"
	"net/http"
	"github.com/gofiber/fiber/v2"
	"github.com/gibbyDev/go-crm-server/models"
	"golang.org/x/crypto/bcrypt"
	"github.com/gibbyDev/go-crm-server/utils"
	"time"
	"fmt"
)

func SignIn(c *fiber.Ctx) error {
    var userInput struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    if err := c.BodyParser(&userInput); err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
    }

    var user models.User
    if err := db.Where("email = ?", userInput.Email).First(&user).Error; err != nil {
        return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid email or password"})
    }

    if !user.Active {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "error": "Please verify your email before signing in",
        })
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil {
        return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid email or password"})
    }

    accessToken, refreshToken, err := utils.GenerateJWT(user)
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Could not generate tokens"})
    }

    log.Printf("Setting access token: %s", accessToken)
    log.Printf("Setting refresh token: %s", refreshToken)

    c.Cookie(&fiber.Cookie{
        Name:     "access_token",
        Value:    accessToken,
        Expires:  time.Now().Add(15 * time.Minute),
        HTTPOnly: true,
        Secure:   false,
        SameSite: "None",
        Domain:   "",
        Path:     "/",
    })

    c.Cookie(&fiber.Cookie{
        Name:     "refresh_token",
        Value:    refreshToken,
        Expires:  time.Now().Add(7 * 24 * time.Hour),
        HTTPOnly: true,
        Secure:   false,
        SameSite: "None",
        Domain:   "",
        Path:     "/",
    })

    return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Sign in successful"})
}