package handler

import "github.com/gofiber/fiber/v2"

// GetOrderByCode Getting Order by Code
//
//	@Summary		Getting Order by Code
//	@Description	Getting Order by Code in detail
//	@Tags			Orders
//	@Accept			json
//	@Produce		json
//	@Param			x-correlationid	header		string	true	"code of Order"
//	@Param			orderCode		path		string	true	"code of Order"
//	@Success		200				{string}	string
//	@Router			/orders/code/{orderCode} [get]
func GetProvinsi(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Provinsi Found", "data": ""})
}
