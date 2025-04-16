package routes

import (
    "github.com/gofiber/fiber/v2"
    "go-crm-server/internal/auth/handlers"
    "gorm.io/gorm"
)

// func RegisterAuthRoutes(app *fiber.App) {
//     auth := app.Group("/auth")
//     auth.Post("/login", handlers.LoginHandler)
//     auth.Post("/register", handlers.RegisterHandler)
// }

func SetupRoutes(app *fiber.App, db *gorm.DB) {
    api := app.Group("/api/v1")
    
    // Public routes (no authentication required)
    api.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Server is running!")
    })
    api.Post("/signup", handlers.SignUp)
    api.Post("/signin", handlers.SignIn)
    api.Get("/verify/:token", handlers.VerifyEmail)
    api.Post("/forgot-password", handlers.RequestPasswordReset)
    api.Post("/reset-password", handlers.ResetPassword)
    // Other protected routes
    api.Post("/signout", handlers.SignOut)
    api.Post("/auth/refresh", handlers.RefreshToken)
}