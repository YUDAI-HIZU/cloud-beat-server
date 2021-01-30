package repository

import (
	"app/domain/models"

	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	Create(db *gorm.DB, user *models.User) (*models.User, error)
	GetByID(db *gorm.DB, id int) (*models.User, error)
}