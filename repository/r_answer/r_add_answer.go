package ranswer

import (
	"context"
	"lms/model"

	"gorm.io/gorm"
)

type AnswerRepository interface {
	Create(ctx context.Context, answer *model.Answer) error
}

type answerRepository struct {
	db *gorm.DB
}

func NewAnswerRepository(db *gorm.DB) AnswerRepository {
	return &answerRepository{db: db}
}

func (r *answerRepository) Create(ctx context.Context, answer *model.Answer) error {

	// find quiz by id_question
	var question model.Question
	if err := r.db.Where("id_question = ?", answer.IDQuestion).First(&question).Error; err != nil {
		return err
	}

	// create new answer
	if err := r.db.Create(answer).Error; err != nil {
		return err
	}

	// add answer to question's answer
	if err := r.db.Model(&question).Association("Answers").Append(answer); err != nil {
		return err
	}

	return nil
}
