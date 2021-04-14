package usecase

import (
	"app/domain/models"
	"app/domain/repository"
	"app/graph/model"
	"app/usecase/validations"
)

type userUsecase struct {
	userRepository  repository.UserRepository
	imageRepository repository.ImageRepository
	authRepository  repository.AuthRepository
}

type UserUsecase interface {
	Get(id int) (*models.User, error)
	Create(input model.CreateUserInput) (*models.User, error)
	Update(id int, input model.UpdateUserInput) (*models.User, error)
}

func NewUserUsecase(
	u repository.UserRepository,
	i repository.ImageRepository,
	a repository.AuthRepository,
) UserUsecase {
	return &userUsecase{
		u,
		i,
		a,
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

	if err := validations.UserCreateValidation(user); err != nil {
		return nil, err
	}

	user, err := u.userRepository.Create(user)
	if err != nil {
		return nil, err
	}

	if err = u.authRepository.SetIDToClaims(input.UID, user.ID); err != nil {
		return nil, err
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

	return u.userRepository.Update(user)
}
