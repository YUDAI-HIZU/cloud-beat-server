package usecase

import (
	"app/domain/models"
	"app/domain/repository"
	"context"
)

type genreUsecase struct {
	genreRepository repository.GenreRepository
}

type GenreUsecase interface {
	List(ctx context.Context) ([]*models.Genre, error)
}

func NewGenreUsecase(g repository.GenreRepository) GenreUsecase {
	return &genreUsecase{
		g,
	}
}

func (u *genreUsecase) List(ctx context.Context) ([]*models.Genre, error) {
	return u.genreRepository.List(ctx)
}
