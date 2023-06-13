package s_quiz

import (
	"context"
	"lms/model"
	rquiz "lms/repository/r_quiz"
	"lms/validation/v_quiz"
)

type QuizService interface {
	AddQuiz(ctx context.Context, input *model.Quiz) (*model.Quiz, error)
}

type quizService struct {
	quizRepo rquiz.QuizRepository
}

func NewQuizService(quizRepo rquiz.QuizRepository) QuizService {
	return &quizService{
		quizRepo: quizRepo,
	}
}

func (s *quizService) AddQuiz(ctx context.Context, input *model.Quiz) (*model.Quiz, error) {
	if err := v_quiz.ValidateQuiz(*input); err != nil {
		return nil, err
	}
	if err := s.quizRepo.Create(ctx, input); err != nil {
		return nil, err
	}
	return input, nil
}
