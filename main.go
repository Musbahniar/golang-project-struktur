package main

import (
	_ "github.com/Musbahniar/golang-project-struktur/docs"
	"github.com/Musbahniar/golang-project-struktur/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// @title			Belajar Resful Api
// @version		1.0
// @description	Documentation for Belajar Resful Api
// @termsOfService	http://swagger.io/terms/
func main() {
	// Start a new fiber app
	app := fiber.New()

	app.Get("/swagger/*", swagger.HandlerDefault)

	// Health check
	app.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString("And the API is UP!")
		return err
	})

	router.SetupRoutes(app)

	// Listen on PORT 3000
	app.Listen(":3000")
}
