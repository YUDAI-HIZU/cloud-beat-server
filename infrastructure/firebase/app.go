package firebase

import (
	"app/config"
	"context"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

type firebaseApp struct {
	App *firebase.App
}

func NewFirebaseApp(ctx context.Context) firebaseApp {
	app, err := firebase.NewApp(ctx, nil, option.WithCredentialsJSON([]byte(config.FirebaseAccount)))
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	return firebaseApp{
		app,
	}
}
