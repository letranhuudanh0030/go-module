package module

import (
	todoMigrate "todo/module/item/migrate"
)

func AutoMigrate() {
	todoMigrate.AutoMigrate()
}
