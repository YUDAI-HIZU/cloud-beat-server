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

func (p *userPersistence) Create(user *models.User) (*models.User, error) {
	if err := p.db.Model(user).Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (p *userPersistence) Update(user *models.User) (*models.User, error) {
	if err := p.db.Model(user).Update(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (p *userPersistence) GetByID(id int) (*models.User, error) {
	var user models.User
	if err := p.db.Model(user).Take(&user, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, err
		}
		return nil, err
	}
	return &user, nil
}
