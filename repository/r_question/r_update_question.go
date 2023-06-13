package rquestion

import (
	"lms/config/db"
	"lms/model"
)

type QuestionUpdateRepository interface {
	UpdateQuestionByID(id_question int, input model.Question) (question model.Question, err error)
}

type questionUpdateRepository struct{}

func UpdateQuestionRepository() QuestionUpdateRepository {
	return &questionUpdateRepository{}
}

func (r *questionUpdateRepository) UpdateQuestionByID(id_question int, input model.Question) (question model.Question, err error) {

	var existingQuestion model.Question
	if err := db.Server().First(&existingQuestion, id_question).Error; err != nil {
		return question, err
	}

	question = model.Question{
		IDQuestion:   id_question,
		IDQuiz:       input.IDQuiz,
		QuestionName: input.QuestionName,
	}

	if err = db.Server().Where("id_question = ?", id_question).Updates(&question).Error; err != nil {
		return question, err
	}

	return question, nil
}
