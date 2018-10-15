package migration

import (
	"figin/system"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

var db *gorm.DB

func Migrate() {
	var err error
	if config := system.Config(); config.SqlConnect != "" {
		// 连接数据库
		db, err = gorm.Open("mysql", config.SqlConnect)
		if err != nil {
			panic(err)
		}
		// 执行迁移
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Test{})
	}
}