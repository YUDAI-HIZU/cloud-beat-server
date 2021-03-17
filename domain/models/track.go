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
	UserID        int
	User          *User
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
