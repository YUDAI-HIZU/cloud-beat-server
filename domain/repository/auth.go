package repository

import "context"

type AuthRepository interface {
	SetIDToClaims(ctx context.Context, uid string, id int) error
}
