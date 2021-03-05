package persistence

import (
	"app/domain/models"
	"app/domain/repository"

	"github.com/jinzhu/gorm"
)

type userPersistence struct {
	db *gorm.DB
}

func NewUserPersistence(db *gorm.DB) repository.UserRepository {
	return &userPersistence{
		db: db,
	}
}

func (u *userPersistence) Create(user *models.User) (*models.User, error) {
	r := u.db.Create(user)
	if err := r.Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userPersistence) Update(user *models.User) (*models.User, error) {
	r := u.db.Update(user)
	if err := r.Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userPersistence) GetByID(id int) (*models.User, error) {
	var user models.User
	if err := u.db.Preload("Icon", "owner_type = ?", models.ImageOwnerTypeUserIcon).Preload("Cover", "owner_type = ?", models.ImageOwnerTypeUserCover).Take(&user, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, err
		}
		return nil, err
	}
	return &user, nil
}
