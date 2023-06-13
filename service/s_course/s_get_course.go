package s_course

import (
	"lms/model"
	rcourse "lms/repository/r_course"
)

type CourseGetService interface {
	GetAllCourse(limit, page int) ([]model.Course, int64, error)
	GetByIdCourse(id int) (model.Course, error)
}

type courseGetService struct {
	courseRepo rcourse.CourseGetRepository
}

func NewCourseGetService(courseRepo rcourse.CourseGetRepository) CourseGetService {
	return &courseGetService{
		courseRepo: courseRepo,
	}
}

func (qs *courseGetService) GetAllCourse(limit, page int) ([]model.Course, int64, error) {
	offset := (page - 1) * limit
	return qs.courseRepo.GetAllCourse(limit, offset)
}

func (qs *courseGetService) GetByIdCourse(IDCourse int) (model.Course, error) {
	return qs.courseRepo.GetByIdCourse(IDCourse)
}
