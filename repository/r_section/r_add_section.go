package rsection

import (
	"lms/config/db"
	"lms/model"
)


type SectionRepository interface {
	AddSection(input *model.Section) (*model.Section, error)
}

type sectionRepository struct{}

func NewSectionRepository() SectionRepository {
	return &sectionRepository{}
}

func (r *sectionRepository) AddSection(input *model.Section) (*model.Section, error) {
	section := model.Section{
		SectionTitle: input.SectionTitle,
		IDcourse:     input.IDcourse,
	}

	if err := section.Validate(); err != nil {
		return nil, err
	}

	if err := db.Server().Create(&section).Error; err != nil {
		return nil, err
	}

	return &section, nil
}
