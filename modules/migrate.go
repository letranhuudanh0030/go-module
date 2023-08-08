package modules

import (
	todoMigrate "todo/modules/item/migrate"
)

func AutoMigrate() {
	todoMigrate.AutoMigrate()
}
