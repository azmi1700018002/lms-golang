package rcourse

import (
	"context"
	"lms/model"

	"gorm.io/gorm"
)

type CourseRepository interface {
	AddCourse(ctx context.Context, course *model.Course) error
}

type courseRepository struct {
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) CourseRepository {
	return &courseRepository{db: db}
}

func (r *courseRepository) AddCourse(ctx context.Context, course *model.Course) error {
	// find user by IDUser
	var user model.User
	if err := r.db.Where("id_user = ?", course.IDUser).First(&user).Error; err != nil {
		return err
	}

	// create new course
	if err := r.db.Create(&course).Error; err != nil {
		return err
	}

	// add course to user's courses
	if err := r.db.Model(&user).Association("Course").Append(course); err != nil {
		return err
	}

	return nil
}
