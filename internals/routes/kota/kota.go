package kotaRoutes

import (
	kotaHandler "github.com/Musbahniar/golang-project-struktur/internals/handlers/kota"
	"github.com/gofiber/fiber/v2"
)

func SetupKotaRoutes(router fiber.Router) {
	kota := router.Group("/kota")

	kota.Get("/", kotaHandler.GetKota)
}
