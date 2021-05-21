package usecase

import (
	"app/domain/models"
	"app/domain/repository"
	"context"
)

type trackUsecase struct {
	trackRepository repository.TrackRepository
	imageRepository repository.ImageRepository
	AudioRepository repository.AudioRepository
}

type TrackUsecase interface {
	List(ctx context.Context) ([]*models.Track, error)
	ListByUserID(ctx context.Context, userID string) ([]*models.Track, error)
	Get(ctx context.Context, id int) (*models.Track, error)
	Create(ctx context.Context, track *models.Track, image *models.Image, audio *models.Audio) (*models.Track, error)
	Delete(ctx context.Context, track *models.Track) (*models.Track, error)
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

func (u *trackUsecase) List(ctx context.Context) ([]*models.Track, error) {
	return u.trackRepository.List(ctx)
}

func (u *trackUsecase) ListByUserID(ctx context.Context, userID string) ([]*models.Track, error) {
	return u.trackRepository.ListByUserID(ctx, userID)
}

func (u *trackUsecase) Get(ctx context.Context, id int) (*models.Track, error) {
	return u.trackRepository.Get(ctx, id)
}

func (u *trackUsecase) Create(ctx context.Context, track *models.Track, image *models.Image, audio *models.Audio) (*models.Track, error) {
	if err := track.Validation(); err != nil {
		return nil, err
	}

	if err := resizeImage(image); err != nil {
		return nil, err
	}

	track, err := u.trackRepository.Create(ctx, track)
	if err != nil {
		return nil, err
	}

	if err := u.imageRepository.Upload(ctx, image); err != nil {
		return nil, err
	}

	if err := u.AudioRepository.Upload(ctx, audio); err != nil {
		return nil, err
	}

	return track, nil
}

func (u *trackUsecase) Delete(ctx context.Context, track *models.Track) (*models.Track, error) {
	track, err := u.trackRepository.Delete(ctx, track)
	if err != nil {
		return nil, err
	}

	if err := u.imageRepository.Delete(ctx, track.ThumbName); err != nil {
		return nil, err
	}

	if err := u.AudioRepository.Delete(ctx, track.AudioName); err != nil {
		return nil, err
	}

	return track, nil
}
