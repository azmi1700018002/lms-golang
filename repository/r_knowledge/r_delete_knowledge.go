package rknowledge

import (
	"lms/config/db"
	"lms/model"
)

type KnowledgeDeleteRepository interface {
	DeleteKnowledgeByID(id int) error
}

type knowledgeDeleteRepository struct{}

func DeleteKnowledgeRepository() KnowledgeDeleteRepository {
	return &knowledgeDeleteRepository{}
}

func (*knowledgeDeleteRepository) DeleteKnowledgeByID(id int) error {
	if err := db.Server().Unscoped().Where("id_knowledge = ?", id).Delete(&model.Knowledge{}).Error; err != nil {
		return err
	}
	return nil
}
