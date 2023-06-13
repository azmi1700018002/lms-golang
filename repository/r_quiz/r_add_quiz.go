package rquiz

import (
	"context"
	"lms/model"

	"gorm.io/gorm"
)

type QuizRepository interface {
	Create(ctx context.Context, quiz *model.Quiz) error
}

type quizRepository struct {
	db *gorm.DB
}

func NewQuizRepository(db *gorm.DB) QuizRepository {
	return &quizRepository{db: db}
}

func (r *quizRepository) Create(ctx context.Context, quiz *model.Quiz) error {

	// find user by IDUser
	var user model.User
	if err := r.db.Where("id_user = ?", quiz.IDUser).First(&user).Error; err != nil {
		return err
	}

	// create new quiz
	if err := r.db.Create(quiz).Error; err != nil {
		return err
	}

	// add quiz to user's quizzez
	//Note : Association melihat dari model example Quizzes []Quiz
	if err := r.db.Model(&user).Association("Quizzes").Append(quiz); err != nil {
		return err
	}

	return nil
}
