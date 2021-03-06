package persistence

import (
	"app/config"
	"app/domain/models"
	"app/domain/repository"
	"app/graph/model"
	"context"
	"fmt"
	"io"

	"cloud.google.com/go/storage"
	"github.com/jinzhu/gorm"
)

type imagePersistence struct {
	db      *gorm.DB
	storage *storage.Client
}

func NewImagePersistence(db *gorm.DB, storage *storage.Client) repository.ImageRepository {
	return &imagePersistence{
		db:      db,
		storage: storage,
	}
}

func (i *imagePersistence) Create(img *models.Image) (*models.Image, error) {
	if err := i.db.Model(img).Create(img).
		Error; err != nil {
		return nil, err
	}
	return img, nil
}

func (i *imagePersistence) Upload(id int, name string, input model.CreateImageInput) error {
	path := fmt.Sprintf("%d/%s", id, name)
	sw := i.storage.Bucket(config.BucketName).Object(path).NewWriter(context.Background())
	sw.ContentType = input.Image.ContentType
	if _, err := io.Copy(sw, input.Image.File); err != nil {
		return err
	}
	if err := sw.Close(); err != nil {
		return err
	}
	return nil
}
