package todomodel

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type ToDoItem struct {
	Id        int            `json:"id" gorm:"column:id;"`
	Title     string         `json:"title" gorm:"column:title;"`
	Status    string         `json:"status" gorm:"column:status;"`
	CreatedAt *time.Time     `json:"createdAt" gorm:"column:created_at;"`
	UpdatedAt *time.Time     `json:"updatedAt" gorm:"column:updated_at;"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"column:deleted_at;"`
}

func (ToDoItem) TableName() string { return "todo_items" }

type postgresqlStorage struct {
	db *gorm.DB
}

// FindItem implements todobiz.TodoItemStorage.
func (s *postgresqlStorage) FindItem(ctx context.Context, id int) (*ToDoItem, error) {
	var item *ToDoItem
	if err := s.db.Find(&item, id).Error; err != nil {
		return nil, err
	}

	return item, nil
}

func (s *postgresqlStorage) FindAll(ctx context.Context) ([]*ToDoItem, error) {
	var items []*ToDoItem
	if err := s.db.Find(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

func NewPostgreSQLStorage(db *gorm.DB) *postgresqlStorage {
	return &postgresqlStorage{db: db}
}

func (s *postgresqlStorage) CreateItem(ctx context.Context, data *ToDoItem) error {
	if err := s.db.Create(data).Error; err != nil {
		return err
	}

	return nil
}
