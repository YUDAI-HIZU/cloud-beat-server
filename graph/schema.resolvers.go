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

func (r *mutationResolver) CreatePlaylist(ctx context.Context, input model.CreatePlaylistInput) (*models.Playlist, error) {
	userID := 1
	return r.playlist.Create(userID, input)
}

func (r *mutationResolver) CreatePlaylistSource(ctx context.Context, input model.CreatePlaylistSourceInput) (*models.PlaylistSource, error) {
	userID := 1
	return r.playlistSource.Create(userID, input)
}

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

func (r *mutationResolver) DeletePlaylist(ctx context.Context, input model.DeletePlaylistInput) (*models.Playlist, error) {
	userID := 1
	return r.playlist.Delete(userID, input)
}

func (r *mutationResolver) DeletePlaylistSource(ctx context.Context, input model.DeletePlaylistSourceInput) ([]*models.PlaylistSource, error) {
	userID := 1
	return r.playlistSource.BatchDelete(userID, input)
}

func (r *queryResolver) Genres(ctx context.Context) ([]*models.Genre, error) {
	return r.genre.List()
}

func (r *queryResolver) PlaylistsByUser(ctx context.Context, userID int) ([]*models.Playlist, error) {
	return r.playlist.ListByUserID(userID)
}

func (r *queryResolver) PlaylistsByMe(ctx context.Context) ([]*models.Playlist, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Track(ctx context.Context, id int) (*models.Track, error) {
	return r.track.Get(id)
}

func (r *queryResolver) TracksByUser(ctx context.Context, userID int) ([]*models.Track, error) {
	return r.track.ListByUserID(userID)
}

func (r *queryResolver) NewReleaseTracks(ctx context.Context) ([]*models.Track, error) {
	return r.track.List()
}

func (r *queryResolver) PickUpTracks(ctx context.Context) ([]*models.Track, error) {
	return r.track.List()
}

func (r *queryResolver) User(ctx context.Context, id int) (*models.User, error) {
	return r.user.Get(id)
}

func (r *queryResolver) Me(ctx context.Context) (*models.User, error) {
	id := ctx.Value("ID").(int)
	return r.user.Get(id)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
