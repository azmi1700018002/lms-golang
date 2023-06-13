package category

import (
	rcategory "lms/repository/r_category"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteCategory(c *gin.Context) {
	idcategory, err := strconv.Atoi(c.Param("idcategory"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	err = rcategory.MDeleteCategoryByID(idcategory)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "Category deleted"})
}
