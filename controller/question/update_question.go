package question

import (
	"net/http"
	"strconv"

	"lms/model"
	"lms/service/s_question"
	"lms/validation/v_question"

	"github.com/gin-gonic/gin"
)

type QuestionUpdateController interface {
	UpdateQuestionByID(c *gin.Context)
}

type questionUpdateController struct {
	questionUpdateService s_question.QuestionUpdateService
}

func UpdateQuestionController(questionUpdateService s_question.QuestionUpdateService) QuestionUpdateController {
	return &questionUpdateController{
		questionUpdateService: questionUpdateService,
	}
}

func (c *questionUpdateController) UpdateQuestionByID(ctx *gin.Context) {
	var question model.Question
	if err := ctx.ShouldBindJSON(&question); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := v_question.ValidateUpdateQuestion(question); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id_question, err := strconv.Atoi(ctx.Param("id_question"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedQuestion, err := c.questionUpdateService.UpdateQuestionByID(id_question, question)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": updatedQuestion})
}
