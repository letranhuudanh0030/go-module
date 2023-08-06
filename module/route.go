package module

import (
	todoroute "todo/module/item/route"

	"github.com/gofiber/fiber/v2"
)

func InitRoute(app *fiber.App) {
	todoroute.InitRoute(app)
}
