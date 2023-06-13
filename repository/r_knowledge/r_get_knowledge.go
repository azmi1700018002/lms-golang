package rknowledge

import (
	"lms/config/db"
	"lms/model"

	"gorm.io/gorm"
)

type KnowledgeGetRepository interface {
	GetAllKnowledge(limit, offset int) ([]model.Knowledge, int64, error)
	GetByIdKnowledge(id int) (model.Knowledge, error)
}

type knowledgeGetRepository struct{}

func NewKnowledgeGetRepository() KnowledgeGetRepository {
	return &knowledgeGetRepository{}
}

func (qr *knowledgeGetRepository) GetAllKnowledge(limit, offset int) ([]model.Knowledge, int64, error) {
	// Query to select knowledge rows with limit and offset
	querySelect := `
		SELECT id_knowledge, idcategory, knowledge_name, description, status
		FROM knowledges
		LIMIT $1
		OFFSET $2
	`
	// Execute the query and scan the rows into a slice of Knowledge structs
	rows, err := db.Server().Raw(querySelect, limit, offset).Rows()
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var knowledgeList []model.Knowledge
	for rows.Next() {
		var knowledge model.Knowledge
		if err := rows.Scan(&knowledge.IdKnowledge, &knowledge.IDcategory, &knowledge.KnowledgeName, &knowledge.Description, &knowledge.Status); err != nil {
			return nil, 0, err
		}

		// Eager loading questions
		if err := db.Server().Where("id_knowledge = ?", knowledge.IdKnowledge).Preload("Courses", func(db *gorm.DB) *gorm.DB {
			return db.Preload("Sections")
		}).Find(&knowledge).Error; err != nil {
			return nil, 0, err
		}

		knowledgeList = append(knowledgeList, knowledge)
	}

	// Query to count the total number of knowledge rows
	queryCount := `
		SELECT COUNT(*)
		FROM knowledges
	`
	var count int64
	if err := db.Server().Raw(queryCount).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	return knowledgeList, count, nil
}

func (qr *knowledgeGetRepository) GetByIdKnowledge(IDKnowledge int) (model.Knowledge, error) {
	var knowledge model.Knowledge
	if err := db.Server().Where("id_knowledge = ?", IDKnowledge).Preload("Courses", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Sections")
	}).First(&knowledge).Error; err != nil {
		return knowledge, err
	}
	return knowledge, nil
}
