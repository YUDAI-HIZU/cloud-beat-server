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
	Introduction string
	IconPath     string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (u *User) IconUrl() string {
	if u.IconPath == "" {
		return ""
	}
	return fmt.Sprintf("%s/%s", config.AssetEndpoint, u.IconPath)
}
