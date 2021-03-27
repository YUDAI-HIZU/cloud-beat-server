package repository

import (
	"app/domain/models"
)

type UserRepository interface {
	Get(id int) (*models.User, error)
	Create(user *models.User) (*models.User, error)
	Update(user *models.User) (*models.User, error)
}
