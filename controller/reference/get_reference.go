package reference

import (
	rreference "lms/repository/r_reference"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetReference(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	offset := (page - 1) * limit

	references, count, err := rreference.MGetAllReference(limit, offset)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	totalPages := int(math.Ceil(float64(count) / float64(limit)))

	c.JSON(http.StatusOK, gin.H{
		"data":       references,
		"count":      count,
		"page":       page,
		"totalPages": totalPages,
	})
}
