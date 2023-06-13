package s_answer

import (
	ranswer "lms/repository/r_answer"
)

type AnswerDeleteService interface {
	DeleteAnswer(IDAnswer int) error
}

type answerDeleteService struct {
	answerDeleteRepo ranswer.AnswerDeleteRepository
}

func DeleteAnswerService(answerDeleteRepo ranswer.AnswerDeleteRepository) AnswerDeleteService {
	return &answerDeleteService{answerDeleteRepo}
}

func (s *answerDeleteService) DeleteAnswer(IDAnswer int) error {
	err := s.answerDeleteRepo.DeleteAnswerByID(IDAnswer)
	if err != nil {
		return err
	}
	return nil
}
