package v_quiz

import (
	"errors"
	"lms/model"
)

func ValidateUpdateQuiz(quiz model.Quiz) error {
	if quiz.QuizName == "" {
		return errors.New("quiz name is required")
	}
	if quiz.QuizDesc == "" {
		return errors.New("quiz desct is required")
	}
	if quiz.QuizType == "" {
		return errors.New("quizType is required")
	}
	return nil
}
