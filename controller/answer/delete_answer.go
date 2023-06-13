package answer

import (
	"lms/service/s_answer"
	"lms/validation/v_answer"
	"net/http"

	"github.com/gin-gonic/gin"
)

type answerDeleteController struct {
	answerDeleteService s_answer.AnswerDeleteService
}

func DeleteAnswerController(answerDeleteService s_answer.AnswerDeleteService) *answerDeleteController {
	return &answerDeleteController{answerDeleteService}
}

func (c *answerDeleteController) DeleteAnswer(ctx *gin.Context) {
	id := ctx.Param("id_answer")
	IDAnswer, err := v_answer.ValidateID(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = c.answerDeleteService.DeleteAnswer(IDAnswer)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": "Answer deleted"})
}
