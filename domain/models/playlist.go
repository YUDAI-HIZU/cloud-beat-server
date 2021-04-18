package models

import (
	"errors"
	"time"
)

type Playlist struct {
	ID     int
	Title  string
	UserID int
	User   *User
	// TODO: infrastructureに依存させない
	Tracks          []*Track `gorm:"many2many:playlist_sources"`
	PlaylistSources []*PlaylistSource
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func (p *Playlist) Validation() error {
	if p.Title == "" {
		return errors.New("required title")
	}
	if len(p.Title) > 20 {
		return errors.New("title is 20 characters or less")
	}
	return nil
}
