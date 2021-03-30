package persistence

import (
	"app/domain/models"
	"app/domain/repository"

	"github.com/jinzhu/gorm"
)

type musicVideoPersistence struct {
	db *gorm.DB
}

func NewMusicVideoPersistence(db *gorm.DB) repository.MusicVideoRepository {
	return &musicVideoPersistence{
		db,
	}
}

func (p *musicVideoPersistence) Create(musicVideo *models.MusicVideo) (*models.MusicVideo, error) {
	if err := p.db.Create(musicVideo).Error; err != nil {
		return nil, err
	}
	return musicVideo, nil
}

func (p *musicVideoPersistence) Delete(id int, userID int) (*models.MusicVideo, error) {
	musicVideo := &models.MusicVideo{}
	if err := p.db.Where("id = ? AND user_id = ?", id, userID).Take(musicVideo).Delete(musicVideo).Error; err != nil {
		return nil, err
	}
	return musicVideo, nil
}
