package repository

import (
	"app/domain/models"
	"context"
)

type PlaylistSourceRepository interface {
	Create(ctx context.Context, playlistSource *models.PlaylistSource) (*models.PlaylistSource, error)
	Delete(ctx context.Context, playlistSource *models.PlaylistSource) (*models.PlaylistSource, error)
}
