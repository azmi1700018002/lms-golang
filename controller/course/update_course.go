package course

import (
	"lms/model"
	rcourse "lms/repository/r_course"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateCourse(c *gin.Context) {
	var input model.Course
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idcourse, err := strconv.Atoi(c.Param("idcourse"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	_, err = rcourse.MUpdateCourseByID(idcourse, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Course updated"})
}
