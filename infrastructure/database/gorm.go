package database

import (
	"app/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type DatabaseClient struct {
	Client *gorm.DB
}

func NewDatabase() *DatabaseClient {
	client, err := gorm.Open("mysql", config.DatabaseURL)
	if err != nil {
		panic(err)
	}
	return &DatabaseClient{
		Client: client,
	}
}
