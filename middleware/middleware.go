package middleware

import (
	"strings"
	"todo/config"
	"todo/database"
	"todo/util"

	"github.com/gofiber/fiber/v2"
)

// Check Key App
func AppInfo(c *fiber.Ctx) error {
	app_key := c.Get("x-csv-key")
	response := new(config.DataResponse)
	userToken, _ := util.ExtractTokenData(c)
	if userToken != nil {
		c.Locals("username", userToken.Username)
	}
	c.Locals("timezone", "America/Los_Angeles") // header request: Timezone from FE
	if len(app_key) == 0 || app_key != config.Get("APP_KEY") {
		response.Status = false
		response.Message = config.KEY_NOT_FOUND
		return c.JSON(response)
	}

	return c.Next()
}

// Authen
func AppAuthen(c *fiber.Ctx) error {
	store := database.Store
	is_error := 0
	response := new(config.DataResponse)

	// Check token valid
	tokenData, err := util.ExtractTokenData(c)
	if err != nil {
		response.Status = false
		response.Message = config.TOKEN_INCORRECT
		return c.JSON(response)
	}

	// Check session Exist and comparse token
	sess, err := store.Get(tokenData.Username)
	authen := strings.Split(c.Get("x-csv-token"), " ")
	if err != nil || len(sess) == 0 || string(sess) != authen[1] {
		is_error = 1
	}

	if is_error == 1 {
		response.Status = false
		response.Message = config.TOKEN_INCORRECT
		return c.JSON(response)
	}
	return c.Next()
}
