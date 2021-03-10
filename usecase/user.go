package usecase

import (
	"app/domain/models"
	"app/domain/repository"
	"app/graph/model"
	"app/infrastructure/auth"
	"context"
	"log"
)

type userUseCase struct {
	userRepository  repository.UserRepository
	imageRepository repository.ImageRepository
}

type UserUseCase interface {
	Create(input model.CreateUserInput) (*models.User, error)
	Update(id int, input model.UpdateUserInput) (*models.User, error)
	GetByID(id int) (*models.User, error)
}

func NewUserUseCase(
	u repository.UserRepository,
	i repository.ImageRepository,
) UserUseCase {
	return &userUseCase{
		u,
		i,
	}
}

func (u userUseCase) Create(input model.CreateUserInput) (*models.User, error) {
	user := &models.User{
		UID:         input.UID,
		DisplayName: input.DisplayName,
	}

	user, err := u.userRepository.Create(user)
	if err != nil {
		return nil, err
	}

	auth := auth.NewAuthClient(context.Background())

	if err = auth.SetIDToClaims(input.UID, user.ID); err != nil {
		log.Fatalf("error setting custom claims %v\n", err)
	}

	return user, nil
}

func (u userUseCase) Update(id int, input model.UpdateUserInput) (*models.User, error) {
	var err error
	user := &models.User{
		ID:           id,
		DisplayName:  *input.DisplayName,
		WebURL:       *input.WebURL,
		Introduction: *input.Introduction,
	}

	if input.IconImage != nil {
		user.IconPath, err = u.imageRepository.Upload("icons", input.IconImage)
		if err != nil {
			return nil, err
		}
	}

	if input.CoverImage != nil {
		user.CoverPath, err = u.imageRepository.Upload("covers", input.CoverImage)
		if err != nil {
			return nil, err
		}
	}

	return u.userRepository.Update(user)
}

func (u userUseCase) GetByID(id int) (*models.User, error) {
	return u.userRepository.GetByID(id)
}
