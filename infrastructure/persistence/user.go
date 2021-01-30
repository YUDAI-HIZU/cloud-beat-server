package persistence

import (
	"app/domain/models"
	"app/domain/repository"

	"github.com/jinzhu/gorm"
)

type userPersistence struct{}

func NewUserPersistence() repository.UserRepository {
	return &userPersistence{}
}

func (u *userPersistence) Create(db *gorm.DB, user *models.User) (*models.User, error) {
	r := db.Create(user)
	if err := r.Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userPersistence) GetByID(db *gorm.DB, id int) (*models.User, error) {
	var user models.User
	if err := db.Take(&user, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, err
		}
		return nil, err
	}
	return &user, nil
}

func (u *userPersistence) GetByEmail(db *gorm.DB, email string) (*models.User, error) {
	var user models.User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, err
		}
		return nil, err
	}
	return &user, nil
}
