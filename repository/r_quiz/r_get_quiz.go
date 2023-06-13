package rquiz

import (
	"lms/config/db"
	"lms/model"

	"gorm.io/gorm"
)

type QuizGetRepository interface {
	GetAllQuiz(limit, offset int) ([]model.Quiz, int64, error)
	GetByIdQuiz(id int) (model.Quiz, error)
}

type quizGetRepository struct{}

func NewQuizGetRepository() QuizGetRepository {
	return &quizGetRepository{}
}

func (qr *quizGetRepository) GetAllQuiz(limit, offset int) ([]model.Quiz, int64, error) {
	querySelect := `
        SELECT id_quiz, id_user, quiz_name, quiz_desc, quiz_type, created_at, score, date_start, date_end
        FROM quizzes
        LIMIT $1
        OFFSET $2
    `
	rows, err := db.Server().Raw(querySelect, limit, offset).Rows()
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var quizList []model.Quiz
	for rows.Next() {
		var quiz model.Quiz
		if err := rows.Scan(&quiz.IDQuiz, &quiz.IDUser, &quiz.QuizName, &quiz.QuizDesc, &quiz.QuizType, &quiz.CreatedAt, &quiz.Score, &quiz.DateStart, &quiz.DateEnd); err != nil {
			return nil, 0, err
		}

		// Eager loading questions
		if err := db.Server().Where("id_quiz = ?", quiz.IDQuiz).Preload("Questions", func(db *gorm.DB) *gorm.DB {
			return db.Preload("Answers")
		}).Find(&quiz).Error; err != nil {
			return nil, 0, err
		}

		quizList = append(quizList, quiz)
	}

	queryCount := `
        SELECT COUNT(*)
        FROM quizzes
    `
	var count int64
	if err := db.Server().Raw(queryCount).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	return quizList, count, nil
}

func (qr *quizGetRepository) GetByIdQuiz(IDQuiz int) (model.Quiz, error) {
	var quiz model.Quiz
	if err := db.Server().Where("id_quiz = ?", IDQuiz).Preload("Questions", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Answers")
	}).First(&quiz).Error; err != nil {
		return quiz, err
	}
	return quiz, nil
}
