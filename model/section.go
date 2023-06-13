package model

import "github.com/go-playground/validator/v10"

type Section struct {
	IDsection    int       `json:"idsection" gorm:"primaryKey"`
	IDcourse     int       `json:"idcourse"`
	SectionTitle string    `json:"section_title" validate:"required,min=3,max=50"`
	Content      []Content `gorm:"foreignKey:IDsection"`
	Quizzes      []Quiz    `gorm:"many2many:quiz_sections;"`
	// CourseSections   []CourseSections   `gorm:"foreignKey:IDsection"`
	// KnowledgeSection []KnowledgeSection `gorm:"foreignKey:IDsection"`
}

func (s *Section) Validate() error {
	validate := validator.New()
	return validate.Struct(s)
}
