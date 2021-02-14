// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"github.com/99designs/gqlgen/graphql"
)

type CreateUserInput struct {
	UID         string `json:"uid"`
	DisplayName string `json:"displayName"`
}

type UpdateUserInput struct {
	DisplayName  *string         `json:"displayName"`
	WebURL       *string         `json:"webUrl"`
	Introduction *string         `json:"introduction"`
	IconImage    *graphql.Upload `json:"iconImage"`
	CoverImage   *graphql.Upload `json:"coverImage"`
}
