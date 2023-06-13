package knowledge

import (
	"lms/model"
	"lms/service/s_knowledge"
	"net/http"

	"github.com/gin-gonic/gin"
)

type knowledgeController struct {
	knowledgeService s_knowledge.KnowledgeService
}

func NewKnowledgeController(knowledgeService s_knowledge.KnowledgeService) *knowledgeController {
	return &knowledgeController{
		knowledgeService: knowledgeService,
	}
}

func (c *knowledgeController) AddKnowledge(ctx *gin.Context) {
	var input model.Knowledge
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	knowledge, err := c.knowledgeService.AddKnowledge(input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": knowledge})
}
