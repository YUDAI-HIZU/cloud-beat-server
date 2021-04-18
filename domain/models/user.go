package models

import (
	"app/config"
	"errors"
	"fmt"
	"time"
)

type User struct {
	ID           int
	UID          string
	DisplayName  string
	WebURL       string
	Introduction string
	IconName     string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (u *User) IconUrl() string {
	if u.IconName == "" {
		return ""
	}
	return fmt.Sprintf("%s/%s/%s", config.AssetEndpoint, "icons", u.IconName)
}

func (u *User) Validation() error {
	if u.UID == "" {
		return errors.New("required uid")
	}
	if u.DisplayName == "" {
		return errors.New("required display name")
	}
	if len(u.DisplayName) > 255 {
		return errors.New("display name is 255 characters or less")
	}
	if len(u.WebURL) > 255 {
		return errors.New("web url is 255 characters or less")
	}
	if len(u.Introduction) > 255 {
		return errors.New("introduction is 255 characters or less")
	}
	return nil
}
