package repository

import (
	"app/domain/models"
	"context"
)

type ImageRepository interface {
	Upload(ctx context.Context, image *models.Image) error
	Delete(ctx context.Context, objName string) error
}
