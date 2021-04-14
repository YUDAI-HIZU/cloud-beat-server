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
	db := database.NewDatabase()
	storage := storage.NewStorage()
	auth := auth.NewAuthClient(context.Background())

	externalLink := usecase.NewExternalLinkUsecase(
		persistence.NewExternalLinkPersistence(db),
	)
	musicVideo := usecase.NewMusicVideoUsecase(
		persistence.NewMusicVideoPersistence(db),
		persistence.NewVideoPersistence(storage),
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
		musicVideo,
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
