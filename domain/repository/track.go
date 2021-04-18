package repository

import (
	"app/domain/models"
	"context"
)

type TrackRepository interface {
	List(ctx context.Context) ([]*models.Track, error)
	ListByUserID(ctx context.Context, userID int) ([]*models.Track, error)
	Get(ctx context.Context, id int) (*models.Track, error)
	Create(ctx context.Context, track *models.Track) (*models.Track, error)
	Delete(ctx context.Context, track *models.Track) (*models.Track, error)
}
