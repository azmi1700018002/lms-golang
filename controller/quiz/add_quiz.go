package quiz

import (
	"lms/model"
	"lms/service/s_quiz"
	"net/http"

	"github.com/gin-gonic/gin"
)

type quizController struct {
	quizService s_quiz.QuizService
}

func NewQuizController(quizService s_quiz.QuizService) *quizController {
	return &quizController{
		quizService: quizService,
	}
}

func (c *quizController) AddQuiz(ctx *gin.Context) {
	var input model.Quiz
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	quiz, err := c.quizService.AddQuiz(ctx, &input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": quiz})
}
