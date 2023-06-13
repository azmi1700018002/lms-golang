package quiz

import (
	"lms/service/s_quiz"
	"lms/validation/v_quiz"
	"net/http"

	"github.com/gin-gonic/gin"
)

type quizDeleteController struct {
	quizDeleteService s_quiz.QuizDeleteService
}

func DeleteQuizController(quizDeleteService s_quiz.QuizDeleteService) *quizDeleteController {
	return &quizDeleteController{quizDeleteService}
}

func (c *quizDeleteController) DeleteQuiz(ctx *gin.Context) {
	id := ctx.Param("id_quiz")
	IDQuiz, err := v_quiz.ValidateID(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = c.quizDeleteService.DeleteQuiz(IDQuiz)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": "Quiz deleted"})
}
