package models

import (
	"fmt"
	"time"
)

type User struct {
	ID           string `json:"id"`
	UID          string `json:"uid"`
	DisplayName  string `json:"displayName"`
	WebURL       string `json:"webUrl"`
	Introduction string `json:"introduction"`
	IconName     string `json:"iconName"`
	CoverName    string `json:"coverName"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (u *User) IconURL() string {
	if u.CoverName == "" {
		return ""
	}
	return fmt.Sprintf("%s/%s", "dev", "uesr")
}

func (u *User) CoverURL() string {
	if u.CoverName == "" {
		return ""
	}
	return fmt.Sprintf("%s/%s", "dev", "uesr")
}
