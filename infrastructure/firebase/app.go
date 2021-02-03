package firebase

import (
	"app/config"
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

var app *firebase.App

func init() {
	var err error
	ctx := context.Background()
	app, err = firebase.NewApp(ctx, nil, option.WithCredentialsJSON([]byte(config.FirebaseAccount)))
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
		panic(err)
	}
}

func InitAuth(ctx context.Context) *auth.Client {
	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error initializing auth: %v\n", err)
		panic(err)
	}
	return client
}
