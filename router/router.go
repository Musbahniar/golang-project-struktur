package router

import (
	"github.com/Musbahniar/golang-project-struktur/internals/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {

	// Group api calls with param '/api'
	api := app.Group("/api", logger.New())

	// Setup note routes, can use same syntax to add routes for more models
	routes.SetupProvinsiRoutes(api)
	routes.SetupKotaRoutes(api)
}
