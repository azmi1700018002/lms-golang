package v_answer

import (
	"errors"
	"lms/model"
)

func ValidateAnswer(answer model.Answer) error {
	if answer.AnswerText == "" {
		return errors.New("answer name is required")
	}
	return nil
}
