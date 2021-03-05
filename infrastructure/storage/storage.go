package storage

import (
	"app/config"
	"context"
	"log"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

type storageClient struct {
	Client *storage.Client
}

func NewStorageClient() *storageClient {
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsJSON([]byte(config.GCSAccount)))
	if err != nil {
		log.Fatalf("error getting storage client: %v\n", err)
	}
	return &storageClient{
		Client: client,
	}
}
