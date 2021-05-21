package repository

import (
	"app/domain/models"
	"context"
)

type ExternalLinkRepository interface {
	Get(ctx context.Context, userID string) (*models.ExternalLink, error)
	Create(ctx context.Context, externalLink *models.ExternalLink) (*models.ExternalLink, error)
	Update(ctx context.Context, externalLink *models.ExternalLink) (*models.ExternalLink, error)
}
