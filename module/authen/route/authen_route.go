package route

import (
	"todo/middleware"
	authenctrl "todo/module/authen/controller"

	"github.com/gofiber/fiber/v2"
)

func InitRoute(app *fiber.App) {
	/**
	*
	*	System User
	* - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
	*
	**/
	api := app.Group("/", middleware.AppInfo)

	/**
	*
	*	Authen
	* - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
	*
	**/
	api.Post("/login", authenctrl.Login)
	api.Post("/check-token", middleware.AppAuthen, authenctrl.CheckToken)

}
