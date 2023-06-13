package question

import (
	"lms/service/s_question"
	"lms/validation/v_question"
	"net/http"

	"github.com/gin-gonic/gin"
)

type questionDeleteController struct {
	questionDeleteService s_question.QuestionDeleteService
}

func DeleteQuestionController(questionDeleteService s_question.QuestionDeleteService) *questionDeleteController {
	return &questionDeleteController{questionDeleteService}
}

func (c *questionDeleteController) DeleteQuestion(ctx *gin.Context) {
	id := ctx.Param("id_question")
	IDQuestion, err := v_question.ValidateID(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = c.questionDeleteService.DeleteQuestion(IDQuestion)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": "Question deleted"})
}
