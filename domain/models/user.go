package models

import "time"

type User struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	Email             string `json:"email"`
	Password          string `gorm:"-"`
	EncryptedPassword string `json:"encryptedPassword"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
