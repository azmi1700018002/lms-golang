package category

import (
	rcategory "lms/repository/r_category"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCategory(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	offset := (page - 1) * limit

	category, count, err := rcategory.MGetAllCategory(limit, offset)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	totalPages := int(math.Ceil(float64(count) / float64(limit)))

	c.JSON(http.StatusOK, gin.H{
		"data":       category,
		"count":      count,
		"page":       page,
		"totalPages": totalPages,
	})
}
