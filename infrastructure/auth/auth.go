package auth

import (
	"app/infrastructure/firebase"
	"context"
	"log"

	"firebase.google.com/go/auth"
)

type authClient struct {
	Client *auth.Client
}

func NewAuthClient(ctx context.Context) authClient {
	firebase := firebase.NewFirebaseApp(ctx)
	client, err := firebase.App.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}
	return authClient{
		client,
	}
}

func (a *authClient) VerifyIDToken(ctx context.Context, idToken string) (*auth.Token, error) {
	return a.Client.VerifyIDToken(ctx, idToken)
}

func (a *authClient) SetIDToClaims(uid string, id int) error {
	ctx := context.Background()
	claims := map[string]interface{}{"ID": id}

	return a.Client.SetCustomUserClaims(ctx, uid, claims)
}
