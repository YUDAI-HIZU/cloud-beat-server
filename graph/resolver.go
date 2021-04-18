package graph

import "app/usecase"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	externalLink   usecase.ExternalLinkUsecase
	genre          usecase.GenreUsecase
	playlist       usecase.PlaylistUsecase
	playlistSource usecase.PlaylistSourceUsecase
	user           usecase.UserUsecase
	track          usecase.TrackUsecase
}

func NewResolver(
	externalLink usecase.ExternalLinkUsecase,
	genre usecase.GenreUsecase,
	playlist usecase.PlaylistUsecase,
	playlistSource usecase.PlaylistSourceUsecase,
	user usecase.UserUsecase,
	track usecase.TrackUsecase,
) *Resolver {
	return &Resolver{
		externalLink,
		genre,
		playlist,
		playlistSource,
		user,
		track,
	}
}
