package repository

import (
	"app/domain/models"
	"context"
)

type AudioRepository interface {
	Upload(ctx context.Context, audio *models.Audio) error
	Delete(ctx context.Context, objName string) error
}
