package todomodel

import (
	"errors"
	"time"

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
	t.LogVersion += 1

	if t.DeletedAt.Valid {
		if err := tx.Model(t).Unscoped().UpdateColumns(map[string]interface{}{
			"log_version": t.LogVersion,
			"deleted_by":  username.(string),
		}).Error; err != nil {
			return err
		}
	}

	return
}
