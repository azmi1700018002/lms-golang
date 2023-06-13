package knowledge

import (
	"net/http"
	"strconv"

	"lms/model"
	"lms/service/s_knowledge"
	"lms/validation/v_knowledge"

	"github.com/gin-gonic/gin"
)

type KnowledgeUpdateController interface {
	UpdateKnowledgeByID(c *gin.Context)
}

type knowledgeUpdateController struct {
	knowledgeUpdateService s_knowledge.KnowledgeUpdateService
}

func UpdateKnowledgeController(knowledgeUpdateService s_knowledge.KnowledgeUpdateService) KnowledgeUpdateController {
	return &knowledgeUpdateController{
		knowledgeUpdateService: knowledgeUpdateService,
	}
}

func (c *knowledgeUpdateController) UpdateKnowledgeByID(ctx *gin.Context) {
	var knowledge model.Knowledge
	if err := ctx.ShouldBindJSON(&knowledge); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := v_knowledge.ValidateUpdateKnowledge(knowledge); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id_knowledge, err := strconv.Atoi(ctx.Param("id_knowledge"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedKnowledge, err := c.knowledgeUpdateService.UpdateKnowledgeByID(id_knowledge, knowledge)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": updatedKnowledge})
}
