package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log")

var SqlDB *sql.DB

func init() {
	var err error
	SqlDB, err = sql.Open("mysql", "root:root@tcp(192.168.197.130:3306)/golang?parseTime=true")
	if err != nil {
		log.Fatal(err.Error())
	}
	err = SqlDB.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
}
