package database

import (
	"app/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func InitDB() *gorm.DB {
	conn, err := gorm.Open("mysql", config.DatabaseURL)
	if err != nil {
		panic(err)
	}
	return conn
}
