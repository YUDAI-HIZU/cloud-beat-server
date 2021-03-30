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

type videoPersistence struct {
	storage *storage.Client
}

func NewVideoPersistence(storage *storage.Client) repository.VideoRepository {
	return &videoPersistence{
		storage,
	}
}

func (p *videoPersistence) Upload(prefix string, Video *graphql.Upload) (string, error) {
	path := fmt.Sprintf("%s/%s/%s", config.ENV, prefix, uuid.New().String())
	sw := p.storage.Bucket(config.BucketName).Object(path).NewWriter(context.Background())
	sw.ContentType = Video.ContentType
	if _, err := io.Copy(sw, Video.File); err != nil {
		return "", err
	}
	if err := sw.Close(); err != nil {
		return "", err
	}
	return path, nil
}
