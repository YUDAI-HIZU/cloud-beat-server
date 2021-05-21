package repository

import (
	"app/domain/models"
	"context"
)

type UserRepository interface {
	Get(ctx context.Context, id string) (*models.User, error)
	Create(ctx context.Context, user *models.User) (*models.User, error)
	Update(ctx context.Context, user *models.User) (*models.User, error)
}
