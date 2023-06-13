package rsection

import (
	"lms/config/db"
	"lms/model"
)

type SectionGetRepository interface {
	GetAllSection(limit, offset int) ([]model.Section, int64, error)
	GetByIdSection(id int) (model.Section, error)
}

type sectionGetRepository struct{}

func NewSectionGetRepository() SectionGetRepository {
	return &sectionGetRepository{}
}

func (qr *sectionGetRepository) GetAllSection(limit, offset int) ([]model.Section, int64, error) {
	// Query to select section rows with limit and offset
	querySelect := `
			SELECT idsection, COALESCE(idcourse, 0), section_title
		FROM sections
		LIMIT $1
		OFFSET $2
	`
	// Execute the query and scan the rows into a slice of Section structs
	rows, err := db.Server().Raw(querySelect, limit, offset).Rows()
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var sectionList []model.Section
	for rows.Next() {
		var section model.Section
		if err := rows.Scan(&section.IDsection, &section.IDcourse, &section.SectionTitle); err != nil {
			return nil, 0, err
		}

		// Eager loading questions
		if err := db.Server().Where("idsection = ?", section.IDsection).Preload("Content").Find(&section).Error; err != nil {
			return nil, 0, err
		}

		sectionList = append(sectionList, section)
	}

	// Query to count the total number of section rows
	queryCount := `
		SELECT COUNT(*)
		FROM sections
	`
	var count int64
	if err := db.Server().Raw(queryCount).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	return sectionList, count, nil
}

func (qr *sectionGetRepository) GetByIdSection(IDSection int) (model.Section, error) {
	var section model.Section
	if err := db.Server().Where("idsection = ?", IDSection).Preload("Content").First(&section).Error; err != nil {
		return section, err
	}
	return section, nil
}
