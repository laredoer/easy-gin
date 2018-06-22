package database

import (
	"github.com/gohouse/gorose"        //import Gorose
	_ "github.com/go-sql-driver/mysql" //import DB driver
	"fmt"
)

// DB Config.(Recommend to use configuration file to import)
var DbConfig = map[string]interface{}{
	// Default database configuration
	"Default": "mysql_dev",
	// (Connection pool) Max open connections, default value 0 means unlimit.
	"SetMaxOpenConns": 300,
	// (Connection pool) Max idle connections, default value is 1.
	"SetMaxIdleConns": 10,

	// Define the database configuration character "mysql_dev".
	"Connections":map[string]map[string]string{
		"mysql_dev": map[string]string{
			"host":     "192.168.197.130",
			"username": "root",
			"password": "root",
			"port":     "3306",
			"database": "golang",
			"charset":  "utf8",
			"protocol": "tcp",
			"prefix":   "",      // Table prefix
			"driver":   "mysql", // Database driver(mysql,sqlite,postgres,oracle,mssql)
		},
	},
}

var DB *gorose.Database

func init() {
	connection, err := gorose.Open(DbConfig)
	if err != nil {
		fmt.Println(err)
		return
	}
	// close DB
	//defer connection.Close()

	DB = connection.GetInstance()

}