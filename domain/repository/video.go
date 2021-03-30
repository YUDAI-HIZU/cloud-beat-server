package repository

import "github.com/99designs/gqlgen/graphql"

type VideoRepository interface {
	Upload(prefix string, video *graphql.Upload) (string, error)
}
