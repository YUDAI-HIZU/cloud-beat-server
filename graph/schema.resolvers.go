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
	"strconv"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*models.User, error) {
	u := usecase.NewUserUseCase(persistence.NewUserPersistence())
	return u.Create(r.DB, input)
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUserInput) (*models.User, error) {
	u := usecase.NewUserUseCase(persistence.NewUserPersistence())
	return u.Update(r.DB, ctx.Value("UID").(string), input)
}

func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
	intID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	return usecase.NewUserUseCase(persistence.NewUserPersistence()).GetByID(r.DB, intID)
}

func (r *queryResolver) CurrentUser(ctx context.Context) (*models.User, error) {
	uid := ctx.Value("UID").(string)
	return usecase.NewUserUseCase(persistence.NewUserPersistence()).GetByUID(r.DB, uid)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
