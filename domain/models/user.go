package models

import "time"

type User struct {
	ID           string `json:"id"`
	UID          string `gorm:"primaryKey" json:"uid"`
	DisplayName  string `json:"displayName"`
	WebURL       string `json:"webUrl"`
	Introduction string `json:"introduction"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
