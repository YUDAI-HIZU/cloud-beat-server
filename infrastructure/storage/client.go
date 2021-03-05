package storage

import (
	"app/config"
	"context"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

type Client struct {
	client     *storage.Client
	rootPrefix string
}

func InitClient() *storage.Client {
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsJSON([]byte(config.GCSAccount)))

	if err != nil {
		panic(err)
	}
	return client
}
