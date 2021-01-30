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
	"fmt"
)

func (r *mutationResolver) SignUp(ctx context.Context, input model.SignUpInput) (*model.SignUpPayload, error) {
	user, err := usecase.NewUserUseCase(persistence.NewUserPersistence()).Create(r.DB, &models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	})
	fmt.Println(user)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *mutationResolver) SignIn(ctx context.Context, input model.SignInInput) (*model.SignInPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
	var user *models.User
	return user, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
