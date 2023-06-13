package s_question

import (
	"lms/model"
	rquestion "lms/repository/r_question"
)

type QuestionGetService interface {
	GetAllQuestion(limit, page int) ([]model.Question, int64, error)
}

type questionGetService struct {
	questionRepo rquestion.QuestionGetRepository
}

func NewQuestionGetService(questionRepo rquestion.QuestionGetRepository) QuestionGetService {
	return &questionGetService{
		questionRepo: questionRepo,
	}
}

func (qs *questionGetService) GetAllQuestion(limit, page int) ([]model.Question, int64, error) {
	offset := (page - 1) * limit
	return qs.questionRepo.GetAllQuestion(limit, offset)
}
