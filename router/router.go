package router

import (
	kotaRoutes "github.com/Musbahniar/golang-project-struktur/internals/routes/kota"
	provinsiRoutes "github.com/Musbahniar/golang-project-struktur/internals/routes/provinsi"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {

	// Group api calls with param '/api'
	api := app.Group("/api", logger.New())

	// Setup note routes, can use same syntax to add routes for more models
	provinsiRoutes.SetupProvinsiRoutes(api)
	kotaRoutes.SetupKotaRoutes(api)
}
