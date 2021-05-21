package repository

import (
	"app/domain/models"
	"context"
)

type PlaylistRepository interface {
	Get(ctx context.Context, id int) (*models.Playlist, error)
	ListByUserID(ctx context.Context, userID string) ([]*models.Playlist, error)
	Create(ctx context.Context, playlist *models.Playlist) (*models.Playlist, error)
	Delete(ctx context.Context, playlist *models.Playlist) (*models.Playlist, error)
}
