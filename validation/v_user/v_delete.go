package vuser

import (
	"errors"

	"github.com/google/uuid"
)

func ValidateUserID(id string) error {
	// Validate UUID format
	_, err := uuid.Parse(id)
	if err != nil {
		return errors.New("invalid ID")
	}

	return nil
}
