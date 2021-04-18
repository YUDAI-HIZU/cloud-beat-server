package repository

import (
	"app/domain/models"
	"context"
)

type GenreRepository interface {
	List(ctx context.Context) ([]*models.Genre, error)
}
