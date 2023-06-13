package section

import (
	"lms/service/s_section"

	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type sectionGetController struct {
	sectionGetService s_section.SectionGetService
}

func NewSectionGetController(sectionGetService s_section.SectionGetService) *sectionGetController {
	return &sectionGetController{
		sectionGetService: sectionGetService,
	}
}

func (qc *sectionGetController) GetAllSection(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	sectionList, count, err := qc.sectionGetService.GetAllSection(limit, page)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	totalPages := int(math.Ceil(float64(count) / float64(limit)))

	c.JSON(http.StatusOK, gin.H{
		"data":       sectionList,
		"count":      count,
		"page":       page,
		"totalPages": totalPages,
	})
}

func (qc *sectionGetController) GetByIdSection(c *gin.Context) {
	IDSection, err := strconv.Atoi(c.Param("idsection"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid section ID"})
		return
	}

	section, err := qc.sectionGetService.GetByIdSection(IDSection)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": section,
	})
}
