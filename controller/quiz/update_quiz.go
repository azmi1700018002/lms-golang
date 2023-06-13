package quiz

import (
	"net/http"
	"strconv"

	"lms/model"
	"lms/service/s_quiz"
	"lms/validation/v_quiz"

	"github.com/gin-gonic/gin"
)

type QuizUpdateController interface {
	UpdateQuizByID(c *gin.Context)
}

type quizUpdateController struct {
	quizUpdateService s_quiz.QuizUpdateService
}

func UpdateQuizController(quizUpdateService s_quiz.QuizUpdateService) QuizUpdateController {
	return &quizUpdateController{
		quizUpdateService: quizUpdateService,
	}
}

func (c *quizUpdateController) UpdateQuizByID(ctx *gin.Context) {
	var quiz model.Quiz
	if err := ctx.ShouldBindJSON(&quiz); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := v_quiz.ValidateUpdateQuiz(quiz); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id_quiz, err := strconv.Atoi(ctx.Param("id_quiz"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedQuiz, err := c.quizUpdateService.UpdateQuizByID(id_quiz, quiz)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": updatedQuiz})
}
