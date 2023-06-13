package validator

import (
	"fmt"
	"lms/model"

	"github.com/go-playground/validator/v10"
)

var (
	validate *validator.Validate
)

func init() {
	validate = validator.New()
}

func ValidateUser(user *model.User) error {
	if err := validate.Struct(user); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	return nil
}
