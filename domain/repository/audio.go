package repository

import "github.com/99designs/gqlgen/graphql"

type AudioRepository interface {
	Upload(prefix string, audio *graphql.Upload) (string, error)
}
