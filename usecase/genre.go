package usecase

import (
	"app/domain/models"
	"app/domain/repository"
)

type genreUsecase struct {
	genreRepository repository.GenreRepository
}

type GenreUsecase interface {
	List() ([]*models.Genre, error)
}

func NewGenreUsecase(g repository.GenreRepository) GenreUsecase {
	return &genreUsecase{
		g,
	}
}

func (u *genreUsecase) List() ([]*models.Genre, error) {
	return u.genreRepository.List()
}
