package handlers

import (
	"log"

	"github.com/aticish/log-service/internal"
	"github.com/gofiber/fiber/v2"
)

func VersionOne(c *fiber.Ctx) error {
	c.Context().SetContentType("application/json")

	// Read request body
	var request *internal.Request
	if err := c.BodyParser(&request); err != nil {
		log.Printf("Error parsing request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(&internal.Response{
			Code:    fiber.StatusBadRequest,
			Status:  "error",
			Message: internal.ErrorInvalidJSON.Error(),
		})
	}

	// Check token
	err := internal.CheckToken(c.Get("Authorization"))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(&internal.Response{
			Code:    fiber.StatusUnauthorized,
			Status:  "error",
			Message: err.Error(),
		})
	}
	return nil
}
