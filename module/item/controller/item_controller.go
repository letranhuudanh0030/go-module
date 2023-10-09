package todoctrl

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"

	"todo/config"
	"todo/database"
	domainmdl "todo/domain/model"
	todorepo "todo/module/item/repository"
	todoservice "todo/module/item/service"
	"todo/util"
)

func setup(c *fiber.Ctx) *todoservice.ToDoBiz {
	// setup dependencies
	db := database.DB.Set("username", c.Locals("username"))
	db.Set("timezone", c.Locals("timezone"))
	storage := todorepo.NewPostgreSQLStorage(db)
	return todoservice.ToDoItemBiz(storage)
}

// Create a new item ================================================================================
// @Tags Item
// @Summary Create a new item
// @Description Create a new item
// @Param body body domainmdl.CreateItem true " "
// @Accept  json
// @Produce  json
// @Success 200 {object} config.DataResponse "desc"
// @Router /v1/items [post]
// @Security ApiKeyAuth
// @Security ApiTokenAuth
func HanleCreateItem(c *fiber.Ctx) error {
	var response config.DataResponse
	var dataItem domainmdl.ToDoItem

	if err := c.BodyParser(&dataItem); err != nil {
		response.Message = config.PARAM_ERROR
		return c.JSON(response)
	}

	if errors := util.Validator(dataItem); errors != nil {
		response.Message = config.VALIDATE
		response.ValidateError = errors
		return c.JSON(response)
	}

	// pre-process title - trim all spaces
	dataItem.Title = strings.TrimSpace(dataItem.Title)

	if err := setup(c).CreateItem(c, &dataItem); err != nil {
		response.Message = err.Error()
		return c.JSON(response)
	}

	response.Status = true
	response.Data = dataItem
	return c.JSON(response)
}

// Find an item ================================================================================
// @Tags Item
// @Summary Find an item
// @Description Find an item
// @Param id path string true "Item ID"
// @Accept  json
// @Produce  json
// @Success 200 {object} config.DataResponse "desc"
// @Router /v1/items/{id} [get]
// @Security ApiKeyAuth
// @Security ApiTokenAuth
func HandleFindItem(c *fiber.Ctx) error {
	var response config.DataResponse
	var dataItem domainmdl.ToDoItem

	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		response.Message = err.Error()
		return c.JSON(response)
	}

	dataItem.Id = id

	if err := setup(c).FindItem(c, &dataItem); err != nil {
		response.Message = err.Error()
		return c.JSON(response)
	}

	response.Status = true
	response.Data = &dataItem
	return c.JSON(response)
}

// Find all ================================================================================
// @Tags Item
// @Summary Find all
// @Description Find all
// @Accept  json
// @Produce  json
// @Success 200 {object} config.DataResponse "desc"
// @Router /v1/items [get]
// @Security ApiKeyAuth
// @Security ApiTokenAuth
func HandleFindAll(c *fiber.Ctx) error {
	var response config.DataResponse
	var dataItem []domainmdl.ToDoItem

	if err := setup(c).FindAll(c, &dataItem); err != nil {
		response.Message = err.Error()
		return c.JSON(response)
	}

	response.Status = true
	response.Data = &dataItem
	return c.JSON(response)
}

// Update an item ================================================================================
// @Tags Item
// @Summary Update an item
// @Description Update an item
// @Param id path string true "Item ID"
// @Param body body domainmdl.UpdateItem true " "
// @Accept  json
// @Produce  json
// @Success 200 {object} config.DataResponse "desc"
// @Router /v1/items/{id} [put]
// @Security ApiKeyAuth
// @Security ApiTokenAuth
func HandleEditItem(c *fiber.Ctx) error {
	var response config.DataResponse
	var dataItem domainmdl.ToDoItem

	if err := c.BodyParser(&dataItem); err != nil {
		response.Message = err.Error()
		return c.JSON(response)
	}

	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return err
	}

	dataItem.Id = id

	if err := setup(c).UpdateItem(c, &dataItem); err != nil {
		response.Message = err.Error()
		return c.JSON(response)
	}

	response.Status = true
	response.Data = dataItem
	return c.JSON(response)
}

// Delete an item ================================================================================
// @Tags Item
// @Summary Delete an item
// @Description Delete an item
// @Param id path string true "Item ID"
// @Accept  json
// @Produce  json
// @Success 200 {object} config.DataResponse "desc"
// @Router /v1/items/{id} [delete]
// @Security ApiKeyAuth
// @Security ApiTokenAuth
func HandleDeleteItem(c *fiber.Ctx) error {
	var response config.DataResponse
	var dataItem domainmdl.ToDoItem

	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return err
	}

	dataItem.Id = id

	if err := setup(c).DeleteItem(c, &dataItem); err != nil {
		response.Message = err.Error()
		return c.JSON(response)
	}

	response.Status = true
	response.Data = dataItem
	return c.JSON(response)
}
