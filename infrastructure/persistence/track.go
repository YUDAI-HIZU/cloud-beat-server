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
		db,
	}
}

func (p *trackPersistence) List() ([]*models.Track, error) {
	tracks := make([]*models.Track, 0)
	if err := p.db.Preload("User").Order("created_at desc").Find(&tracks).Error; err != nil {
		return nil, err
	}
	return tracks, nil
}

func (p *trackPersistence) ListByUserID(userID int) ([]*models.Track, error) {
	tracks := make([]*models.Track, 0)
	if err := p.db.Preload("User").Where("user_id = ?", userID).Find(&tracks).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, err
		}
		return nil, err
	}
	return tracks, nil
}

func (p *trackPersistence) Get(id int) (*models.Track, error) {
	var track *models.Track
	if err := p.db.Preload("User").Take(track).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, err
		}
		return nil, err
	}
	return track, nil
}

func (p *trackPersistence) Create(track *models.Track) (*models.Track, error) {
	if err := p.db.Create(track).Error; err != nil {
		return nil, err
	}
	return track, nil
}

func (p *trackPersistence) Update(track *models.Track) (*models.Track, error) {
	if err := p.db.Update(track).Error; err != nil {
		return nil, err
	}
	return track, nil
}
