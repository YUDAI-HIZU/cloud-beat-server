package repository

import "app/domain/models"

type TrackRepository interface {
	List() ([]*models.Track, error)
	ListByUserID(userID int) ([]*models.Track, error)
	Get(id int) (*models.Track, error)
	Create(track *models.Track) (*models.Track, error)
	Update(track *models.Track) (*models.Track, error)
	Delete(id int, userID int) (*models.Track, error)
}
