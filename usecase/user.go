package usecase

import (
	"app/domain/models"
	"app/domain/repository"
	"context"

	"github.com/google/uuid"
)

type userUsecase struct {
	userRepository  repository.UserRepository
	imageRepository repository.ImageRepository
}

type UserUsecase interface {
	Get(ctx context.Context, id string) (*models.User, error)
	Create(ctx context.Context, user *models.User) (*models.User, error)
	Update(ctx context.Context, user *models.User, image *models.Image) (*models.User, error)
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

func (u *userUsecase) Get(ctx context.Context, id string) (*models.User, error) {
	return u.userRepository.Get(ctx, id)
}

func (u *userUsecase) Create(ctx context.Context, user *models.User) (*models.User, error) {
	if err := user.Validation(); err != nil {
		return nil, err
	}

	user, err := u.userRepository.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userUsecase) Update(ctx context.Context, user *models.User, image *models.Image) (*models.User, error) {
	user.IconName = uuid.New().String()

	if err := user.Validation(); err != nil {
		return nil, err
	}

	if err := resizeImage(image); err != nil {
		return nil, err
	}

	user, err := u.userRepository.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	if err := u.imageRepository.Upload(ctx, image); err != nil {
		return nil, err
	}

	return user, nil
}
