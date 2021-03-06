package persistence

import (
	"app/domain/models"
	"app/domain/repository"
	"context"

	"github.com/jinzhu/gorm"
)

type genrePersistence struct {
	db *gorm.DB
}

func NewGenrePersistence(db *gorm.DB) repository.GenreRepository {
	return &genrePersistence{
		db,
	}
}

func (p *genrePersistence) List(ctx context.Context) ([]*models.Genre, error) {
	genres := make([]*models.Genre, 0)
	if err := p.db.Find(&genres).Error; err != nil {
		return nil, err
	}
	return genres, nil
}
