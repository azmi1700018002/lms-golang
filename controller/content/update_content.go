package content

import (
	"lms/model"
	rcontent "lms/repository/r_content"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateContent(c *gin.Context) {
	var input model.Content
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idcontent, err := strconv.Atoi(c.Param("idcontent"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Content ID"})
		return
	}

	_, err = rcontent.MUpdateContentByID(idcontent, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Content updated"})
}
