package persistence

import (
	"app/config"
	"app/domain/models"
	"app/domain/repository"
	"context"
	"io"

	"cloud.google.com/go/storage"
)

type audioPersistence struct {
	storage *storage.Client
}

func NewAudioPersistence(storage *storage.Client) repository.AudioRepository {
	return &audioPersistence{
		storage,
	}
}

func (p *audioPersistence) Upload(ctx context.Context, audio *models.Audio) error {
	wc := p.storage.Bucket(config.BucketName).Object(audio.ObjName()).NewWriter(ctx)
	if _, err := io.Copy(wc, audio.Buf); err != nil {
		return err
	}
	if err := wc.Close(); err != nil {
		return err
	}
	return nil
}

func (p *audioPersistence) Delete(ctx context.Context, objName string) error {
	return p.storage.Bucket(config.BucketName).Object(objName).Delete(ctx)
}
