package usecase

import (
	"app/domain/models"
	"app/domain/repository"
	"app/graph/model"
)

type trackUseCase struct {
	trackRepository repository.TrackRepository
	imageRepository repository.ImageRepository
}

type TrackUseCase interface {
	Create(userID int, input model.CreateTrackInput) (*models.Track, error)
	Update(userID int, input model.UpdateTrackInput) (*models.Track, error)
}

func NewTrackUseCase(
	t repository.TrackRepository,
	i repository.ImageRepository,
) TrackUseCase {
	return &trackUseCase{
		t,
		i,
	}
}

func (u trackUseCase) Create(userID int, input model.CreateTrackInput) (*models.Track, error) {
	var err error
	track := &models.Track{
		Title:  input.Title,
		UserID: userID,
	}

	if input.ThumbnailImage != nil {
		track.ThumbnailPath, err = u.imageRepository.Upload("tracks", input.ThumbnailImage)
		if err != nil {
			return nil, err
		}
	}
	return u.trackRepository.Create(track)
}

func (u trackUseCase) Update(userID int, input model.UpdateTrackInput) (*models.Track, error) {
	var err error
	track := &models.Track{
		Title:  input.Title,
		UserID: userID,
	}

	if input.ThumbnailImage != nil {
		track.ThumbnailPath, err = u.imageRepository.Upload("tracks", input.ThumbnailImage)
		if err != nil {
			return nil, err
		}
	}
	return u.trackRepository.Update(track)
}
