package todomigrate

import (
	"todo/database"
	todomodel "todo/modules/item/model"
)

func AutoMigrate() {
	db := database.DB
	db.AutoMigrate(&todomodel.ToDoItem{})
}
