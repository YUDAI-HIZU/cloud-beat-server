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
	img := &models.Image{
		Name:      uuid.New().String() + input.Image.Filename,
		OwnerID:   id,
		OwnerType: input.OwnerType,
	}

	img, err := i.imageRepository.Create(img)
	if err != nil {
		return nil, err
	}

	if err := i.imageRepository.Upload(img.ID, img.Name, input); err != nil {
		return nil, err
	}

	return img, nil
}
