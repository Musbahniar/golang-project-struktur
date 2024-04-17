package routes

import (
	handler "github.com/Musbahniar/golang-project-struktur/internals/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupProvinsiRoutes(router fiber.Router) {
	provinsi := router.Group("/provinsi")

	// Read all Notes
	provinsi.Get("/", handler.GetProvinsi)
}
