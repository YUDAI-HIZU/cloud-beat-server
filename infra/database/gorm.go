package database

import (
	"app/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

func InitDB() *gorm.DB {
	conn, err := gorm.Open("mysql", config.DatabaseURL)
	if err != nil {
		logrus.Error(err)
	}
	return conn
}
