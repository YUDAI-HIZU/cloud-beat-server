package main

import (
	"app/graph"
	"app/infrastructure/database"
	"app/infrastructure/handler"
	"app/infrastructure/persistence"
	"app/infrastructure/router"
	"app/infrastructure/storage"
	"app/usecase"
)

func main() {
	db := database.NewDatabase()
	storage := storage.NewStorage()

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
	)
	track := usecase.NewTrackUsecase(
		persistence.NewTrackPersistence(db),
		persistence.NewImagePersistence(storage),
		persistence.NewAudioPersistence(storage),
	)

	resolver := graph.NewResolver(genre, playlist, playlistSource, user, track)

	h := handler.NewGraphQLHandler(resolver)
	p := handler.NewPlayGroundHandler()

	r := router.NewRouter(h, p)
	r.Run()
}
