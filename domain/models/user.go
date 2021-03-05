package models

import (
	"time"
)

type User struct {
	ID           int
	UID          string
	DisplayName  string
	WebURL       string
	Introduction string
	Icon         *Image `gorm:"foreignkey:OwnerID"`
	Cover        *Image `gorm:"foreignkey:OwnerID"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
