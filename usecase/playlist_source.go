package usecase

import (
	"app/domain/models"
	"app/domain/repository"
	"app/graph/model"
)

type playlistSourceUsecase struct {
	playlistSourceRepository repository.PlaylistSourceRepository
}

type PlaylistSourceUsecase interface {
	Create(userID int, input model.CreatePlaylistSourceInput) (*models.PlaylistSource, error)
	BatchDelete(userID int, input model.DeletePlaylistSourceInput) ([]*models.PlaylistSource, error)
}

func NewPlaylistSourceUsecase(p repository.PlaylistSourceRepository) PlaylistSourceUsecase {
	return &playlistSourceUsecase{
		p,
	}
}

func (u *playlistSourceUsecase) Create(userID int, input model.CreatePlaylistSourceInput) (*models.PlaylistSource, error) {
	playlistSource := &models.PlaylistSource{
		PlaylistID: input.PlaylistID,
		TrackID:    input.TrackID,
	}
	return u.playlistSourceRepository.Create(userID, playlistSource)
}

func (u *playlistSourceUsecase) BatchDelete(userID int, input model.DeletePlaylistSourceInput) ([]*models.PlaylistSource, error) {
	return u.playlistSourceRepository.BatchDelete(userID, input.PlaylistID, input.TrackIDs)
}
