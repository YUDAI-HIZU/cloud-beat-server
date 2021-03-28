package usecase

import (
	"app/domain/models"
	"app/domain/repository"
	"app/graph/model"
)

type externalLinkUsecase struct {
	externalLinkRepository repository.ExternalLinkRepository
}

type ExternalLinkUsecase interface {
	Get(userID int) (*models.ExternalLink, error)
	Create(userID int, input model.CreateExternalLinkInput) (*models.ExternalLink, error)
	Update(userID int, input model.UpdateExternalLinkInput) (*models.ExternalLink, error)
}

func NewExternalLinkUsecase(e repository.ExternalLinkRepository) ExternalLinkUsecase {
	return &externalLinkUsecase{
		e,
	}
}

func (u *externalLinkUsecase) Get(userID int) (*models.ExternalLink, error) {
	return u.externalLinkRepository.Get(userID)
}

func (u *externalLinkUsecase) Create(userID int, input model.CreateExternalLinkInput) (*models.ExternalLink, error) {
	externalLink := &models.ExternalLink{
		Twitter:    *input.Twitter,
		SoundCloud: *input.SoundCloud,
		Facebook:   *input.Facebook,
		Youtube:    *input.Youtube,
		Instagram:  *input.Instagram,
		UserID:     userID,
	}
	return u.externalLinkRepository.Create(externalLink)
}

func (u *externalLinkUsecase) Update(userID int, input model.UpdateExternalLinkInput) (*models.ExternalLink, error) {
	externalLink := &models.ExternalLink{
		Twitter:    *input.Twitter,
		SoundCloud: *input.SoundCloud,
		Facebook:   *input.Facebook,
		Youtube:    *input.Youtube,
		Instagram:  *input.Instagram,
		UserID:     userID,
	}
	return u.externalLinkRepository.Update(externalLink)
}
