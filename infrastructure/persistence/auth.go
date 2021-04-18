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

func (p *authPersistence) SetIDToClaims(ctx context.Context, uid string, id int) error {
	claims := map[string]interface{}{"id": id}
	return p.auth.SetCustomUserClaims(ctx, uid, claims)
}
