package repository

import "app/domain/models"

type TrackRepository interface {
	Create(track *models.Track) (*models.Track, error)
	Update(track *models.Track) (*models.Track, error)
}
