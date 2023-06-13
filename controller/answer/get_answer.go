package answer

import (
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	s_answer "lms/service/s_answer"
)

type answerGetController struct {
	answerGetService s_answer.AnswerGetService
}

func NewAnswerGetController(answerGetService s_answer.AnswerGetService) *answerGetController {
	return &answerGetController{
		answerGetService: answerGetService,
	}
}

func (qc *answerGetController) GetAllAnswer(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	answerList, count, err := qc.answerGetService.GetAllAnswer(limit, page)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	totalPages := int(math.Ceil(float64(count) / float64(limit)))

	c.JSON(http.StatusOK, gin.H{
		"data":       answerList,
		"count":      count,
		"page":       page,
		"totalPages": totalPages,
	})
}
