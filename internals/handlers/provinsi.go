package handler

import "github.com/gofiber/fiber/v2"

func GetProvinsi(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Provinsi Found", "data": ""})
}
