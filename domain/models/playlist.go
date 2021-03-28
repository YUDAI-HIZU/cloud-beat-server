package models

import "time"

type Playlist struct {
	ID              int
	Title           string
	UserID          int
	User            *User
	Tracks          []*Track `gorm:"many2many:playlist_sources"`
	PlaylistSources []*PlaylistSource
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
