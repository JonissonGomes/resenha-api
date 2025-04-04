package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func GetPlayers(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Lista de jogadores",
	})
}
