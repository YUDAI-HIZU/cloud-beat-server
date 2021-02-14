package usecase

import (
	"app/domain/models"
	"app/domain/repository"
	"app/graph/model"
	"app/infrastructure/persistence"
	"fmt"

	"github.com/jinzhu/gorm"
)

type userUseCase struct {
	userRepository repository.UserRepository
}

type UserUseCase interface {
	Create(db *gorm.DB, input model.CreateUserInput) (*models.User, error)
	Update(db *gorm.DB, uid string, input model.UpdateUserInput) (*models.User, error)
	GetByID(db *gorm.DB, id int) (*models.User, error)
	GetByUID(db *gorm.DB, uid string) (*models.User, error)
}

func NewUserUseCase(u repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: u,
	}
}

func (u userUseCase) Create(db *gorm.DB, input model.CreateUserInput) (*models.User, error) {
	user := &models.User{
		UID:         input.UID,
		DisplayName: input.DisplayName,
	}
	user, err := u.userRepository.Create(db, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u userUseCase) Update(db *gorm.DB, uid string, input model.UpdateUserInput) (*models.User, error) {
	user := &models.User{
		UID:          uid,
		DisplayName:  *input.DisplayName,
		WebURL:       *input.WebURL,
		Introduction: *input.Introduction,
	}

	image := persistence.NewImagePersistence("user")
	fmt.Println("=======image upload======", input.IconImage.Filename, input.IconImage.File)
	image.Upload(input.IconImage.Filename, input.IconImage.File)
	image.Upload(input.CoverImage.Filename, input.CoverImage.File)

	user, err := u.userRepository.Update(db, user)
	if err != nil {
		return nil, err
	}
	return user, nil
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
