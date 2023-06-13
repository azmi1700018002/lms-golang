package s_quiz

import (
	rquiz "lms/repository/r_quiz"
)

type QuizDeleteService interface {
	DeleteQuiz(IDQuiz int) error
}

type quizDeleteService struct {
	quizDeleteRepo rquiz.QuizDeleteRepository
}

func DeleteQuizService(quizDeleteRepo rquiz.QuizDeleteRepository) QuizDeleteService {
	return &quizDeleteService{quizDeleteRepo}
}

func (s *quizDeleteService) DeleteQuiz(IDQuiz int) error {
	err := s.quizDeleteRepo.DeleteQuizByID(IDQuiz)
	if err != nil {
		return err
	}
	return nil
}
