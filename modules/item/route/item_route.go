package todoroute

import (
	"todo/middleware"
	todoctrl "todo/modules/item/controller"

	"github.com/gofiber/fiber/v2"
)

func InitRoute(app *fiber.App) {
	v1 := app.Group("/v1", middleware.AppInfo)
	v1.Post("/items", todoctrl.HanleCreateItem)
	v1.Get("/items", todoctrl.HandleFindAll)
	v1.Get("/items/:id", todoctrl.HandleFindItem)
	v1.Put("/items/:id", todoctrl.HandleEditItem)
	v1.Delete("/items/:id", todoctrl.HandleDeleteItem)
}
