package s_answer

import (
	"lms/model"
	ranswer "lms/repository/r_answer"
)

type AnswerUpdateService interface {
	UpdateAnswerByID(id_answer int, input model.Answer) (answer model.Answer, err error)
}

type answerUpdateService struct {
	answerUpdateRepo ranswer.AnswerUpdateRepository
}

func UpdateAnswerService(answerUpdateRepo ranswer.AnswerUpdateRepository) AnswerUpdateService {
	return &answerUpdateService{
		answerUpdateRepo: answerUpdateRepo,
	}
}

func (s *answerUpdateService) UpdateAnswerByID(id_answer int, input model.Answer) (answer model.Answer, err error) {
	return s.answerUpdateRepo.UpdateAnswerByID(id_answer, input)
}
