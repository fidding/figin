package database

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
	"figin/config"
)

type DataObject struct {
	Conf *config.Configuration
	pool DbStruct
}

func (data *DataObject) Load(conf *config.Configuration) *DataObject {
	// 加载配置
	data.Conf = conf
	// 加载sql连接池
	if data.Conf.SqlConnect != "" {
		data.connectSql()
	}
	return data
}

func (data *DataObject) DbInit() {
	if data.Conf = config.Config(); data.Conf.SqlConnect != "" {
		data.connectSql()
	}
}

func (data *DataObject) connectSql() {
	var err error
	fmt.Println("connect db start")
	data.pool.sql, err = sql.Open("mysql", data.Conf.SqlConnect)
	if err != nil {
		fmt.Println("connect db fail")
		fmt.Println(err)
	}
	defer data.pool.sql.Close()
}

// 返回DB连接池
func (data *DataObject) Db(dbType string) *sql.DB {
	fmt.Println(data)
	if dbType == "sql" {
		// 连接池为空，重新连接
		if data.Conf.SqlConnect != "" && data.pool.sql == nil {
			data.connectSql()
		}
		return data.pool.sql
	}
	return nil
}