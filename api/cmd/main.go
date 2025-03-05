package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/aticish/log-service/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New(fiber.Config{
		IdleTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
		Concurrency:  256 * 1024,
	})

	// Write & Read log
	app.Post("/api/v1", func(c *fiber.Ctx) error {
		return handlers.VersionOne(c)
	})

	// Default
	app.Use(func(c *fiber.Ctx) error {
		return handlers.NotFound(c)
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

	// Wait for shutdown signal
	<-sigChan
	log.Print("Shutting down server...")

	// Shutdown after 5 seconds
	if err := app.ShutdownWithTimeout(5 * time.Second); err != nil {
		log.Fatal("Error during server shutdown. %w", err)
	}
	log.Print("Server gracefully stopped!")
}
