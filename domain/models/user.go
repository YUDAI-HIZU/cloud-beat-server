package models

import "time"

type User struct {
	ID          string `json:"id"`
	UID         string `json:"uid"`
	DisplayName string `json:"displayName"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
