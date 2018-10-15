package database

import (
	"database/sql"
)

// 定义DB结构
type DbStruct struct {
	sql *sql.DB
}

// 定义DB连接配置结构
type DbConfig struct {
	SqlConnect string
	MongodbConnect string
}