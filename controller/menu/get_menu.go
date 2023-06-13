package menu

import (
	"math"
	"net/http"
	"strconv"

	rmenu "lms/repository/r_menu"

	"github.com/gin-gonic/gin"
)

func GetMenu(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	offset := (page - 1) * limit

	menu, count, err := rmenu.MGetAllMenu(limit, offset)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	totalPages := int(math.Ceil(float64(count) / float64(limit)))

	c.JSON(http.StatusOK, gin.H{
		"data":       menu,
		"count":      count,
		"page":       page,
		"totalPages": totalPages,
	})
}
