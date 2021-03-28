package repository

import "app/domain/models"

type ExternalLinkRepository interface {
	Get(userID int) (*models.ExternalLink, error)
	Create(externalLink *models.ExternalLink) (*models.ExternalLink, error)
	Update(externalLink *models.ExternalLink) (*models.ExternalLink, error)
}
