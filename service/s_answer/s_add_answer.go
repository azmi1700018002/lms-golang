package s_answer

import (
	"context"
	"errors"
	"lms/model"
	ranswer "lms/repository/r_answer"
	"lms/validation/v_answer"
)

type AnswerService interface {
	AddAnswer(ctx context.Context, input *model.Answer) (*model.Answer, error)
}

type answerService struct {
	answerRepo ranswer.AnswerRepository
}

func NewAnswerService(answerRepo ranswer.AnswerRepository) AnswerService {
	return &answerService{
		answerRepo: answerRepo,
	}
}

func (s *answerService) AddAnswer(ctx context.Context, input *model.Answer) (*model.Answer, error) {
	if err := v_answer.ValidateAnswer(*input); err != nil {
		return nil, err
	}
	if err := s.answerRepo.Create(ctx, input); err != nil {
		return nil, errors.New("failed to add answer")
	}
	return input, nil
}
