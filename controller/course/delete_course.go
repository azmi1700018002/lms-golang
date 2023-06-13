package course

import (
	rcourse "lms/repository/r_course"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteCourse(c *gin.Context) {
	idcourse, err := strconv.Atoi(c.Param("idcourse"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Course ID"})
		return
	}

	err = rcourse.MDeleteCourseByID(idcourse)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "Course deleted"})
}
