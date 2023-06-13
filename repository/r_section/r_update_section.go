package rsection

import (
	"lms/config/db"
	"lms/model"
)

func MUpdateSectionByID(idsection int, input model.Section) (section model.Section, err error) {
	section = model.Section{
		SectionTitle: input.SectionTitle,
	}
	if err = db.Server().Where("idsection = ?", idsection).Updates(&section).Error; err != nil {
		return section, err
	}
	return section, nil
}
