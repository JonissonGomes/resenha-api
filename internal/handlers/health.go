package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// HealthCheck retorna o status da API
func HealthCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "ok",
		"message": "API is running",
	})
} 