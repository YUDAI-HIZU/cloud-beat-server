package persistence

import (
	"app/config"
	"app/domain/models"
	"app/domain/repository"
	"context"
	"io"

	"cloud.google.com/go/storage"
)

type imagePersistence struct {
	storage *storage.Client
}

func NewImagePersistence(storage *storage.Client) repository.ImageRepository {
	return &imagePersistence{
		storage,
	}
}

func (p *imagePersistence) Upload(ctx context.Context, image *models.Image) error {
	wc := p.storage.Bucket(config.BucketName).Object(image.ObjName()).NewWriter(ctx)
	if _, err := io.Copy(wc, image.Buf); err != nil {
		return err
	}
	if err := wc.Close(); err != nil {
		return err
	}
	return nil
}

func (p *imagePersistence) Delete(ctx context.Context, objName string) error {
	return p.storage.Bucket(config.BucketName).Object(objName).Delete(ctx)
}
