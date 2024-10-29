package storage

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var msqldb *gorm.DB

type MysqlConf struct {
	Host         string
	Port         string
	User         string
	DatabaseName string
	Password     string
}

func (msql MysqlConf) getDSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		msql.User,
		msql.Password,
		msql.Host,
		msql.Port,
		msql.DatabaseName,
	)
}

func GetMysqlClient() *gorm.DB {
	mysqlConf := MysqlConf{
		Host:         os.Getenv("msql_host"),
		Port:         os.Getenv("mysql_port"),
		User:         os.Getenv("mysql_user"),
		Password:     os.Getenv("mysql_password"),
		DatabaseName: os.Getenv("mysql_db_name"),
	}

	dsn := mysqlConf.getDSN()
	mysqlDb, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}))

	if err != nil {
		panicMsg := fmt.Sprintf("Failed to connect to mysql database %v", err)
		panic(panicMsg)
	}
	fmt.Println("Mysql Connection Successful")

	return mysqlDb
}
