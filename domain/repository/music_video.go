package repository

import "app/domain/models"

type MusicVideoRepository interface {
	Create(musicVideo *models.MusicVideo) (*models.MusicVideo, error)
	Delete(id int, userID int) (*models.MusicVideo, error)
}
