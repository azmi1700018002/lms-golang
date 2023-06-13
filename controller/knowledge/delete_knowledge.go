package knowledge

import (
	"lms/service/s_knowledge"
	"lms/validation/v_knowledge"
	"net/http"

	"github.com/gin-gonic/gin"
)

type knowledgeDeleteController struct {
	knowledgeDeleteService s_knowledge.KnowledgeDeleteService
}

func DeleteKnowledgeController(knowledgeDeleteService s_knowledge.KnowledgeDeleteService) *knowledgeDeleteController {
	return &knowledgeDeleteController{knowledgeDeleteService}
}

func (c *knowledgeDeleteController) DeleteKnowledge(ctx *gin.Context) {
	idStr := ctx.Param("id_knowledge")
	id, err := v_knowledge.ValidateID(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = c.knowledgeDeleteService.DeleteKnowledge(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": "Knowledge deleted"})
}
