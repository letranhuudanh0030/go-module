package todoroute

import (
	"todo/middleware"
	todoctr "todo/module/item/controller"

	"github.com/gofiber/fiber/v2"
)

func InitRoute(app *fiber.App) {
	v1 := app.Group("/v1", middleware.AppInfo)
	v1.Post("/items", todoctr.HanleCreateItem)
	v1.Get("/items", todoctr.HandleFindAll)
	v1.Get("/items/:id", todoctr.HandleFindItem)
	v1.Put("/items/:id", todoctr.HandleEditItem)
	v1.Delete("/items/:id", todoctr.HandleDeleteItem)
}
