package persistence

import (
	"app/config"
	"app/domain/repository"
	"context"
	"fmt"
	"io"

	"cloud.google.com/go/storage"
	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
)

type imagePersistence struct {
	storage *storage.Client
}

func NewImagePersistence(storage *storage.Client) repository.ImageRepository {
	return &imagePersistence{
		storage: storage,
	}
}

func (i *imagePersistence) Upload(prefix string, img *graphql.Upload) (string, error) {
	path := fmt.Sprintf("%s/%s", prefix, uuid.New().String())
	sw := i.storage.Bucket(config.BucketName).Object(path).NewWriter(context.Background())
	sw.ContentType = img.ContentType
	if _, err := io.Copy(sw, img.File); err != nil {
		return "", err
	}
	if err := sw.Close(); err != nil {
		return "", err
	}
	return path, nil
}
