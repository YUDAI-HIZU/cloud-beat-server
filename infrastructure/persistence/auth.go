package persistence

import (
	"app/domain/repository"
	"context"

	"firebase.google.com/go/auth"
)

type authPersistence struct {
	auth *auth.Client
}

func NewAuthPersistence(auth *auth.Client) repository.AuthRepository {
	return &authPersistence{
		auth,
	}
}

func (p *authPersistence) SetIDToClaims(uid string, id int) error {
	ctx := context.Background()
	claims := map[string]interface{}{"ID": id}

	return p.auth.SetCustomUserClaims(ctx, uid, claims)
}
