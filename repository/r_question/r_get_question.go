package rquestion

import (
	"lms/config/db"
	"lms/model"
)

type QuestionGetRepository interface {
	GetAllQuestion(limit, offset int) ([]model.Question, int64, error)
}

type questionGetRepository struct{}

func NewQuestionGetRepository() QuestionGetRepository {
	return &questionGetRepository{}
}

func (qr *questionGetRepository) GetAllQuestion(limit, offset int) ([]model.Question, int64, error) {
	querySelect := `
        SELECT id_question, id_quiz, question_name, is_correct, created_at, updated_at
        FROM questions
        LIMIT $1
        OFFSET $2
    `
	rows, err := db.Server().Raw(querySelect, limit, offset).Rows()
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var questionList []model.Question
	for rows.Next() {
		var question model.Question
		if err := rows.Scan(&question.IDQuestion, &question.IDQuiz, &question.QuestionName, &question.IsCorrect, &question.CreatedAt, &question.UpdatedAt); err != nil {
			return nil, 0, err
		}

		// Eager loading answer
		if err := db.Server().Where("id_question = ?", question.IDQuestion).Preload("Answers").Find(&question).Error; err != nil {
			return nil, 0, err
		}
		questionList = append(questionList, question)
	}

	queryCount := `
        SELECT COUNT(*)
        FROM questions
    `
	var count int64
	if err := db.Server().Raw(queryCount).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	return questionList, count, nil
}
