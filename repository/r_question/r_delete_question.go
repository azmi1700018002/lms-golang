package rquestion

import (
	"lms/config/db"
	"lms/model"
)

type QuestionDeleteRepository interface {
	DeleteQuestionByID(IDQuestion int) error
}

type questionDeleteRepository struct{}

func DeleteQuestionRepository() QuestionDeleteRepository {
	return &questionDeleteRepository{}
}

func (*questionDeleteRepository) DeleteQuestionByID(IDQuestion int) error {
	// Delete all answers associated with the question
	if err := db.Server().Unscoped().Where("id_question = ?", IDQuestion).Delete(&model.Answer{}).Error; err != nil {
		return err
	}

	// Delete the question
	if err := db.Server().Unscoped().Where("id_question = ?", IDQuestion).Delete(&model.Question{}).Error; err != nil {
		return err
	}
	return nil
}
