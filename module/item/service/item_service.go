package todoservice

import (
	domainifc "todo/domain/interface"
	domainmdl "todo/domain/model"

	"github.com/gofiber/fiber/v2"
)

type ToDoBiz struct {
	store domainifc.TodoItemStorage
}

func ToDoItemBiz(store domainifc.TodoItemStorage) *ToDoBiz {
	return &ToDoBiz{store: store}
}

func (biz *ToDoBiz) CreateItem(ctx *fiber.Ctx, data *domainmdl.ToDoItem) error {
	// do not allow "finished" status when creating a new task
	data.Status = "Doing" // set to default
	if err := biz.store.CreateItem(ctx, data); err != nil {
		return err
	}
	return nil
}

func (biz *ToDoBiz) FindItem(ctx *fiber.Ctx, data *domainmdl.ToDoItem) error {
	if err := biz.store.FindItem(ctx, data); err != nil {
		return err
	}
	return nil
}

func (biz *ToDoBiz) FindAll(ctx *fiber.Ctx, data *[]domainmdl.ToDoItem) error {
	if err := biz.store.FindAll(ctx, data); err != nil {
		return err
	}
	return nil
}

func (biz *ToDoBiz) UpdateItem(ctx *fiber.Ctx, data *domainmdl.ToDoItem) error {
	if err := biz.store.UpdateItem(ctx, data); err != nil {
		return err
	}
	return nil
}

func (biz *ToDoBiz) DeleteItem(ctx *fiber.Ctx, data *domainmdl.ToDoItem) error {
	if err := biz.store.DeleteItem(ctx, data); err != nil {
		return err
	}
	return nil
}
