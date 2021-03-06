package models

import (
	"app/config"
	"bytes"
	"fmt"
	"time"
)

type Image struct {
	ID        int
	Name      string
	Buf       *bytes.Buffer
	OwnerID   int
	OwnerType ImageOwnerType
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ImageOwnerType string

const (
	ImageOwnerTypeUserIcon  ImageOwnerType = "USER_ICON"
	ImageOwnerTypeUserCover ImageOwnerType = "USER_COVER"
)

func (i *Image) Url() string {
	return fmt.Sprintf("%s/%d/%s", config.AssetEndpoint, i.ID, i.Name)
}
