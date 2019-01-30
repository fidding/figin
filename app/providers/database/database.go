package database

import (
	"figin/config"
	"fmt"

	"github.com/jinzhu/gorm"

	"strconv"

	// _ mysql driver
	_ "github.com/go-sql-driver/mysql"
)

var db *gorm.DB

// GetDB 获取DB实例
func GetDB() *gorm.DB {
	return db
}

// Setup 数据库启动入口
func Setup() {
	var err error
	var conf = config.GetConf()
	var connect = fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.DB.User,
		conf.DB.Password,
		conf.DB.Host+":"+strconv.Itoa(conf.DB.Port),
		conf.DB.DBName,
	)
	// 连接数据库
	db, err = gorm.Open(conf.DB.DBType, connect)
	if err != nil {
		fmt.Printf("db: "+conf.DB.DBType+" connect error %v", err)
	}
}
