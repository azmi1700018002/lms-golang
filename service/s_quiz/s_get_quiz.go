package s_quiz

import (
	"lms/model"
	rquiz "lms/repository/r_quiz"
)

type QuizGetService interface {
	GetAllQuiz(limit, page int) ([]model.Quiz, int64, error)
	GetByIdQuiz(id int) (model.Quiz, error)
}

type quizGetService struct {
	quizRepo rquiz.QuizGetRepository
}

func NewQuizGetService(quizRepo rquiz.QuizGetRepository) QuizGetService {
	return &quizGetService{
		quizRepo: quizRepo,
	}
}

func (qs *quizGetService) GetAllQuiz(limit, page int) ([]model.Quiz, int64, error) {
	offset := (page - 1) * limit
	return qs.quizRepo.GetAllQuiz(limit, offset)
}

func (qs *quizGetService) GetByIdQuiz(IDQuiz int) (model.Quiz, error) {
	return qs.quizRepo.GetByIdQuiz(IDQuiz)
}
