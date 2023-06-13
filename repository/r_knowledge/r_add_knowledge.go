package rknowledge

import (
	"lms/config/db"
	"lms/model"
)

type KnowledgeRepository interface {
	Create(knowledge model.Knowledge) (model.Knowledge, error)
}

type knowledgeRepository struct{}

func NewKnowledgeRepository() KnowledgeRepository {
	return &knowledgeRepository{}
}

func (r *knowledgeRepository) Create(knowledge model.Knowledge) (model.Knowledge, error) {
	if err := db.Server().Create(&knowledge).Error; err != nil {
		return model.Knowledge{}, err
	}
	return knowledge, nil
}
