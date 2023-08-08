package todomodel

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type (
	ToDoItem struct {
		Id         int            `json:"id" gorm:"column:id;"`
		Title      string         `json:"title" gorm:"column:title;" validate:"required,min=5"`
		Status     string         `json:"status" gorm:"column:status;"`
		LogVersion int            `json:"logVersion" gorm:"column:log_version;default:0;"`
		CreatedAt  *time.Time     `json:"createdAt" gorm:"column:created_at;"`
		UpdatedAt  *time.Time     `json:"updatedAt" gorm:"column:updated_at;"`
		DeletedAt  gorm.DeletedAt `json:"deletedAt" gorm:"column:deleted_at;"`
		CreatedBy  string         `json:"createdBy" gorm:"column:created_by;"`
		UpdatedBy  string         `json:"updatedBy" gorm:"column:updated_by;"`
		DeletedBy  string         `json:"deletedBy" gorm:"column:deleted_by;"`
	}

	CreateItem struct {
		Title string `json:"title" gorm:"column:title;"`
	}

	UpdateItem struct {
		Title  string `json:"title" gorm:"column:title;"`
		Status string `json:"status" gorm:"column:status;"`
	}
)

func (ToDoItem) TableName() string { return "todo_items" }

func (t *ToDoItem) BeforeCreate(tx *gorm.DB) (err error) {
	username, _ := tx.Get("username")
	t.CreatedBy = username.(string)
	return
}

func (t *ToDoItem) BeforeUpdate(tx *gorm.DB) (err error) {
	username, _ := tx.Get("username")

	if !tx.Statement.Changed("Title", "Status", "DeletedBy") {
		return errors.New("data not change")
	}

	t.LogVersion += 1
	tx.Statement.SetColumn("LogVersion", t.LogVersion)

	t.UpdatedBy = username.(string)
	tx.Statement.SetColumn("UpdatedBy", t.UpdatedBy)

	return
}

func (t *ToDoItem) AfterDelete(tx *gorm.DB) (err error) {
	username, _ := tx.Get("username")

	if t.DeletedAt.Valid {
		if err := tx.Model(t).Unscoped().Updates(map[string]interface{}{
			"deleted_by": username.(string),
		}).Error; err != nil {
			return err
		}
	}

	return
}

type postgresqlStorage struct {
	db *gorm.DB
}

func NewPostgreSQLStorage(db *gorm.DB) *postgresqlStorage {
	return &postgresqlStorage{db: db}
}

func (s *postgresqlStorage) CreateItem(ctx *fiber.Ctx, data *ToDoItem) error {
	if err := s.db.Create(data).Error; err != nil {
		return err
	}

	return nil
}

// FindItem implements todobiz.TodoItemStorage.
func (s *postgresqlStorage) FindItem(ctx *fiber.Ctx, data *ToDoItem) error {
	if err := s.db.First(data).Error; err != nil {
		return err
	}

	return nil
}

func (s *postgresqlStorage) FindAll(ctx *fiber.Ctx, data *[]ToDoItem) error {
	if err := s.db.Find(data).Error; err != nil {
		return err
	}

	return nil
}

func (s *postgresqlStorage) UpdateItem(ctx *fiber.Ctx, data *ToDoItem) error {
	var params ToDoItem

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

func (s *postgresqlStorage) DeleteItem(ctx *fiber.Ctx, data *ToDoItem) error {

	if err := s.db.First(data).Error; err != nil {
		return err
	}

	if err := s.db.Delete(data).Error; err != nil {
		return err
	}

	return nil
}
