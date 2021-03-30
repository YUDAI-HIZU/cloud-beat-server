package usecase

import (
	"app/domain/models"
	"app/domain/repository"
	"app/graph/model"
)

type musicVideoUsecase struct {
	musicVideoRepository repository.MusicVideoRepository
	videoRepository      repository.VideoRepository
}

type MusicVideoUsecase interface {
	Create(userID int, input model.CreateMusicVideoInput) (*models.MusicVideo, error)
	Delete(userID int, input model.DeleteMusicVideoInput) (*models.MusicVideo, error)
}

func NewMusicVideoUsecase(
	m repository.MusicVideoRepository,
	v repository.VideoRepository,
) MusicVideoUsecase {
	return &musicVideoUsecase{
		m,
		v,
	}
}

func (u *musicVideoUsecase) Create(userID int, input model.CreateMusicVideoInput) (*models.MusicVideo, error) {
	var err error
	musicVideo := &models.MusicVideo{
		Text:   input.Text,
		UserID: userID,
	}

	musicVideo.VideoPath, err = u.videoRepository.Upload("musicVideos", &input.Video)
	if err != nil {
		return nil, err
	}

	return u.musicVideoRepository.Create(musicVideo)
}

func (u *musicVideoUsecase) Delete(userID int, input model.DeleteMusicVideoInput) (*models.MusicVideo, error) {
	return u.musicVideoRepository.Delete(input.ID, userID)
}
