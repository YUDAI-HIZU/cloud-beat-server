package usecase

import (
	"app/domain/models"
	"app/domain/repository"
	"context"
)

type playlistUsecase struct {
	playlistRepository repository.PlaylistRepository
}

type PlaylistUsecase interface {
	Get(ctx context.Context, id int) (*models.Playlist, error)
	ListByUserID(ctx context.Context, userID int) ([]*models.Playlist, error)
	Create(ctx context.Context, playlist *models.Playlist) (*models.Playlist, error)
	Delete(ctx context.Context, playlist *models.Playlist) (*models.Playlist, error)
}

func NewPlaylistUsecase(p repository.PlaylistRepository) PlaylistUsecase {
	return &playlistUsecase{
		p,
	}
}

func (u *playlistUsecase) Get(ctx context.Context, id int) (*models.Playlist, error) {
	return u.playlistRepository.Get(ctx, id)
}

func (u *playlistUsecase) ListByUserID(ctx context.Context, userID int) ([]*models.Playlist, error) {
	return u.playlistRepository.ListByUserID(ctx, userID)
}

func (u *playlistUsecase) Create(ctx context.Context, playlist *models.Playlist) (*models.Playlist, error) {
	return u.playlistRepository.Create(ctx, playlist)
}

func (u *playlistUsecase) Delete(ctx context.Context, playlist *models.Playlist) (*models.Playlist, error) {
	return u.playlistRepository.Delete(ctx, playlist)
}
