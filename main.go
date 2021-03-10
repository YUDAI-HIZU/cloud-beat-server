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

	user := usecase.NewUserUseCase(persistence.NewUserPersistence(db), persistence.NewImagePersistence(storage))
	track := usecase.NewTrackUseCase(persistence.NewTrackPersistence(db), persistence.NewImagePersistence(storage))

	resolver := graph.NewResolver(user, track)

	h := handler.NewGraphQLHandler(resolver)
	p := handler.NewPlayGroundHandler()

	r := router.NewRouter(h, p)
	r.Run()
}
