package section

import (
	"lms/model"
	rsection "lms/repository/r_section"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateSection(c *gin.Context) {
	var input model.Section
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idsection, err := strconv.Atoi(c.Param("idsection"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid section ID"})
		return
	}

	_, err = rsection.MUpdateSectionByID(idsection, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Section updated"})
}
