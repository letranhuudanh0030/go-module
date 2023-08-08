package todoctrl

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"

	"todo/config"
	"todo/database"
	todomodel "todo/modules/item/model"
	todostorage "todo/modules/item/repository"
	todobiz "todo/modules/item/service"
	"todo/utils"
)

func biz(c *fiber.Ctx) *todobiz.ToDoBiz {
	// setup dependencies
	db := database.DB.Set("username", c.Locals("username"))
	storage := todostorage.NewPostgreSQLStorage(db)
	return todobiz.ToDoItemBiz(storage)
}

// Create a new item ================================================================================
// @Tags Item
// @Summary Create a new item
// @Description Create a new item
// @Param body body todomodel.CreateItem true " "
// @Accept  json
// @Produce  json
// @Success 200 {object} config.DataResponse "desc"
// @Router /v1/items [post]
// @Security ApiKeyAuth
func HanleCreateItem(c *fiber.Ctx) error {
	var response config.DataResponse
	var dataItem todomodel.ToDoItem

	if err := c.BodyParser(&dataItem); err != nil {
		response.Message = config.PARAM_ERROR
		return c.JSON(response)
	}

	if errors := utils.Validator(dataItem); errors != nil {
		response.Message = config.VALIDATE
		response.ValidateError = errors
		return c.JSON(response)
	}

	// pre-process title - trim all spaces
	dataItem.Title = strings.TrimSpace(dataItem.Title)

	if err := biz(c).CreateNewItem(c, &dataItem); err != nil {
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
func HandleFindItem(c *fiber.Ctx) error {
	var response config.DataResponse
	var dataItem todomodel.ToDoItem

	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		response.Message = err.Error()
		return c.JSON(response)
	}

	dataItem.Id = id

	if err := biz(c).FindItem(c, &dataItem); err != nil {
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
func HandleFindAll(c *fiber.Ctx) error {
	var response config.DataResponse
	var dataItem []todomodel.ToDoItem

	if err := biz(c).FindAll(c, &dataItem); err != nil {
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
// @Param body body todomodel.UpdateItem true " "
// @Accept  json
// @Produce  json
// @Success 200 {object} config.DataResponse "desc"
// @Router /v1/items/{id} [put]
// @Security ApiKeyAuth
func HandleEditItem(c *fiber.Ctx) error {
	var response config.DataResponse
	var dataItem todomodel.ToDoItem

	if err := c.BodyParser(&dataItem); err != nil {
		response.Message = err.Error()
		return c.JSON(response)
	}

	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return err
	}

	dataItem.Id = id

	if err := biz(c).UpdateItem(c, &dataItem); err != nil {
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
func HandleDeleteItem(c *fiber.Ctx) error {
	var response config.DataResponse
	var dataItem todomodel.ToDoItem

	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return err
	}

	dataItem.Id = id

	if err := biz(c).DeleteItem(c, &dataItem); err != nil {
		response.Message = err.Error()
		return c.JSON(response)
	}

	response.Status = true
	response.Data = dataItem
	return c.JSON(response)
}
