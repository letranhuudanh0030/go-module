package middleware

import (
	"todo/config"

	"github.com/gofiber/fiber/v2"
)

// Check Key App
func AppInfo(c *fiber.Ctx) error {
	app_key := c.Get("x-csv-key")
	response := new(config.DataResponse)
	c.Locals("username", "1520")
	if len(app_key) == 0 || app_key != config.Get("APP_KEY") {
		response.Status = false
		// response.Message = config.GetMessageCode("KEY_NOT_FOUND")
		return c.JSON(response)
	}

	return c.Next()
}
