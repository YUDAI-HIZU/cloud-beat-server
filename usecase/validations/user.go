package validations

import (
	"app/domain/models"
	"errors"
)

func UserCreateValidation(user *models.User) error {
	if user.UID == "" {
		return errors.New("required uid")
	}
	if user.DisplayName == "" {
		return errors.New("required display name")
	}
	if len(user.DisplayName) > 20 {
		return errors.New("display name is 20 characters or less")
	}
	return nil
}

func UserUpdateValidation(user *models.User) error {
	if len(user.WebURL) > 255 {
		return errors.New("web url is 255 characters or less")
	}
	if len(user.Introduction) > 255 {
		return errors.New("introduction is 255 characters or less")
	}
	return nil
}
