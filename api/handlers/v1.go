package handlers

import (
	"github.com/aticish/log-service/internal"
	"github.com/gofiber/fiber/v2"
)

func VersionOne(c *fiber.Ctx) error {
	c.Context().SetContentType("application/json")

	// Read request body
	var request *internal.Request
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&internal.Response{
			Code:    fiber.StatusBadRequest,
			Status:  "error",
			Message: internal.ErrorInvalidJSON.Error(),
		})
	}
	return nil
}
