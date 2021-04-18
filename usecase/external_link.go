package usecase

import (
	"app/domain/models"
	"app/domain/repository"
	"context"
)

type externalLinkUsecase struct {
	externalLinkRepository repository.ExternalLinkRepository
}

type ExternalLinkUsecase interface {
	Get(ctx context.Context, userID int) (*models.ExternalLink, error)
	Create(ctx context.Context, externalLink *models.ExternalLink) (*models.ExternalLink, error)
	Update(ctx context.Context, externalLink *models.ExternalLink) (*models.ExternalLink, error)
}

func NewExternalLinkUsecase(e repository.ExternalLinkRepository) ExternalLinkUsecase {
	return &externalLinkUsecase{
		e,
	}
}

func (u *externalLinkUsecase) Get(ctx context.Context, userID int) (*models.ExternalLink, error) {
	return u.externalLinkRepository.Get(ctx, userID)
}

func (u *externalLinkUsecase) Create(ctx context.Context, externalLink *models.ExternalLink) (*models.ExternalLink, error) {
	return u.externalLinkRepository.Create(ctx, externalLink)
}

func (u *externalLinkUsecase) Update(ctx context.Context, externalLink *models.ExternalLink) (*models.ExternalLink, error) {
	return u.externalLinkRepository.Update(ctx, externalLink)
}
