package ranswer

import (
	"lms/config/db"
	"lms/model"
)

type AnswerDeleteRepository interface {
	DeleteAnswerByID(IDAnswer int) error
}

type answerDeleteRepository struct{}

func DeleteAnswerRepository() AnswerDeleteRepository {
	return &answerDeleteRepository{}
}

func (*answerDeleteRepository) DeleteAnswerByID(IDAnswer int) error {
	if err := db.Server().Unscoped().Where("id_answer = ?", IDAnswer).Delete(&model.Answer{}).Error; err != nil {
		return err
	}
	return nil
}
