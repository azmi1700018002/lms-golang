package rquestion

import (
	"context"
	"lms/model"

	"gorm.io/gorm"
)

type QuestionRepository interface {
	Create(ctx context.Context, question *model.Question) error
}

type questionRepository struct {
	db *gorm.DB
}

func NewQuestionRepository(db *gorm.DB) QuestionRepository {
	return &questionRepository{db: db}
}

func (r *questionRepository) Create(ctx context.Context, question *model.Question) error {

	// find quiz by id_quiz
	var quiz model.Quiz
	if err := r.db.Where("id_quiz = ?", question.IDQuiz).First(&quiz).Error; err != nil {
		return err
	}

	// create new question
	if err := r.db.Create(question).Error; err != nil {
		return err
	}

	// add question to quiz's questions
	if err := r.db.Model(&quiz).Association("Questions").Append(question); err != nil {
		return err
	}

	return nil
}
