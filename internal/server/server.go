package server

import (
	"github.com/gofiber/fiber/v2"

	"go-crm-server/internal/database"
)

type FiberServer struct {
	*fiber.App

	db database.Service
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "go-crm-server",
			AppName:      "go-crm-server",
		}),

		db: database.New(),
	}

	return server
}
