package persistence

import (
	"app/domain/models"
	"app/domain/repository"
	"context"

	"github.com/jinzhu/gorm"
)

type playlistPersistence struct {
	db *gorm.DB
}

func NewPlaylistPersistence(db *gorm.DB) repository.PlaylistRepository {
	return &playlistPersistence{
		db,
	}
}

func (p *playlistPersistence) Get(ctx context.Context, id int) (*models.Playlist, error) {
	playlist := &models.Playlist{}
	if err := p.db.Preload("Tracks").Preload("User").First(playlist, id).Error; err != nil {
		return nil, err
	}
	return playlist, nil
}

func (p *playlistPersistence) ListByUserID(ctx context.Context, userID int) ([]*models.Playlist, error) {
	playlists := make([]*models.Playlist, 0)
	if err := p.db.Preload("Tracks").Preload("User").Where("user_id = ?", userID).Find(&playlists).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, err
		}
		return nil, err
	}
	return playlists, nil
}

func (p *playlistPersistence) Create(ctx context.Context, playlist *models.Playlist) (*models.Playlist, error) {
	if err := p.db.Create(playlist).Error; err != nil {
		return nil, err
	}
	return playlist, nil
}

func (p *playlistPersistence) Delete(ctx context.Context, playlist *models.Playlist) (*models.Playlist, error) {
	if err := p.db.Where("id = ? AND user_id = ?", playlist.ID, playlist.UserID).Take(playlist).Delete(playlist).Error; err != nil {
		return nil, err
	}
	return playlist, nil
}
