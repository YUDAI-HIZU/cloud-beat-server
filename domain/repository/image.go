package repository

import (
	"app/domain/models"
	"app/graph/model"
)

type ImageRepository interface {
	Create(image *models.Image) (*models.Image, error)
	Upload(id int, name string, input model.CreateImageInput) error
}
