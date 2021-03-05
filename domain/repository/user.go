package repository

import (
	"app/domain/models"
)

type UserRepository interface {
	Create(user *models.User) (*models.User, error)
	Update(user *models.User) (*models.User, error)
	GetByID(id int) (*models.User, error)
}
