package router

import (
	"github.com/Musbahniar/golang-project-struktur/internals/routes"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	apiV1 := app.Group("/api/v1")

	// Setup note routes, can use same syntax to add routes for more models
	routes.SetupProvinsiRoutes(apiV1)
	routes.SetupKotaRoutes(apiV1)
}
