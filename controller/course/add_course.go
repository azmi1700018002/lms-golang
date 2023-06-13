package course

import (
	"lms/model"
	rcourse "lms/repository/r_course"
	"net/http"

	"github.com/gin-gonic/gin"
)

type courseController struct {
	repo rcourse.CourseRepository
}

func NewCourseController(repo rcourse.CourseRepository) *courseController {
	return &courseController{repo}
}

func (c *courseController) AddCourse(ctx *gin.Context) {
	var course model.Course

	if err := ctx.ShouldBindJSON(&course); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.repo.AddCourse(ctx, &course); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, course)
}
