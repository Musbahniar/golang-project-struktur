package provinsiRoutes

import (
	provinsiHandler "github.com/Musbahniar/golang-project-struktur/internals/handlers/provinsi"
	"github.com/gofiber/fiber/v2"
)

func SetupProvinsiRoutes(router fiber.Router) {
	provinsi := router.Group("/provinsi")

	// Read all Notes
	provinsi.Get("/", provinsiHandler.GetProvinsi)
}
