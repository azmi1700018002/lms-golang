package ranswer

import (
	"lms/config/db"
	"lms/model"
)

type AnswerUpdateRepository interface {
	UpdateAnswerByID(id_answer int, input model.Answer) (answer model.Answer, err error)
}

type answerUpdateRepository struct{}

func UpdateAnswerRepository() AnswerUpdateRepository {
	return &answerUpdateRepository{}
}

func (r *answerUpdateRepository) UpdateAnswerByID(id_answer int, input model.Answer) (answer model.Answer, err error) {

	var existingAnswer model.Answer
	if err := db.Server().First(&existingAnswer, id_answer).Error; err != nil {
		return answer, err
	}

	answer = model.Answer{
		IDAnswer:   id_answer,
		IDQuestion: input.IDQuestion,
		AnswerText: input.AnswerText,
		IsCorrect:  input.IsCorrect,
	}

	if err = db.Server().Where("id_answer = ?", id_answer).Updates(&answer).Error; err != nil {
		return answer, err
	}

	return answer, nil
}
