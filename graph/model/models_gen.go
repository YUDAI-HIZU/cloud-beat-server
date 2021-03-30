// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"github.com/99designs/gqlgen/graphql"
)

type CreateExternalLinkInput struct {
	Twitter    *string `json:"twitter"`
	SoundCloud *string `json:"soundCloud"`
	Facebook   *string `json:"facebook"`
	Youtube    *string `json:"youtube"`
	Instagram  *string `json:"instagram"`
}

type CreateMusicVideoInput struct {
	Text  string         `json:"text"`
	Video graphql.Upload `json:"video"`
}

type CreatePlaylistInput struct {
	Title    string `json:"title"`
	TrackIDs []int  `json:"trackIDs"`
}

type CreatePlaylistSourceInput struct {
	PlaylistID int `json:"playlistID"`
	TrackID    int `json:"trackID"`
}

type CreateTrackInput struct {
	Title       string          `json:"title"`
	Sound       graphql.Upload  `json:"sound"`
	Thumbnail   *graphql.Upload `json:"thumbnail"`
	Description string          `json:"description"`
	YoutubeLink *string         `json:"youtube_link"`
	GenreID     int             `json:"genre_id"`
}

type CreateUserInput struct {
	UID         string `json:"uid"`
	DisplayName string `json:"displayName"`
}

type DeleteMusicVideoInput struct {
	ID int `json:"id"`
}

type DeletePlaylistInput struct {
	ID int `json:"id"`
}

type DeletePlaylistSourceInput struct {
	PlaylistID int   `json:"playlistID"`
	TrackIDs   []int `json:"trackIDs"`
}

type DeleteTrackInput struct {
	ID int `json:"id"`
}

type UpdateExternalLinkInput struct {
	Twitter    *string `json:"twitter"`
	SoundCloud *string `json:"soundCloud"`
	Facebook   *string `json:"facebook"`
	Youtube    *string `json:"youtube"`
	Instagram  *string `json:"instagram"`
}

type UpdateTrackInput struct {
	Title       string          `json:"title"`
	Sound       *graphql.Upload `json:"sound"`
	Thumbnail   *graphql.Upload `json:"thumbnail"`
	Description string          `json:"description"`
	YoutubeLink *string         `json:"youtube_link"`
	GenreID     int             `json:"genre_id"`
}

type UpdateUserInput struct {
	Icon         *graphql.Upload `json:"icon"`
	DisplayName  *string         `json:"displayName"`
	WebURL       *string         `json:"webUrl"`
	Introduction *string         `json:"introduction"`
}
