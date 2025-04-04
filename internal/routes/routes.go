package routes

import (
	"github.com/gofiber/fiber/v2"
	"resenha-api/internal/handlers"
)

func SetupRoutes(app *fiber.App) {
	// Health Check
	app.Get("/health", handlers.HealthCheck)

	// API Routes
	api := app.Group("/api")
	api.Get("/players", handlers.GetPlayers)
}
