package answer

import (
	"lms/model"
	"lms/service/s_answer"
	"net/http"

	"github.com/gin-gonic/gin"
)

type answerController struct {
	answerService s_answer.AnswerService
}

func NewAnswerController(answerService s_answer.AnswerService) *answerController {
	return &answerController{
		answerService: answerService,
	}
}

func (c *answerController) AddAnswer(ctx *gin.Context) {
	var input model.Answer
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	answer, err := c.answerService.AddAnswer(ctx, &input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": answer})
}
