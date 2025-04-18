package routes

import (
    "github.com/gofiber/fiber/v2"
    "go-crm-server/internal/auth/handlers"
    "gorm.io/gorm"
)
func SetupRoutes(app *fiber.App, db *gorm.DB) {
    auth := app.Group("/auth/v1")
    
    // Public routes (no authentication required)
    auth.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Server is running!")
    })

    auth.Get("/signup", handlers.SignUpPage)
    auth.Get("/signin", handlers.SignInPage)
    auth.Get("/verify/:token", handlers.VerifyEmail)
    auth.Get("/signup", handlers.GetSignUp)
    auth.Get("/signin", handlers.GetSignIn)

    auth.Post("/signup", handlers.SignUp)
    auth.Post("/signin", handlers.SignIn)
    auth.Post("/forgot-password", handlers.RequestPasswordReset)
    auth.Post("/reset-password", handlers.ResetPassword)
    // Other protected routes
    auth.Post("/signout", handlers.SignOut)
    auth.Post("/refresh", handlers.RefreshToken)
}