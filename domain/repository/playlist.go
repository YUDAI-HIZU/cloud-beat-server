package repository

import "app/domain/models"

type PlaylistRepository interface {
	Get(id int) (*models.Playlist, error)
	ListByUserID(userID int) ([]*models.Playlist, error)
	Create(playlist *models.Playlist) (*models.Playlist, error)
	Delete(id int, userID int) (*models.Playlist, error)
}
