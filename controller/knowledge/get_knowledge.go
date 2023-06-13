package knowledge

import (
	"lms/service/s_knowledge"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type knowledgeGetController struct {
	knowledgeGetService s_knowledge.KnowledgeGetService
}

func NewKnowledgeGetController(knowledgeGetService s_knowledge.KnowledgeGetService) *knowledgeGetController {
	return &knowledgeGetController{
		knowledgeGetService: knowledgeGetService,
	}
}

func (qc *knowledgeGetController) GetAllKnowledge(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	knowledgeList, count, err := qc.knowledgeGetService.GetAllKnowledge(limit, page)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	totalPages := int(math.Ceil(float64(count) / float64(limit)))

	c.JSON(http.StatusOK, gin.H{
		"data":       knowledgeList,
		"count":      count,
		"page":       page,
		"totalPages": totalPages,
	})
}

func (qc *knowledgeGetController) GetByIdKnowledge(c *gin.Context) {
	IDKnowledge, err := strconv.Atoi(c.Param("id_knowledge"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid knowledge ID"})
		return
	}

	knowledge, err := qc.knowledgeGetService.GetByIdKnowledge(IDKnowledge)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": knowledge,
	})
}
