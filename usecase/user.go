package usecase

import (
	"app/domain/models"
	"app/domain/repository"
	"app/graph/model"
	"app/infrastructure/auth"
	"context"
	"log"
)

type userUsecase struct {
	userRepository  repository.UserRepository
	imageRepository repository.ImageRepository
}

type UserUsecase interface {
	Get(id int) (*models.User, error)
	Create(input model.CreateUserInput) (*models.User, error)
	Update(id int, input model.UpdateUserInput) (*models.User, error)
}

func NewUserUsecase(
	u repository.UserRepository,
	i repository.ImageRepository,
) UserUsecase {
	return &userUsecase{
		u,
		i,
	}
}

func (u *userUsecase) Get(id int) (*models.User, error) {
	return u.userRepository.Get(id)
}

func (u *userUsecase) Create(input model.CreateUserInput) (*models.User, error) {
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

func (u *userUsecase) Update(id int, input model.UpdateUserInput) (*models.User, error) {
	var err error
	user := &models.User{
		ID:           id,
		DisplayName:  *input.DisplayName,
		WebURL:       *input.WebURL,
		Introduction: *input.Introduction,
	}

	if input.Icon != nil {
		user.IconPath, err = u.imageRepository.Upload("icons", input.Icon)
		if err != nil {
			return nil, err
		}
	}

	if input.Cover != nil {
		user.CoverPath, err = u.imageRepository.Upload("covers", input.Cover)
		if err != nil {
			return nil, err
		}
	}

	return u.userRepository.Update(user)
}
