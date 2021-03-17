package persistence

import (
	"app/domain/models"
	"app/domain/repository"

	"github.com/jinzhu/gorm"
)

type trackPersistence struct {
	db *gorm.DB
}

func NewTrackPersistence(db *gorm.DB) repository.TrackRepository {
	return &trackPersistence{
		db: db,
	}
}

func (p *trackPersistence) List() ([]*models.Track, error) {
	tracks := make([]*models.Track, 0)

	if err := p.db.Preload("User").Find(&tracks).Error; err != nil {

		return nil, err
	}
	return tracks, nil
}

func (p *trackPersistence) Create(track *models.Track) (*models.Track, error) {
	if err := p.db.Model(track).Create(track).Error; err != nil {
		return nil, err
	}
	return track, nil
}

func (p *trackPersistence) Update(track *models.Track) (*models.Track, error) {
	if err := p.db.Model(track).Update(track).Error; err != nil {
		return nil, err
	}
	return track, nil
}
