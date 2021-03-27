package repository

import "app/domain/models"

type GenreRepository interface {
	List() ([]*models.Genre, error)
}
