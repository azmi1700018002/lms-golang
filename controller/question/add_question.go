package question

import (
	"lms/model"
	"lms/service/s_question"
	"net/http"

	"github.com/gin-gonic/gin"
)

type questionController struct {
	questionService s_question.QuestionService
}

func NewQuestionController(questionService s_question.QuestionService) *questionController {
	return &questionController{
		questionService: questionService,
	}
}

func (c *questionController) AddQuestion(ctx *gin.Context) {
	var input model.Question
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	question, err := c.questionService.AddQuestion(ctx, &input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": question})
}
