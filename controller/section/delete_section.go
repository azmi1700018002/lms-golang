package section

import (
	rsection "lms/repository/r_section"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteSection(c *gin.Context) {
	idsection, err := strconv.Atoi(c.Param("idsection"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid section ID"})
		return
	}

	err = rsection.MDeleteSectionByID(idsection)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "Section deleted"})
}
