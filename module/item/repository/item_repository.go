package todobiz

import (
	"context"
	"errors"
	todomodel "todo/module/item/model"
)

type TodoItemStorage interface {
	CreateItem(ctx context.Context, data *todomodel.ToDoItem) error
	FindItem(ctx context.Context, id int) (*todomodel.ToDoItem, error)
	FindAll(ctx context.Context) ([]*todomodel.ToDoItem, error)
}

type ToDoBiz struct {
	store TodoItemStorage
}

func ToDoItemBiz(store TodoItemStorage) *ToDoBiz {
	return &ToDoBiz{store: store}
}

func (biz *ToDoBiz) CreateNewItem(ctx context.Context, data *todomodel.ToDoItem) error {
	if data.Title == "" {
		return errors.New("title can not be blank")
	}

	// do not allow "finished" status when creating a new task
	data.Status = "Doing" // set to default

	if err := biz.store.CreateItem(ctx, data); err != nil {
		return err
	}

	return nil
}

func (biz *ToDoBiz) FindItem(ctx context.Context, id int) (*todomodel.ToDoItem, error) {
	if id == 0 {
		return nil, errors.New("id can not be blank")
	}

	item, err := biz.store.FindItem(ctx, id)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (biz *ToDoBiz) FindAll(ctx context.Context) ([]*todomodel.ToDoItem, error) {
	items, err := biz.store.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return items, nil
}
