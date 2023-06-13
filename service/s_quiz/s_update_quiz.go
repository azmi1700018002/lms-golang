package s_quiz

import (
	"lms/model"
	rquiz "lms/repository/r_quiz"
)

type QuizUpdateService interface {
	UpdateQuizByID(id_quiz int, input model.Quiz) (quiz model.Quiz, err error)
}

type quizUpdateService struct {
	quizUpdateRepo rquiz.QuizUpdateRepository
}

func UpdateQuizService(quizUpdateRepo rquiz.QuizUpdateRepository) QuizUpdateService {
	return &quizUpdateService{
		quizUpdateRepo: quizUpdateRepo,
	}
}

func (s *quizUpdateService) UpdateQuizByID(id_quiz int, input model.Quiz) (quiz model.Quiz, err error) {
	return s.quizUpdateRepo.UpdateQuizByID(id_quiz, input)
}
