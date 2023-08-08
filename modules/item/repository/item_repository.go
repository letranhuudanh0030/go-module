package repository

import (
	todomodel "todo/modules/item/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type postgresqlStorage struct {
	db *gorm.DB
}

func NewPostgreSQLStorage(db *gorm.DB) *postgresqlStorage {
	return &postgresqlStorage{db: db}
}

func (s *postgresqlStorage) CreateItem(ctx *fiber.Ctx, data *todomodel.ToDoItem) error {
	if err := s.db.Create(data).Error; err != nil {
		return err
	}

	return nil
}

// FindItem implements todobiz.TodoItemStorage.
func (s *postgresqlStorage) FindItem(ctx *fiber.Ctx, data *todomodel.ToDoItem) error {
	if err := s.db.First(data).Error; err != nil {
		return err
	}

	return nil
}

func (s *postgresqlStorage) FindAll(ctx *fiber.Ctx, data *[]todomodel.ToDoItem) error {
	if err := s.db.Find(data).Error; err != nil {
		return err
	}

	return nil
}

func (s *postgresqlStorage) UpdateItem(ctx *fiber.Ctx, data *todomodel.ToDoItem) error {
	var params todomodel.ToDoItem

	params.Title = data.Title
	params.Status = data.Status

	if err := s.db.First(data).Error; err != nil {
		return err
	}

	if err := s.db.Model(data).Updates(params).Error; err != nil {
		return err
	}

	return nil
}

func (s *postgresqlStorage) DeleteItem(ctx *fiber.Ctx, data *todomodel.ToDoItem) error {

	if err := s.db.First(data).Error; err != nil {
		return err
	}

	if err := s.db.Delete(data).Error; err != nil {
		return err
	}

	return nil
}
