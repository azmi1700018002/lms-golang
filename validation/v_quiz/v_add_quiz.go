package v_quiz

import (
	"errors"
	"lms/model"
)

func ValidateQuiz(quiz model.Quiz) error {
	if quiz.QuizName == "" {
		return errors.New("quiz name is required")
	}
	if quiz.QuizDesc == "" {
		return errors.New("quiz desct is required")
	}
	if quiz.QuizType == "" {
		return errors.New("quiz type is required")
	}
	return nil
}
