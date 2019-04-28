package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"os"
	"reflect"
	"sync"
	"time"
)

const (
	default_addr     = "127.0.0.1:3306"
	default_user     = "root"
	default_password = "123456"
	default_DB       = "default"
)

var mysql_DB_Pools = sync.Pool{
	New: func() interface{} {
		db, err := sql.Open("mysql", newConfig())
		if err != nil {
			panic(err.Error())
		}
		return db
	},
}

func main() {
	db1 := mysql_DB_Pools.Get().(*sql.DB)
	defer mysql_DB_Pools.Put(db1)

	err1 := db1.Ping()
	fmt.Println(err1)
}

func newConfig() string {
	mysqlConfig := mysql.Config{
		Net:                  "tcp",
		Params:               map[string]string{"charset": "utf8"},
		Collation:            "utf8_general_ci",
		Loc:                  time.UTC,
		AllowNativePasswords: true,
	}
	config := reflect.TypeOf(mysqlConfig)
	for i := 0; i < config.FieldAlign(); i++ {
		if config.Field(i).Name == "User" {
			if os.Getenv("User") == "" {
				mysqlConfig.User = default_user
			} else {
				mysqlConfig.User = os.Getenv("User")
			}
		} else if config.Field(i).Name == "Passwd" {
			if os.Getenv("Passwd") == "" {
				mysqlConfig.Passwd = default_password
			} else {
				mysqlConfig.Passwd = os.Getenv("Passwd")
			}
		} else if config.Field(i).Name == "Addr" {
			if os.Getenv("Addr") == "" {
				mysqlConfig.Addr = default_addr
			} else {
				mysqlConfig.Addr = os.Getenv("Addr")
			}
		} else if config.Field(i).Name == "DBName" {
			if os.Getenv("DBName") == "" {
				mysqlConfig.DBName = default_DB
			} else {
				mysqlConfig.DBName = os.Getenv("DBName")
			}
		}
	}
	fmt.Printf("The db is %s \n", mysqlConfig.FormatDSN())
	return mysqlConfig.FormatDSN()
}
