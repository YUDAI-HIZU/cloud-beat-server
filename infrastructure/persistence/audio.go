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

type audioPersistence struct {
	storage *storage.Client
}

func NewAudioPersistence(storage *storage.Client) repository.AudioRepository {
	return &audioPersistence{
		storage,
	}
}

func (p *audioPersistence) Upload(prefix string, audio *graphql.Upload) (string, error) {
	path := fmt.Sprintf("%s/%s/%s", config.ENV, prefix, uuid.New().String())
	sw := p.storage.Bucket(config.BucketName).Object(path).NewWriter(context.Background())
	sw.ContentType = audio.ContentType
	if _, err := io.Copy(sw, audio.File); err != nil {
		return "", err
	}
	if err := sw.Close(); err != nil {
		return "", err
	}
	return path, nil
}
