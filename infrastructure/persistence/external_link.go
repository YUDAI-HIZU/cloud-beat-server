package persistence

import (
	"app/domain/models"
	"app/domain/repository"
	"context"

	"github.com/jinzhu/gorm"
)

type externalLinkPersistence struct {
	db *gorm.DB
}

func NewExternalLinkPersistence(db *gorm.DB) repository.ExternalLinkRepository {
	return &externalLinkPersistence{
		db,
	}
}

func (p *externalLinkPersistence) Get(ctx context.Context, userID int) (*models.ExternalLink, error) {
	externalLink := &models.ExternalLink{}
	if err := p.db.Where("user_id = ?", userID).Take(externalLink).Error; err != nil {
		return nil, err
	}
	return externalLink, nil
}

func (p *externalLinkPersistence) Create(ctx context.Context, externalLink *models.ExternalLink) (*models.ExternalLink, error) {
	if err := p.db.Create(externalLink).Error; err != nil {
		return nil, err
	}
	return externalLink, nil
}

func (p *externalLinkPersistence) Update(ctx context.Context, externalLink *models.ExternalLink) (*models.ExternalLink, error) {
	if err := p.db.Update(externalLink).Error; err != nil {
		return nil, err
	}
	return externalLink, nil
}
