package graph

import (
	"cloud.google.com/go/storage"

	"github.com/jinzhu/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB      *gorm.DB
	Storage *storage.Client
}
