package models

import "time"

type UserDeviceToken struct {
	UserID      string `json:"id"`
	DeviceToken string `json:"deviceToken"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
