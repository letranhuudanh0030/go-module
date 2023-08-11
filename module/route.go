package module

import (
	authroute "todo/module/authen/route"
	todoroute "todo/module/item/route"

	_ "todo/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func InitRoute(app *fiber.App) {
	app.Get("/document/*", swagger.HandlerDefault)
	todoroute.InitRoute(app)
	authroute.InitRoute(app)
}
