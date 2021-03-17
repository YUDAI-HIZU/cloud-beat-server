package models

import (
	"app/config"
	"fmt"
	"time"
)

type User struct {
	ID           int
	UID          string
	DisplayName  string
	WebURL       string
	Twitter      string
	SoundCloud   string
	Facebook     string
	Youtube      string
	Instagram    string
	Introduction string
	IconPath     string
	CoverPath    string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (u *User) IconUrl() string {
	if u.IconPath == "" {
		return ""
	}
	return fmt.Sprintf("%s/%s", config.AssetEndpoint, u.IconPath)
}

func (u *User) CoverUrl() string {
	if u.CoverPath == "" {
		return ""
	}
	return fmt.Sprintf("%s/%s", config.AssetEndpoint, u.CoverPath)
}
