package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"app/domain/models"
	"app/graph/generated"
	"app/graph/model"
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (r *mutationResolver) CreateExternalLink(ctx context.Context, input model.CreateExternalLinkInput) (*models.ExternalLink, error) {
	userID := ctx.Value("id").(string)

	e := &models.ExternalLink{
		Twitter:    *input.Twitter,
		SoundCloud: *input.SoundCloud,
		Facebook:   *input.Facebook,
		Youtube:    *input.Youtube,
		Instagram:  *input.Instagram,
		UserID:     userID,
	}

	return r.externalLink.Create(ctx, e)
}

func (r *mutationResolver) CreatePlaylist(ctx context.Context, input model.CreatePlaylistInput) (*models.Playlist, error) {
	userID := ctx.Value("id").(string)

	p := &models.Playlist{
		UserID: userID,
		Title:  input.Title,
	}

	for _, trackID := range input.TrackIDs {
		p.PlaylistSources = append(
			p.PlaylistSources,
			&models.PlaylistSource{
				TrackID: trackID,
			},
		)
	}

	return r.playlist.Create(ctx, p)
}

func (r *mutationResolver) CreatePlaylistSource(ctx context.Context, input model.CreatePlaylistSourceInput) (*models.PlaylistSource, error) {
	userID := ctx.Value("id").(string)

	p := &models.PlaylistSource{
		PlaylistID: input.PlaylistID,
		TrackID:    input.TrackID,
		Playlist: &models.Playlist{
			UserID: userID,
		},
	}

	return r.playlistSource.Create(ctx, p)
}

func (r *mutationResolver) CreateTrack(ctx context.Context, input model.CreateTrackInput) (*models.Track, error) {
	userID := ctx.Value("id").(string)

	t := &models.Track{
		Title:       input.Title,
		Description: input.Description,
		YoutubeLink: *input.YoutubeLink,
		ThumbName:   uuid.New().String(),
		AudioName:   uuid.New().String(),
		GenreID:     input.GenreID,
		UserID:      userID,
	}

	i := &models.Image{
		Prefix: "thumbs",
		Name:   t.ThumbName,
	}

	a := &models.Audio{
		Prefix: "audios",
		Name:   t.AudioName,
	}

	i.Buf.ReadFrom(input.Thumbnail.File)

	a.Buf.ReadFrom(input.Audio.File)

	return r.track.Create(ctx, t, i, a)
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*models.User, error) {
	u := &models.User{
		ID:          input.ID,
		DisplayName: input.DisplayName,
	}

	return r.user.Create(ctx, u)
}

func (r *mutationResolver) UpdateExternalLink(ctx context.Context, input model.UpdateExternalLinkInput) (*models.ExternalLink, error) {
	userID := ctx.Value("id").(string)

	e := &models.ExternalLink{
		Twitter:    *input.Twitter,
		SoundCloud: *input.SoundCloud,
		Facebook:   *input.Facebook,
		Youtube:    *input.Youtube,
		Instagram:  *input.Instagram,
		UserID:     userID,
	}

	return r.externalLink.Update(ctx, e)
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUserInput) (*models.User, error) {
	id := ctx.Value("id").(string)

	u := &models.User{
		ID:           id,
		DisplayName:  *input.DisplayName,
		WebURL:       *input.WebURL,
		IconName:     uuid.New().String(),
		Introduction: *input.Introduction,
	}

	i := &models.Image{
		Prefix: "icons",
		Name:   u.IconName,
	}

	i.Buf.ReadFrom(input.Icon.File)

	return r.user.Update(ctx, u, i)
}

func (r *mutationResolver) DeletePlaylist(ctx context.Context, input model.DeletePlaylistInput) (*models.Playlist, error) {
	userID := ctx.Value("id").(string)

	p := &models.Playlist{
		ID:     input.ID,
		UserID: userID,
	}

	return r.playlist.Delete(ctx, p)
}

func (r *mutationResolver) DeletePlaylistSource(ctx context.Context, input model.DeletePlaylistSourceInput) (*models.PlaylistSource, error) {
	userID := ctx.Value("id").(string)

	p := &models.PlaylistSource{
		PlaylistID: input.PlaylistID,
		TrackID:    input.TrackID,
		Playlist: &models.Playlist{
			UserID: userID,
		},
	}

	return r.playlistSource.Delete(ctx, p)
}

func (r *mutationResolver) DeleteTrack(ctx context.Context, input model.DeleteTrackInput) (*models.Track, error) {
	userID := ctx.Value("id").(string)

	t := &models.Track{
		ID:     input.ID,
		UserID: userID,
	}

	return r.track.Delete(ctx, t)
}

func (r *mutationResolver) DeleteUser(ctx context.Context) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ExternalLinkByUserID(ctx context.Context, userID string) (*models.ExternalLink, error) {
	return r.externalLink.Get(ctx, userID)
}

func (r *queryResolver) Genres(ctx context.Context) ([]*models.Genre, error) {
	return r.genre.List(ctx)
}

func (r *queryResolver) Playlist(ctx context.Context, id int) (*models.Playlist, error) {
	return r.playlist.Get(ctx, id)
}

func (r *queryResolver) PlaylistsByUserID(ctx context.Context, userID string) ([]*models.Playlist, error) {
	return r.playlist.ListByUserID(ctx, userID)
}

func (r *queryResolver) Track(ctx context.Context, id int) (*models.Track, error) {
	return r.track.Get(ctx, id)
}

func (r *queryResolver) TracksByUserID(ctx context.Context, userID string) ([]*models.Track, error) {
	return r.track.ListByUserID(ctx, userID)
}

func (r *queryResolver) NewReleaseTracks(ctx context.Context) ([]*models.Track, error) {
	return r.track.List(ctx)
}

func (r *queryResolver) PickUpTracks(ctx context.Context) ([]*models.Track, error) {
	return r.track.List(ctx)
}

func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
	return r.user.Get(ctx, id)
}

func (r *queryResolver) CurrentUser(ctx context.Context) (*models.User, error) {
	id := ctx.Value("id").(string)
	return r.user.Get(ctx, id)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
