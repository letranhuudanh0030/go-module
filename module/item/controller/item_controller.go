package todoctrl

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"

	"todo/database"
	todomodel "todo/module/item/model"
	todobiz "todo/module/item/repository"
)

func HanleCreateItem(c *fiber.Ctx) error {
	var dataItem todomodel.ToDoItem

	if err := c.BodyParser(&dataItem); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// preprocess title - trim all spaces
	dataItem.Title = strings.TrimSpace(dataItem.Title)

	// setup dependencies
	storage := todomodel.NewPostgreSQLStorage(database.DB)
	biz := todobiz.ToDoItemBiz(storage)

	if err := biz.CreateNewItem(c.Context(), &dataItem); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"data": dataItem.Id})
}

func HandleFindItem(c *fiber.Ctx) error {
	// var dataItem *todomodel.ToDoItem

	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.JSON(fiber.Map{"error": err.Error()})
	}

	// setup dependencies
	storage := todomodel.NewPostgreSQLStorage(database.DB)
	biz := todobiz.ToDoItemBiz(storage)

	dataItem, err := biz.FindItem(c.Context(), id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"data": &dataItem})
}

func HandleFindAll(c *fiber.Ctx) error {
	// setup dependencies
	storage := todomodel.NewPostgreSQLStorage(database.DB)
	biz := todobiz.ToDoItemBiz(storage)

	dataItem, err := biz.FindAll(c.Context())
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"data": &dataItem})
}
