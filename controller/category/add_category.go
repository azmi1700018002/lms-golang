package category

import (
	"lms/model"
	"lms/service/s_category"
	"net/http"

	"github.com/gin-gonic/gin"
)

type categoryController struct {
	categoryService s_category.CategoryService
}

func NewCategoryController(categoryService s_category.CategoryService) *categoryController {
	return &categoryController{
		categoryService: categoryService,
	}
}

func (c *categoryController) AddCategory(ctx *gin.Context) {
	var input model.Category
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	category, err := c.categoryService.AddCategory(input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": category})
}
