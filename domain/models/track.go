package models

import (
	"app/config"
	"errors"
	"fmt"
	"time"
)

type Track struct {
	ID          int
	Title       string
	ThumbName   string
	AudioName   string
	Description string
	YoutubeLink string
	UserID      int
	GenreID     int
	User        *User
	Genre       *Genre
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (t *Track) ThumbnailUrl() string {
	if t.ThumbName == "" {
		return ""
	}
	return fmt.Sprintf("%s/%s/%s", config.AssetsEndpoint, "thumbs", t.ThumbName)
}

func (t *Track) AudioURL() string {
	return fmt.Sprintf("%s/%s/%s", config.AssetsEndpoint, "audios", t.AudioName)
}

func (t *Track) Validation() error {
	if t.Title == "" {
		return errors.New("required title")
	}
	if len(t.Title) > 20 {
		return errors.New("title is 20 characters or less")
	}
	if len(t.Description) > 255 {
		return errors.New("description is 255 characters or less")
	}
	if len(t.YoutubeLink) > 255 {
		return errors.New("youtube link is 255 characters or less")
	}
	return nil
}
