package answer

import (
	"net/http"
	"strconv"

	"lms/model"
	"lms/service/s_answer"
	"lms/validation/v_answer"

	"github.com/gin-gonic/gin"
)

type AnswerUpdateController interface {
	UpdateAnswerByID(c *gin.Context)
}

type answerUpdateController struct {
	answerUpdateService s_answer.AnswerUpdateService
}

func UpdateAnswerController(answerUpdateService s_answer.AnswerUpdateService) AnswerUpdateController {
	return &answerUpdateController{
		answerUpdateService: answerUpdateService,
	}
}

func (c *answerUpdateController) UpdateAnswerByID(ctx *gin.Context) {
	var answer model.Answer
	if err := ctx.ShouldBindJSON(&answer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := v_answer.ValidateUpdateAnswer(answer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id_answer, err := strconv.Atoi(ctx.Param("id_answer"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedAnswer, err := c.answerUpdateService.UpdateAnswerByID(id_answer, answer)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": updatedAnswer})
}
