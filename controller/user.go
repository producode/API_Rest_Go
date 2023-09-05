package controller

import (
	"github.com/apirestgo/services"
	"github.com/gofiber/fiber/v2"
)

func LogIn(c *fiber.Ctx) error {
	token, err := services.LogIn()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	return c.JSON(token)
}
