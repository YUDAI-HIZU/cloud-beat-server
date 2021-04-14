package repository

import "github.com/99designs/gqlgen/graphql"

type ImageRepository interface {
	Upload(prefix string, img *graphql.Upload) (string, error)
	Delete(path string) error
}
