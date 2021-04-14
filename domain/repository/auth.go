package repository

type AuthRepository interface {
	SetIDToClaims(uid string, id int) error
}
