package todobiz

import (
	todomodel "todo/modules/item/model"

	"github.com/gofiber/fiber/v2"
)

type TodoItemStorage interface {
	CreateItem(ctx *fiber.Ctx, data *todomodel.ToDoItem) error
	FindItem(ctx *fiber.Ctx, data *todomodel.ToDoItem) error
	FindAll(ctx *fiber.Ctx, data *[]todomodel.ToDoItem) error
	UpdateItem(ctx *fiber.Ctx, data *todomodel.ToDoItem) error
	DeleteItem(ctx *fiber.Ctx, data *todomodel.ToDoItem) error
}

type ToDoBiz struct {
	store TodoItemStorage
}

func ToDoItemBiz(store TodoItemStorage) *ToDoBiz {
	return &ToDoBiz{store: store}
}

func (biz *ToDoBiz) CreateNewItem(ctx *fiber.Ctx, data *todomodel.ToDoItem) error {
	// do not allow "finished" status when creating a new task
	data.Status = "Doing" // set to default
	if err := biz.store.CreateItem(ctx, data); err != nil {
		return err
	}
	return nil
}

func (biz *ToDoBiz) FindItem(ctx *fiber.Ctx, data *todomodel.ToDoItem) error {
	if err := biz.store.FindItem(ctx, data); err != nil {
		return err
	}
	return nil
}

func (biz *ToDoBiz) FindAll(ctx *fiber.Ctx, data *[]todomodel.ToDoItem) error {
	if err := biz.store.FindAll(ctx, data); err != nil {
		return err
	}
	return nil
}

func (biz *ToDoBiz) UpdateItem(ctx *fiber.Ctx, data *todomodel.ToDoItem) error {
	if err := biz.store.UpdateItem(ctx, data); err != nil {
		return err
	}
	return nil
}

func (biz *ToDoBiz) DeleteItem(ctx *fiber.Ctx, data *todomodel.ToDoItem) error {
	if err := biz.store.DeleteItem(ctx, data); err != nil {
		return err
	}
	return nil
}
