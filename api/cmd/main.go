package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Its working")
	})

	// Create channel for shurtdown signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Run server as goroutine
	go func() {
		if err := app.Listen(":3000"); err != nil {
			log.Fatalf("Failed to start server on localhost: %d", 3000)
			os.Exit(1)
		}
	}()

	log.Printf("Server started at localhost:%d", 3000)

	// Wait for shutdown signal
	<-sigChan
	log.Print("Shutting down server...")

	// Shutdown after 5 seconds
	if err := app.ShutdownWithTimeout(5 * time.Second); err != nil {
		log.Fatal("Error during server shutdown. %w", err)
	}
	log.Print("Server gracefully stopped")
}
