package usecase

import (
	"app/domain/models"
	"app/domain/repository"
	"app/graph/model"

	"github.com/google/uuid"
)

type imageUseCase struct {
	imageRepository repository.ImageRepository
}

type ImageUseCase interface {
	Create(id int, input model.CreateImageInput) (*models.Image, error)
}

func NewImageUseCase(i repository.ImageRepository) ImageUseCase {
	return &imageUseCase{
		imageRepository: i,
	}
}

func (i imageUseCase) Create(id int, input model.CreateImageInput) (*models.Image, error) {
	image := &models.Image{
		Name:      uuid.New().String(),
		OwnerID:   id,
		OwnerType: input.OwnerType,
	}

	image, err := i.imageRepository.Create(image)
	if err != nil {
		return nil, err
	}

	if err := i.imageRepository.Upload(image.ID, image.Name, input); err != nil {
		return nil, err
	}

	return image, nil
}
