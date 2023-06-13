package rknowledge

import (
	"lms/config/db"
	"lms/model"
)

type KnowledgeUpdateRepository interface {
	UpdateKnowledgeByID(id_knowledge int, input model.Knowledge) (knowledge model.Knowledge, err error)
}

type knowledgeUpdateRepository struct{}

func UpdateKnowledgeRepository() KnowledgeUpdateRepository {
	return &knowledgeUpdateRepository{}
}

func (r *knowledgeUpdateRepository) UpdateKnowledgeByID(id_knowledge int, input model.Knowledge) (knowledge model.Knowledge, err error) {
	// cek apakah data dengan id_knowledge yang ingin diupdate ada di database
	var existingKnowledge model.Knowledge
	if err := db.Server().First(&existingKnowledge, id_knowledge).Error; err != nil {
		return knowledge, err
	}

	knowledge = model.Knowledge{
		IdKnowledge:   id_knowledge,
		IDcategory:    input.IDcategory,
		KnowledgeName: input.KnowledgeName,
		Description:   input.Description,
		Status:        input.Status,
	}

	if err = db.Server().Where("id_knowledge = ?", id_knowledge).Updates(&knowledge).Error; err != nil {
		return knowledge, err
	}

	return knowledge, nil
}
