package quiz

import (
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	s_quiz "lms/service/s_quiz"
)

type quizGetController struct {
	quizGetService s_quiz.QuizGetService
}

func NewQuizGetController(quizGetService s_quiz.QuizGetService) *quizGetController {
	return &quizGetController{
		quizGetService: quizGetService,
	}
}

func (qc *quizGetController) GetAllQuiz(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	quizList, count, err := qc.quizGetService.GetAllQuiz(limit, page)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	totalPages := int(math.Ceil(float64(count) / float64(limit)))

	c.JSON(http.StatusOK, gin.H{
		"data":       quizList,
		"count":      count,
		"page":       page,
		"totalPages": totalPages,
	})
}

func (qc *quizGetController) GetByIdQuiz(c *gin.Context) {
	IDQuiz, err := strconv.Atoi(c.Param("id_quiz"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid quiz ID"})
		return
	}

	quiz, err := qc.quizGetService.GetByIdQuiz(IDQuiz)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": quiz,
	})
}
