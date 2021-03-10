package graph

import "app/usecase"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	user  usecase.UserUseCase
	track usecase.TrackUseCase
}

func NewResolver(
	user usecase.UserUseCase,
	track usecase.TrackUseCase,
) *Resolver {
	return &Resolver{
		user,
		track,
	}
}
