package handlers

import "github.com/gofiber/fiber/v2"

func NotFound(c *fiber.Ctx) error {
	c.Context().SetContentType("application/json")
	c.SendStatus(fiber.StatusNotFound)
	return nil
}
