package todomigrate

import (
	"todo/database"
	doaminmdl "todo/domain/model"
)

func AutoMigrate() {
	db := database.DB
	db.AutoMigrate(&doaminmdl.ToDoItem{})
}
