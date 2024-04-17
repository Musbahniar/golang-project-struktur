package routes

import (
	handler "github.com/Musbahniar/golang-project-struktur/internals/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupKotaRoutes(router fiber.Router) {
	kota := router.Group("/kota")

	kota.Get("/", handler.GetKota)
}
