package todoroute

import (
	todoctr "todo/module/item/controller"

	"github.com/gofiber/fiber/v2"
)

func InitRoute(app *fiber.App) {
	v1 := app.Group("/v1")
	v1.Post("/items", todoctr.HanleCreateItem)   // updated
	v1.Get("/items", todoctr.HandleFindAll)      // updated
	v1.Get("/items/:id", todoctr.HandleFindItem) // updated
}
