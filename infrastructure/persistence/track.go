package persistence

import (
	"app/domain/models"
	"app/domain/repository"
	"context"

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

func (p *trackPersistence) List(ctx context.Context) ([]*models.Track, error) {
	tracks := make([]*models.Track, 0)
	if err := p.db.Preload("User").Order("created_at desc").Find(&tracks).Error; err != nil {
		return nil, err
	}
	return tracks, nil
}

func (p *trackPersistence) ListByUserID(ctx context.Context, userID string) ([]*models.Track, error) {
	tracks := make([]*models.Track, 0)
	if err := p.db.Preload("User").Where("user_id = ?", userID).Find(&tracks).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, err
		}
		return nil, err
	}
	return tracks, nil
}

func (p *trackPersistence) Get(ctx context.Context, id int) (*models.Track, error) {
	track := &models.Track{}
	if err := p.db.Preload("User").First(track, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, err
		}
		return nil, err
	}
	return track, nil
}

func (p *trackPersistence) Create(ctx context.Context, track *models.Track) (*models.Track, error) {
	if err := p.db.Create(track).Error; err != nil {
		return nil, err
	}
	return track, nil
}

func (p *trackPersistence) Delete(ctx context.Context, track *models.Track) (*models.Track, error) {
	if err := p.db.Where("id = ? AND user_id = ?", track.ID, track.UserID).Take(track).Delete(track).Error; err != nil {
		return nil, err
	}
	return track, nil
}
