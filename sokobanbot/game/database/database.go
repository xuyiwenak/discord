package database

import (
	"database/sql"
	"fmt"
)

const mysql_hostname = "127.0.0.1"
const mysql_port = 3306
const mysql_database = "sokobot"
const mysql_username = "root"
const mysql_password = "123456"
const mysql_connection = "tcp"

type Database struct {
	DBType int
}
type DatabaseFun interface {
}

var (
	DB  *sql.DB
	err error
)

func init() {
	dbConfig := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", mysql_username, mysql_password, mysql_connection, mysql_hostname, mysql_port, mysql_database)
	DB, err = sql.Open("mysql", dbConfig)
	if err != nil {
		panic(err.Error())
	}
}
