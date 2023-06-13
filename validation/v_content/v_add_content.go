package v_content

import (
	"errors"
	"lms/model"
)

func ValidateContent(content model.Content) error {
	if content.IDsection == 0 {
		return errors.New("idsection is required")
	}
	if content.ContentTitle == "" {
		return errors.New("content Title is required")
	}
	if content.ContentType == 0 {
		return errors.New("status is required")
	}
	if content.Image == "" {
		return errors.New("image is required")
	}
	if content.Link == "" {
		return errors.New("link is required")
	}
	return nil
}
