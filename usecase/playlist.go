package usecase

import (
	"app/domain/models"
	"app/domain/repository"
	"app/graph/model"
)

type playlistUsecase struct {
	playlistRepository repository.PlaylistRepository
}

type PlaylistUsecase interface {
	ListByUserID(userID int) ([]*models.Playlist, error)
	Create(userID int, input model.CreatePlaylistInput) (*models.Playlist, error)
	Delete(userID int, input model.DeletePlaylistInput) (*models.Playlist, error)
}

func NewPlaylistUsecase(p repository.PlaylistRepository) PlaylistUsecase {
	return &playlistUsecase{
		p,
	}
}

func (u *playlistUsecase) ListByUserID(userID int) ([]*models.Playlist, error) {
	return u.playlistRepository.ListByUserID(userID)
}

func (u *playlistUsecase) Create(userID int, input model.CreatePlaylistInput) (*models.Playlist, error) {
	playlist := &models.Playlist{
		UserID: userID,
		Title:  input.Title,
		Public: input.Public,
	}
	for _, trackID := range input.TrackIDs {
		playlist.PlaylistSources = append(
			playlist.PlaylistSources,
			&models.PlaylistSource{
				TrackID: trackID,
			},
		)
	}
	return u.playlistRepository.Create(playlist)
}

func (u *playlistUsecase) Delete(userID int, input model.DeletePlaylistInput) (*models.Playlist, error) {
	return u.playlistRepository.Delete(input.ID, userID)
}
