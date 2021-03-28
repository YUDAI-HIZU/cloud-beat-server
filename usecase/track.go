package usecase

import (
	"app/domain/models"
	"app/domain/repository"
	"app/graph/model"
)

type trackUsecase struct {
	trackRepository repository.TrackRepository
	imageRepository repository.ImageRepository
	AudioRepository repository.AudioRepository
}

type TrackUsecase interface {
	List() ([]*models.Track, error)
	ListByUserID(userID int) ([]*models.Track, error)
	Get(id int) (*models.Track, error)
	Create(userID int, input model.CreateTrackInput) (*models.Track, error)
	Update(userID int, input model.UpdateTrackInput) (*models.Track, error)
}

func NewTrackUsecase(
	t repository.TrackRepository,
	i repository.ImageRepository,
	a repository.AudioRepository,
) TrackUsecase {
	return &trackUsecase{
		t,
		i,
		a,
	}
}

func (u *trackUsecase) List() ([]*models.Track, error) {
	return u.trackRepository.List()
}

func (u *trackUsecase) ListByUserID(userID int) ([]*models.Track, error) {
	return u.trackRepository.ListByUserID(userID)
}

func (u *trackUsecase) Get(id int) (*models.Track, error) {
	return u.trackRepository.Get(id)
}

func (u *trackUsecase) Create(userID int, input model.CreateTrackInput) (*models.Track, error) {
	var err error
	track := &models.Track{
		Title:       input.Title,
		Description: input.Description,
		YoutubeLink: *input.YoutubeLink,
		GenreID:     input.GenreID,
		UserID:      userID,
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

func (u *trackUsecase) Update(userID int, input model.UpdateTrackInput) (*models.Track, error) {
	var err error
	track := &models.Track{
		Title:       input.Title,
		Description: input.Description,
		YoutubeLink: *input.YoutubeLink,
		GenreID:     input.GenreID,
		UserID:      userID,
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
