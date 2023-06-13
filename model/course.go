package model

import (
	"time"

	"github.com/google/uuid"
)

type Course struct {
	IDcourse    int        `json:"id_course" gorm:"primaryKey"`
	IdKnowledge int        `json:"id_knowledge"`
	IDUser      uuid.UUID  `gorm:"type:uuid; column:id_user;"`
	CourseName  string     `json:"course_name"`
	CourseDesc  string     `json:"course_desc"`
	Datestart   *time.Time `json:"date_start"`
	Dateend     *time.Time `json:"date_end"`
	Sections    []Section  `json:"sections" gorm:"foreignKey:IDcourse"`
}
