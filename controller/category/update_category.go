package category

import (
	"lms/model"
	rcategory "lms/repository/r_category"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateCategory(c *gin.Context) {
	var input model.Category
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(c.Param("idcategory"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	_, err = rcategory.MUpdateCategoryByID(id, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category updated"})
}
