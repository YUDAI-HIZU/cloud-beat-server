package graph

import "app/usecase"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	genre          usecase.GenreUsecase
	playlist       usecase.PlaylistUsecase
	playlistSource usecase.PlaylistSourceUsecase
	user           usecase.UserUsecase
	track          usecase.TrackUsecase
}

func NewResolver(
	genre usecase.GenreUsecase,
	playlist usecase.PlaylistUsecase,
	playlistSource usecase.PlaylistSourceUsecase,
	user usecase.UserUsecase,
	track usecase.TrackUsecase,
) *Resolver {
	return &Resolver{
		genre,
		playlist,
		playlistSource,
		user,
		track,
	}
}
