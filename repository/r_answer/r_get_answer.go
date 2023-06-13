package ranswer

import (
	"lms/config/db"
	"lms/model"
)

type AnswerGetRepository interface {
	GetAllAnswer(limit, offset int) ([]model.Answer, int64, error)
}

type answerGetRepository struct{}

func NewAnswerGetRepository() AnswerGetRepository {
	return &answerGetRepository{}
}

func (qr *answerGetRepository) GetAllAnswer(limit, offset int) ([]model.Answer, int64, error) {
	querySelect := `
        SELECT id_answer, id_question, answer_text, is_correct
        FROM answers
        LIMIT $1
        OFFSET $2
    `
	rows, err := db.Server().Raw(querySelect, limit, offset).Rows()
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var answerList []model.Answer
	for rows.Next() {
		var answer model.Answer
		if err := rows.Scan(&answer.IDAnswer, &answer.IDQuestion, &answer.AnswerText, &answer.IsCorrect); err != nil {
			return nil, 0, err
		}
		answerList = append(answerList, answer)
	}

	queryCount := `
        SELECT COUNT(*)
        FROM answers
    `
	var count int64
	if err := db.Server().Raw(queryCount).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	return answerList, count, nil
}
