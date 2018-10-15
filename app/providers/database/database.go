package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

var pool DbStruct
var config DbConfig

// 初始化入口
func DbInit(conf *DbConfig) {
	if config = *conf; config.SqlConnect != "" {
		connectSql()
	}
}

// 连接SQL
func connectSql() {
	var err error
	fmt.Println("connect db start")
	pool.sql, err = sql.Open("mysql", config.SqlConnect)
	if err != nil {
		fmt.Println("connect db fail")
		fmt.Println(err)
	}
	defer pool.sql.Close()
}

// 返回DB连接池
func Db(dbType string) *sql.DB {
	if dbType == "sql" {
		// 连接池为空，重新连接
		if config.SqlConnect != "" && pool.sql == nil {
			connectSql()
		}
		return pool.sql
	}
	return nil
}