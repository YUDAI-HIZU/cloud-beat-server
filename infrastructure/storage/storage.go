package storage

import (
	"app/config"
	"context"
	"log"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

func NewStorage(ctx context.Context) *storage.Client {
	client, err := storage.NewClient(ctx, option.WithCredentialsJSON([]byte(config.CloudStorageAccount)))
	if err != nil {
		log.Fatalf("error getting storage client: %v\n", err)
	}
	return client
}
