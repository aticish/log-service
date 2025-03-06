package handlers

import (
	"github.com/aticish/log-service/internal"
	"github.com/gofiber/fiber/v2"
)

func NotFound(c *fiber.Ctx) error {
	c.Context().SetContentType("application/json")
	c.SendStatus(fiber.StatusNotFound)
	return c.JSON(&internal.Response{
		Code:    fiber.StatusNotFound,
		Status:  "error",
		Message: internal.ErrorPageNotFound.Error(),
		Data:    nil,
	})
}
