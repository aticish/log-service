package handlers

import (
	"log"

	"github.com/aticish/log-service/actions"
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

	// Only read and write methods allowed for logs
	if request.Method != "write" && request.Method != "read" {
		return c.Status(fiber.StatusMethodNotAllowed).JSON(&internal.Response{
			Code:    fiber.StatusMethodNotAllowed,
			Status:  "error",
			Message: internal.ErrorInvalidLogMethod.Error(),
		})
	}

	var response *internal.Response

	// Read logs
	if request.Method == "read" {
		response, err = actions.Read(request.Data)
	} else if request.Method == "write" {
		response, err = actions.Write(request.Data)
	}

	// request not completed
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&internal.Response{
			Code:    fiber.StatusInternalServerError,
			Status:  "error",
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
