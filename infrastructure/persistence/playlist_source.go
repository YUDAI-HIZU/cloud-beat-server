package persistence

import (
	"app/domain/models"
	"app/domain/repository"

	"github.com/jinzhu/gorm"
)

type playlistSourcePersistence struct {
	db *gorm.DB
}

func NewPlaylistSourcePersistence(db *gorm.DB) repository.PlaylistSourceRepository {
	return &playlistSourcePersistence{
		db,
	}
}

type User struct {
	Name string
}

func (p *playlistSourcePersistence) Create(userID int, playlistSource *models.PlaylistSource) (*models.PlaylistSource, error) {
	if err := p.db.Create(playlistSource).Error; err != nil {
		return nil, err
	}
	return playlistSource, nil
}

func (p *playlistSourcePersistence) BatchDelete(userID int, playlistID int, trackIDs []int) ([]*models.PlaylistSource, error) {
	playlistSources := make([]*models.PlaylistSource, 0)
	if err := p.db.Where("playlist_id = ? AND track_id IN (?)", playlistID, trackIDs).Find(&playlistSources).Delete(&models.PlaylistSource{}).Error; err != nil {
		return nil, err
	}
	return playlistSources, nil
}
