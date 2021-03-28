package models

import (
	"app/config"
	"fmt"
	"time"
)

type Track struct {
	ID            int
	Title         string
	ThumbnailPath string
	SoundPath     string
	Description   string
	YoutubeLink   string
	UserID        int
	GenreID       int
	User          *User
	Genre         *Genre
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (t *Track) ThumbnailUrl() string {
	if t.ThumbnailPath == "" {
		return ""
	}
	return fmt.Sprintf("%s/%s", config.AssetEndpoint, t.ThumbnailPath)
}

func (t *Track) SoundUrl() string {
	return fmt.Sprintf("%s/%s", config.AssetEndpoint, t.SoundPath)
}
