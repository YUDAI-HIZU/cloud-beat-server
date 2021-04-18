package repository

import (
	"app/domain/models"
	"context"
)

type UserRepository interface {
	Get(ctx context.Context, id int) (*models.User, error)
	Create(ctx context.Context, user *models.User) (*models.User, error)
	Update(ctx context.Context, user *models.User) (*models.User, error)
}
