package models

import "time"

type DeviceToken struct {
	ID        string
	UserID    string
	Token     string
	User      *User
	CreatedAt time.Time
	UpdatedAt time.Time
}
