package question

import (
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	s_question "lms/service/s_question"
)

type questionGetController struct {
	questionGetService s_question.QuestionGetService
}

func NewQuestionGetController(questionGetService s_question.QuestionGetService) *questionGetController {
	return &questionGetController{
		questionGetService: questionGetService,
	}
}

func (qc *questionGetController) GetAllQuestion(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	questionList, count, err := qc.questionGetService.GetAllQuestion(limit, page)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	totalPages := int(math.Ceil(float64(count) / float64(limit)))

	c.JSON(http.StatusOK, gin.H{
		"data":       questionList,
		"count":      count,
		"page":       page,
		"totalPages": totalPages,
	})
}
