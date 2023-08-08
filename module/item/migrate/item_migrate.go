package todomigrate

import (
	"todo/database"
	todomodel "todo/module/item/model"
)

func AutoMigrate() {
	db := database.DB
	db.AutoMigrate(&todomodel.ToDoItem{})
}
