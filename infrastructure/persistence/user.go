package persistence

import (
	"app/domain/models"
	"app/domain/repository"
	"context"

	"github.com/jinzhu/gorm"
)

type userPersistence struct {
	db *gorm.DB
}

func NewUserPersistence(db *gorm.DB) repository.UserRepository {
	return &userPersistence{
		db,
	}
}

func (p *userPersistence) Get(ctx context.Context, id int) (*models.User, error) {
	user := &models.User{}
	if err := p.db.First(user, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, err
		}
		return nil, err
	}
	return user, nil
}

func (p *userPersistence) Create(ctx context.Context, user *models.User) (*models.User, error) {
	if err := p.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (p *userPersistence) Update(ctx context.Context, user *models.User) (*models.User, error) {
	if err := p.db.Update(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
