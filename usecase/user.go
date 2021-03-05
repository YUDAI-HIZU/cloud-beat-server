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
	userRepository repository.UserRepository
}

type UserUseCase interface {
	Create(input model.CreateUserInput) (*models.User, error)
	Update(id int, input model.UpdateUserInput) (*models.User, error)
	GetByID(id int) (*models.User, error)
}

func NewUserUseCase(u repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: u,
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
	user := &models.User{
		ID:           id,
		DisplayName:  *input.DisplayName,
		WebURL:       *input.WebURL,
		Introduction: *input.Introduction,
	}

	user, err := u.userRepository.Update(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u userUseCase) GetByID(id int) (*models.User, error) {
	user, err := u.userRepository.GetByID(id)

	if err != nil {
		return nil, err
	}

	return user, nil
}
