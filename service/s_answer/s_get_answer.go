package s_answer

import (
	"lms/model"
	ranswer "lms/repository/r_answer"
)

type AnswerGetService interface {
	GetAllAnswer(limit, page int) ([]model.Answer, int64, error)
}

type answerGetService struct {
	answerRepo ranswer.AnswerGetRepository
}

func NewAnswerGetService(answerRepo ranswer.AnswerGetRepository) AnswerGetService {
	return &answerGetService{
		answerRepo: answerRepo,
	}
}

func (qs *answerGetService) GetAllAnswer(limit, page int) ([]model.Answer, int64, error) {
	offset := (page - 1) * limit
	return qs.answerRepo.GetAllAnswer(limit, offset)
}
