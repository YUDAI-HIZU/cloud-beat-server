package repository

import "app/domain/models"

type PlaylistSourceRepository interface {
	Create(userID int, playlistSource *models.PlaylistSource) (*models.PlaylistSource, error)
	BatchDelete(userID int, playlistID int, trackIDs []int) ([]*models.PlaylistSource, error)
}
