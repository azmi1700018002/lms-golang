package course

import (
	"lms/service/s_course"

	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type courseGetController struct {
	courseGetService s_course.CourseGetService
}

func NewCourseGetController(courseGetService s_course.CourseGetService) *courseGetController {
	return &courseGetController{
		courseGetService: courseGetService,
	}
}

func (qc *courseGetController) GetAllCourse(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	courseList, count, err := qc.courseGetService.GetAllCourse(limit, page)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	totalPages := int(math.Ceil(float64(count) / float64(limit)))

	c.JSON(http.StatusOK, gin.H{
		"data":       courseList,
		"count":      count,
		"page":       page,
		"totalPages": totalPages,
	})
}

func (qc *courseGetController) GetByIdCourse(c *gin.Context) {
	IDCourse, err := strconv.Atoi(c.Param("idcourse"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	course, err := qc.courseGetService.GetByIdCourse(IDCourse)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": course,
	})
}
