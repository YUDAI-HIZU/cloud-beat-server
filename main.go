package main

import (
	"app/graph"
	"app/infrastructure/auth"
	"app/infrastructure/database"
	"app/infrastructure/handler"
	"app/infrastructure/persistence"
	"app/infrastructure/router"
	"app/infrastructure/storage"
	"app/usecase"
	"context"
)

func main() {
	ctx := context.Background()

	db := database.NewDatabase()
	storage := storage.NewStorage(ctx)
	auth := auth.NewAuthClient(ctx)

	externalLink := usecase.NewExternalLinkUsecase(
		persistence.NewExternalLinkPersistence(db),
	)
	genre := usecase.NewGenreUsecase(
		persistence.NewGenrePersistence(db),
	)
	playlist := usecase.NewPlaylistUsecase(
		persistence.NewPlaylistPersistence(db),
	)
	playlistSource := usecase.NewPlaylistSourceUsecase(
		persistence.NewPlaylistSourcePersistence(db),
	)
	user := usecase.NewUserUsecase(
		persistence.NewUserPersistence(db),
		persistence.NewImagePersistence(storage),
		persistence.NewAuthPersistence(auth),
	)
	track := usecase.NewTrackUsecase(
		persistence.NewTrackPersistence(db),
		persistence.NewImagePersistence(storage),
		persistence.NewAudioPersistence(storage),
	)

	resolver := graph.NewResolver(
		externalLink,
		genre,
		playlist,
		playlistSource,
		user,
		track,
	)

	h := handler.NewGraphQLHandler(resolver)
	p := handler.NewPlayGroundHandler()

	r := router.NewRouter(h, p)
	r.Run()
}
