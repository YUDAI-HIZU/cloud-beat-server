package repository

import (
	"app/domain/models"
	"app/graph/model"
)

type ImageRepository interface {
	Create(img *models.Image) (*models.Image, error)
	Upload(id int, name string, input model.CreateImageInput) error
}
