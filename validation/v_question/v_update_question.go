package v_question

import (
	"errors"
	"lms/model"
)

func ValidateUpdateQuestion(question model.Question) error {
	if question.QuestionName == "" {
		return errors.New("question name is required")
	}
	return nil
}
