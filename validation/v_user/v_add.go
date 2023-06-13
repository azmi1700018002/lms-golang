package vuser

import (
	"fmt"
	"lms/model"
	"regexp"
)

func ValidateUser(user *model.User) error {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(user.Email) {
		return fmt.Errorf("invalid email format")
	}

	if len(user.Password) < 6 {
		return fmt.Errorf("password length must be at least 6 characters")
	}

	// add more validation rules if necessary

	return nil
}
