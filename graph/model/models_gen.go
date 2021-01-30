// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"app/domain/models"
)

type SignInInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInPayload struct {
	Token string       `json:"token"`
	User  *models.User `json:"user"`
}

type SignUpInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpPayload struct {
	Token string       `json:"token"`
	User  *models.User `json:"user"`
}
