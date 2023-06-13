package s_question

import (
	"lms/model"
	rquestion "lms/repository/r_question"
)

type QuestionUpdateService interface {
	UpdateQuestionByID(id_question int, input model.Question) (question model.Question, err error)
}

type questionUpdateService struct {
	questionUpdateRepo rquestion.QuestionUpdateRepository
}

func UpdateQuestionService(questionUpdateRepo rquestion.QuestionUpdateRepository) QuestionUpdateService {
	return &questionUpdateService{
		questionUpdateRepo: questionUpdateRepo,
	}
}

func (s *questionUpdateService) UpdateQuestionByID(id_question int, input model.Question) (question model.Question, err error) {
	return s.questionUpdateRepo.UpdateQuestionByID(id_question, input)
}
