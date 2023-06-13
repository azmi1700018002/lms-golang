package v_category

import (
	"errors"
	"lms/model"
)

func ValidateCategory(category model.Category) error {
	if category.Categoryname == "" {
		return errors.New("category name is required")
	}
	return nil
}
