package domainifc

import (
	domainmdl "todo/domain/model"

	"github.com/gofiber/fiber/v2"
)

type TodoItemStorage interface {
	CreateItem(ctx *fiber.Ctx, data *domainmdl.ToDoItem) error
	FindItem(ctx *fiber.Ctx, data *domainmdl.ToDoItem) error
	FindAll(ctx *fiber.Ctx, data *[]domainmdl.ToDoItem) error
	UpdateItem(ctx *fiber.Ctx, data *domainmdl.ToDoItem) error
	DeleteItem(ctx *fiber.Ctx, data *domainmdl.ToDoItem) error
}
