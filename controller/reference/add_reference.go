package reference

import (
	"lms/model"
	rreference "lms/repository/r_reference"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddReference(c *gin.Context) {
	var input model.Reference
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	references, err := rreference.MaddReference(input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": references})
}
