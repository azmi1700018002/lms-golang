package rquiz

import (
	"lms/config/db"
	"lms/model"
)

type QuizDeleteRepository interface {
	DeleteQuizByID(IDQuiz int) error
}

type quizDeleteRepository struct{}

func DeleteQuizRepository() QuizDeleteRepository {
	return &quizDeleteRepository{}
}

func (*quizDeleteRepository) DeleteQuizByID(IDQuiz int) error {
	if err := db.Server().Unscoped().Where("id_quiz = ?", IDQuiz).Delete(&model.Quiz{}).Error; err != nil {
		return err
	}
	return nil
}
