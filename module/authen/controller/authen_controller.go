package controller

import (
	"fmt"
	"todo/config"
	"todo/module/authen/model"
	"todo/module/authen/service"
	"todo/util"

	"github.com/gofiber/fiber/v2"
)

type LoginParam struct {
	Username string `json:"username" `
	Password string `json:"password"`
}

// Login ================================================================================
// @Tags User
// @Summary Login
// @Description User login
// @Param body body LoginParam true " "
// @Accept  json
// @Produce  json
// @Success 200 {object} config.DataResponse "desc"
// @Router /login [post]
// @Security ApiKeyAuth
func Login(c *fiber.Ctx) error {

	response := new(config.DataResponse)
	response.Status = false
	var input model.LoginInput

	if err := c.BodyParser(&input); err != nil {
		response.Message = config.PARAM_ERROR
		return c.JSON(response)
	}

	if errors := util.Validator(input); errors != nil {
		response.Message = config.VALIDATE
		response.ValidateError = errors
		return c.JSON(response)
	}

	resultData, message, status := service.LoginApi(input)

	// Return response
	response.Status = status
	response.Message = fmt.Sprintf("%v", message)
	response.Data = resultData
	return c.JSON(response)
}

// Check token ================================================================================
func CheckToken(c *fiber.Ctx) error {
	response := new(config.DataResponse)
	response.Status = true

	// Return response
	return c.JSON(response)
}
