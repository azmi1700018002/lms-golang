package rquiz

import (
	"lms/config/db"
	"lms/model"
)

type QuizUpdateRepository interface {
	UpdateQuizByID(id_quiz int, input model.Quiz) (quiz model.Quiz, err error)
}

type quizUpdateRepository struct{}

func UpdateQuizRepository() QuizUpdateRepository {
	return &quizUpdateRepository{}
}

func (r *quizUpdateRepository) UpdateQuizByID(id_quiz int, input model.Quiz) (quiz model.Quiz, err error) {

	var existingQuiz model.Quiz
	if err := db.Server().First(&existingQuiz, id_quiz).Error; err != nil {
		return quiz, err
	}

	quiz = model.Quiz{
		IDQuiz:   id_quiz,
		IDUser:   input.IDUser,
		QuizName: input.QuizName,
		QuizDesc: input.QuizDesc,
		QuizType: input.QuizType,
		Score:    input.Score,
	}

	if err = db.Server().Where("id_quiz = ?", id_quiz).Updates(&quiz).Error; err != nil {
		return quiz, err
	}

	return quiz, nil
}
