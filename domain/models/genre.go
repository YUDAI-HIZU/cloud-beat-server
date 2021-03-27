package models

import "time"

type Genre struct {
	ID        int
	Name      string
	Tracks    []*Track
	CreatedAt time.Time
	UpdatedAt time.Time
}
