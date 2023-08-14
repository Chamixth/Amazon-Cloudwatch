package main

import (
	"amazon-cloudwatch/api"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		BodyLimit: 40 * 1024 * 1024, // Set the body limit to 40MB
	})

	app.Post("putDashboard", api.PutDashboard)

	if err := app.Listen(":7070"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}

}
