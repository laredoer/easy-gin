package database

import (
	"github.com/gohouse/gorose"        //import Gorose
	_ "github.com/go-sql-driver/mysql" //import DB driver
	"fmt"
	"github.com/kylelemons/go-gypsy/yaml"
	"github.com/astaxie/beego/logs"
)


var DB *gorose.Database

func init() {

	config, err := yaml.ReadFile("conf.yaml")
	if err != nil {
		fmt.Println(err)
	}
	host,err := config.Get("mysql.host")
	if err != nil{
		panic("数据库地址未设置")
	}
	username,err := config.Get("mysql.username")
	if err != nil{
		panic("数据库用户名未配置")
	}
	password,err := config.Get("mysql.password")
	if err != nil{
		panic("数据库密码未配置")
	}
	port,err := config.Get("mysql.port")
	if err != nil{
		logs.Info("数据库端口号未配置,使用默认端口:3306")
		port = "3306"
	}
	database,err := config.Get("mysql.database")
	if err != nil{
		panic("数据库名称未配置")
	}
	prefix,_ := config.Get("mysql.prefix")
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
				"host":     host,
				"username": username,
				"password": password,
				"port":     port,
				"database": database,
				"charset":  "utf8",
				"protocol": "tcp",
				"prefix":   prefix,      // Table prefix
				"driver":   "mysql", // Database driver(mysql,sqlite,postgres,oracle,mssql)
			},
		},
	}

	connection, err := gorose.Open(DbConfig)
	if err != nil {
		fmt.Println(err)
		return
	}
	// close DB
	//defer connection.Close()

	DB = connection.GetInstance()

}