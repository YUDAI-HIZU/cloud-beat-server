package usecase

import (
	"app/domain/models"
	"app/domain/repository"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	userRepository repository.UserRepository
}

type UserUseCase interface {
	Create(db *gorm.DB, user *models.User) (*models.User, error)
	GetByID(db *gorm.DB, id int) (*models.User, error)
}

func NewUserUseCase(u repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: u,
	}
}

func (u userUseCase) Create(db *gorm.DB, user *models.User) (*models.User, error) {
	encryptedPassword, err := generateEncryptedPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.EncryptedPassword = encryptedPassword
	user, err = u.userRepository.Create(db, user)
	if err != nil {
		return nil, err
	}
	user.Password = ""
	user.EncryptedPassword = ""
	return user, nil
}

func (u userUseCase) GetByID(db *gorm.DB, id int) (*models.User, error) {
	user, err := u.userRepository.GetByID(db, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func generateEncryptedPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func compareHashAndPassWord(encryptedPassword string, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(password)); err != nil {
		return err
	}
	return nil
}
