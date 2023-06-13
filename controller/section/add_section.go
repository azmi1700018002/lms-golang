package section

import (
	"lms/model"
	"lms/service/s_section"
	"net/http"

	"github.com/gin-gonic/gin"
)

// func AddSection(c *gin.Context) {
// 	var input model.Section
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	section, err := repository.MaddSection(input)

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"data": section})
// }

type sectionController struct {
	sectionService s_section.SectionService
}

func NewSectionController(service s_section.SectionService) *sectionController {
	return &sectionController{
		sectionService: service,
	}
}

func (c *sectionController) AddSection(ctx *gin.Context) {
	var input model.Section
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	section, err := c.sectionService.AddSection(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": section})
}
