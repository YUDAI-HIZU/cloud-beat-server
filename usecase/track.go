package usecase

import (
	"app/domain/models"
	"app/domain/repository"
	"app/graph/model"
)

type trackUseCase struct {
	trackRepository repository.TrackRepository
	imageRepository repository.ImageRepository
	AudioRepository repository.AudioRepository
}

type TrackUseCase interface {
	List() ([]*models.Track, error)
	Create(userID int, input model.CreateTrackInput) (*models.Track, error)
	Update(userID int, input model.UpdateTrackInput) (*models.Track, error)
}

func NewTrackUseCase(
	t repository.TrackRepository,
	i repository.ImageRepository,
	a repository.AudioRepository,
) TrackUseCase {
	return &trackUseCase{
		t,
		i,
		a,
	}
}

func (u trackUseCase) List() ([]*models.Track, error) {
	return u.trackRepository.List()
}

func (u trackUseCase) Create(userID int, input model.CreateTrackInput) (*models.Track, error) {
	var err error
	track := &models.Track{
		Title:  input.Title,
		UserID: userID,
	}

	if input.Thumbnail != nil {
		track.ThumbnailPath, err = u.imageRepository.Upload("thumbnails", input.Thumbnail)
		if err != nil {
			return nil, err
		}
	}

	track.SoundPath, err = u.AudioRepository.Upload("sounds", &input.Sound)
	if err != nil {
		return nil, err
	}

	return u.trackRepository.Create(track)
}

func (u trackUseCase) Update(userID int, input model.UpdateTrackInput) (*models.Track, error) {
	var err error
	track := &models.Track{
		Title:  input.Title,
		UserID: userID,
	}

	if input.Thumbnail != nil {
		track.ThumbnailPath, err = u.imageRepository.Upload("thumbnails", input.Thumbnail)
		if err != nil {
			return nil, err
		}
	}

	if input.Sound != nil {
		track.SoundPath, err = u.AudioRepository.Upload("sounds", input.Sound)
		if err != nil {
			return nil, err
		}
	}

	return u.trackRepository.Update(track)
}
