package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Its working")
	})

	log.Fatal(app.Listen("0.0.0.0:3000"))
}
