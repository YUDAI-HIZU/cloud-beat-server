package persistence

import (
	"app/domain/models"
	"app/domain/repository"
	"context"

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

func (p *playlistSourcePersistence) Create(ctx context.Context, playlistSource *models.PlaylistSource) (*models.PlaylistSource, error) {
	if err := p.db.Create(playlistSource).Error; err != nil {
		return nil, err
	}
	return playlistSource, nil
}

func (p *playlistSourcePersistence) Delete(ctx context.Context, playlistSource *models.PlaylistSource) (*models.PlaylistSource, error) {
	if err := p.db.Where("playlist_id = ? AND track_id = ?", playlistSource.PlaylistID, playlistSource.TrackID).Find(playlistSource).Delete(playlistSource).Error; err != nil {
		return nil, err
	}
	return playlistSource, nil
}
