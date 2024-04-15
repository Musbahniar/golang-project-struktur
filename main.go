package main

import (
	"github.com/Musbahniar/golang-project-struktur/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Start a new fiber app
	app := fiber.New()

	// Health check
	app.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString("And the API is UP!")
		return err
	})

	router.SetupRoutes(app)

	// Listen on PORT 3000
	app.Listen(":3000")
}
