package handler

import "github.com/gofiber/fiber/v2"

func GetKota(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Kota Found", "data": ""})
}
