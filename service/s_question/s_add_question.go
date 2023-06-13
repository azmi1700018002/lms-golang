package s_question

import (
	"context"
	"errors"
	"lms/model"
	rquestion "lms/repository/r_question"
	"lms/validation/v_question"
)

type QuestionService interface {
	AddQuestion(ctx context.Context, input *model.Question) (*model.Question, error)
}

type questionService struct {
	questionRepo rquestion.QuestionRepository
}

func NewQuestionService(questionRepo rquestion.QuestionRepository) QuestionService {
	return &questionService{
		questionRepo: questionRepo,
	}
}

func (s *questionService) AddQuestion(ctx context.Context, input *model.Question) (*model.Question, error) {
	if err := v_question.ValidateQuestion(*input); err != nil {
		return nil, err
	}
	if err := s.questionRepo.Create(ctx, input); err != nil {
		return nil, errors.New("failed to add question")
	}
	return input, nil
}
