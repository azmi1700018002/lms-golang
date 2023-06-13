package role

import (
	rrole "lms/repository/r_role"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetRoles(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	offset := (page - 1) * limit

	roles, count, err := rrole.MGetAllRole(limit, offset)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	totalPages := int(math.Ceil(float64(count) / float64(limit)))

	c.JSON(http.StatusOK, gin.H{
		"data":       roles,
		"count":      count,
		"page":       page,
		"totalPages": totalPages,
	})
}
