package handlers

import "github.com/gofiber/fiber/v2"

func VersionOne(c *fiber.Ctx) error {
	c.Context().SetContentType("application/json")
	return nil
}
