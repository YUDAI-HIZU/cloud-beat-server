package usecase

import (
	"app/domain/models"
	"app/domain/repository"
	"context"
)

type playlistSourceUsecase struct {
	playlistSourceRepository repository.PlaylistSourceRepository
}

type PlaylistSourceUsecase interface {
	Create(ctx context.Context, playlistSource *models.PlaylistSource) (*models.PlaylistSource, error)
	Delete(ctx context.Context, playlistSource *models.PlaylistSource) (*models.PlaylistSource, error)
}

func NewPlaylistSourceUsecase(p repository.PlaylistSourceRepository) PlaylistSourceUsecase {
	return &playlistSourceUsecase{
		p,
	}
}

func (u *playlistSourceUsecase) Create(ctx context.Context, playlistSource *models.PlaylistSource) (*models.PlaylistSource, error) {
	return u.playlistSourceRepository.Create(ctx, playlistSource)
}

func (u *playlistSourceUsecase) Delete(ctx context.Context, playlistSource *models.PlaylistSource) (*models.PlaylistSource, error) {
	return u.playlistSourceRepository.Delete(ctx, playlistSource)
}
