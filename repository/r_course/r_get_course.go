package rcourse

import (
	"fmt"
	"lms/config/db"
	"lms/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CourseGetRepository interface {
	GetAllCourse(limit, offset int) ([]model.Course, int64, error)
	GetByIdCourse(id int) (model.Course, error)
}

type courseGetRepository struct{}

func NewCourseGetRepository() CourseGetRepository {
	return &courseGetRepository{}
}

func (qr *courseGetRepository) GetAllCourse(limit, offset int) ([]model.Course, int64, error) {
	// Query to select course rows with limit and offset
	querySelect := `
		SELECT idcourse, id_knowledge, id_user, course_name, course_desc, datestart, dateend
		FROM courses
		LIMIT $1
		OFFSET $2
	`
	// Execute the query and scan the rows into a slice of Course structs
	rows, err := db.Server().Raw(querySelect, limit, offset).Rows()
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var courseList []model.Course
	for rows.Next() {
		var course model.Course
		if err := rows.Scan(&course.IDcourse, &course.IdKnowledge, &course.IDUser, &course.CourseName, &course.CourseDesc, &course.Datestart, &course.Dateend); err != nil {
			return nil, 0, err
		}

		// Eager loading questions
		if err := db.Server().Where("idcourse = ?", course.IDcourse).Preload("Sections", func(db *gorm.DB) *gorm.DB {
			return db.Preload("Content")
		}).Find(&course).Error; err != nil {
			return nil, 0, err
		}

		courseList = append(courseList, course)
	}

	// Query to count the total number of course rows
	queryCount := `
		SELECT COUNT(*)
		FROM courses
	`
	var count int64
	if err := db.Server().Raw(queryCount).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	return courseList, count, nil
}

func (qr *courseGetRepository) GetByIdCourse(IDCourse int) (model.Course, error) {
	var course model.Course
	if err := db.Server().Where("idcourse = ?", IDCourse).Preload("Sections", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Content")
	}).First(&course).Error; err != nil {
		return course, err
	}
	return course, nil
}

func MGetCourseByID(userID uuid.UUID, page int, limit int) ([]model.Course, error) {
	var courses []model.Course
	offset := (page - 1) * limit

	stmt := `
		SELECT idcourse, id_knowledge, id_user, course_name, course_desc, datestart, dateend
		FROM courses
		WHERE id_user = $1
		LIMIT $2
		OFFSET $3;
	`
	rows, err := db.Server().Raw(stmt, userID, limit, offset).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var course model.Course
		err := rows.Scan(&course.IDcourse, &course.IdKnowledge, &course.IDUser, &course.CourseName, &course.CourseDesc, &course.Datestart, &course.Dateend)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}

	if len(courses) == 0 {
		return nil, fmt.Errorf("no courses found")
	}

	return courses, nil
}
