package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"app/domain/models"
	"app/graph/generated"
	"app/graph/model"
	"app/infrastructure/persistence"
	"app/usecase"
	"context"
)

func (r *mutationResolver) CreateImage(ctx context.Context, input model.CreateImageInput) (*models.Image, error) {
	id := ctx.Value("ID").(int)
	i := usecase.NewImageUseCase(persistence.NewImagePersistence(r.DB, r.Storage))
	return i.Create(id, input)
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*models.User, error) {
	u := usecase.NewUserUseCase(persistence.NewUserPersistence(r.DB))
	return u.Create(input)
}

func (r *mutationResolver) UpdateImage(ctx context.Context, input model.UpdateImageInput) (*models.Image, error) {
	panic("o")
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUserInput) (*models.User, error) {
	id := ctx.Value("ID").(int)
	u := usecase.NewUserUseCase(persistence.NewUserPersistence(r.DB))
	return u.Update(id, input)
}

func (r *queryResolver) User(ctx context.Context, id int) (*models.User, error) {
	u := usecase.NewUserUseCase(persistence.NewUserPersistence(r.DB))
	return u.GetByID(id)
}

func (r *queryResolver) CurrentUser(ctx context.Context) (*models.User, error) {
	id := ctx.Value("ID").(int)
	u := usecase.NewUserUseCase(persistence.NewUserPersistence(r.DB))
	return u.GetByID(id)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
