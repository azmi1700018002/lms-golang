package s_question

import (
	rquestion "lms/repository/r_question"
)

type QuestionDeleteService interface {
	DeleteQuestion(IDQuestion int) error
}

type questionDeleteService struct {
	questionDeleteRepo rquestion.QuestionDeleteRepository
}

func DeleteQuestionService(questionDeleteRepo rquestion.QuestionDeleteRepository) QuestionDeleteService {
	return &questionDeleteService{questionDeleteRepo}
}

func (s *questionDeleteService) DeleteQuestion(IDQuestion int) error {
	err := s.questionDeleteRepo.DeleteQuestionByID(IDQuestion)
	if err != nil {
		return err
	}
	return nil
}
