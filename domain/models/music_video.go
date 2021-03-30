package models

import (
	"app/config"
	"fmt"
)

type MusicVideo struct {
	ID        int
	Text      string
	VideoPath string
	UserID    int
	User      *User
}

func (m *MusicVideo) VideoUrl() string {
	return fmt.Sprintf("%s/%s", config.AssetEndpoint, m.VideoPath)
}
