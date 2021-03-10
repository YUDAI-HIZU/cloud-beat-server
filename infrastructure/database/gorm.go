package database

import (
	"app/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func NewDatabase() *gorm.DB {
	db, err := gorm.Open("mysql", config.DatabaseURL)
	if err != nil {
		panic(err)
	}
	return db
}
