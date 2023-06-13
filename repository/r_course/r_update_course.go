package rcourse

import (
	"lms/config/db"
	"lms/model"
)

func MUpdateCourseByID(idcourse int, input model.Course) (course model.Course, err error) {
	course = model.Course{
		CourseName: input.CourseName,
		CourseDesc: input.CourseDesc,
	}
	if err = db.Server().Where("idcourse = ?", idcourse).Updates(&course).Error; err != nil {
		return course, err
	}
	return course, nil
}
