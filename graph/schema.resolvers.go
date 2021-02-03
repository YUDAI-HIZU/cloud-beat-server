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

func (r *mutationResolver) SignUp(ctx context.Context, input model.SignUpInput) (*models.User, error) {
	user, err := usecase.NewUserUseCase(persistence.NewUserPersistence()).Create(r.DB, &models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	})
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *mutationResolver) SignIn(ctx context.Context, input model.SignInInput) (*model.SignInPayload, error) {
	user, token, err := usecase.NewUserUseCase(persistence.NewUserPersistence()).SignIn(r.DB, input.Email, input.Password)
	if err != nil {
		return nil, err
	}
	payload := &model.SignInPayload{User: user, Token: token}
	return payload, nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
	intID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	user, err := usecase.NewUserUseCase(persistence.NewUserPersistence()).GetByID(r.DB, intID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *queryResolver) CurrentUser(ctx context.Context) (*models.User, error) {
	id := int(ctx.Value("userID").(float64))
	user, err := usecase.NewUserUseCase(persistence.NewUserPersistence()).GetByID(r.DB, id)
	return user, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
