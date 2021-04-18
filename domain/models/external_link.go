package models

import "errors"

type ExternalLink struct {
	UserID     int
	Twitter    string
	SoundCloud string
	Facebook   string
	Youtube    string
	Instagram  string
	User       *User
}

func (e *ExternalLink) Validation() error {
	if len(e.Twitter) > 100 {
		return errors.New("Twitter is 100 characters or less")
	}
	if len(e.SoundCloud) > 100 {
		return errors.New("SoundCloud is 100 characters or less")
	}
	if len(e.Facebook) > 100 {
		return errors.New("Facebook is 100 characters or less")
	}
	if len(e.Youtube) > 100 {
		return errors.New("Youtube is 100 characters or less")
	}
	if len(e.Instagram) > 100 {
		return errors.New("Instagram is 100 characters or less")
	}
	return nil
}
