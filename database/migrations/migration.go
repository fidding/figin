package migration

import (
	"figin/app/providers/database"
)

// Setup 数据迁移启动入口
func Setup() {
	var db = database.GetDB()
	// 执行迁移
	db.Set(
		"gorm:table_options",
		"ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci",
	).AutoMigrate(
		&User{},
	)
}
