package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"app/domain/models"
	"app/graph/generated"
	"app/graph/model"
	"context"
	"fmt"
)

func (r *mutationResolver) CreateTrack(ctx context.Context, input model.CreateTrackInput) (*models.Track, error) {
	userID := ctx.Value("ID").(int)
	return r.track.Create(userID, input)
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*models.User, error) {
	return r.user.Create(input)
}

func (r *mutationResolver) UpdateTrack(ctx context.Context, input model.UpdateTrackInput) (*models.Track, error) {
	userID := ctx.Value("ID").(int)
	return r.track.Update(userID, input)
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUserInput) (*models.User, error) {
	id := ctx.Value("ID").(int)
	return r.user.Update(id, input)
}

func (r *queryResolver) Track(ctx context.Context, id int) (*models.Track, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) User(ctx context.Context, id int) (*models.User, error) {
	return r.user.GetByID(id)
}

func (r *queryResolver) CurrentUser(ctx context.Context) (*models.User, error) {
	id := ctx.Value("ID").(int)
	return r.user.GetByID(id)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
