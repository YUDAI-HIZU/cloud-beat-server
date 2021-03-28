package models

type ExternalLink struct {
	UserID     int
	Twitter    string
	SoundCloud string
	Facebook   string
	Youtube    string
	Instagram  string
	User       *User
}
