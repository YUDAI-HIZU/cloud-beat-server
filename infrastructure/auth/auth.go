package auth

import (
	"app/infrastructure/firebase"
	"context"
	"log"

	"firebase.google.com/go/auth"
)

func NewAuthClient(ctx context.Context) *auth.Client {
	firebase := firebase.NewFirebaseApp(ctx)
	client, err := firebase.App.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}
	return client
}
