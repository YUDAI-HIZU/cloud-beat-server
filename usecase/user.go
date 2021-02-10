package usecase

import (
	"app/domain/models"
	"app/domain/repository"

	"github.com/jinzhu/gorm"
)

type userUseCase struct {
	userRepository repository.UserRepository
}

type UserUseCase interface {
	Create(db *gorm.DB, user *models.User) error
	GetByID(db *gorm.DB, id int) (*models.User, error)
	GetByUID(db *gorm.DB, uid string) (*models.User, error)
}

func NewUserUseCase(u repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: u,
	}
}

func (u userUseCase) Create(db *gorm.DB, user *models.User) error {
	var err error
	err = u.userRepository.Create(db, user)
	if err != nil {
		return err
	}
	return nil
}

func (u userUseCase) GetByID(db *gorm.DB, id int) (*models.User, error) {
	user, err := u.userRepository.GetByID(db, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u userUseCase) GetByUID(db *gorm.DB, uid string) (*models.User, error) {
	user, err := u.userRepository.GetByUID(db, uid)
	if err != nil {
		return nil, err
	}
	return user, nil
}
