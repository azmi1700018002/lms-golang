package s_knowledge

import (
	"lms/model"
	rknowledge "lms/repository/r_knowledge"
)

type KnowledgeGetService interface {
	GetAllKnowledge(limit, page int) ([]model.Knowledge, int64, error)
	GetByIdKnowledge(id int) (model.Knowledge, error)
}

type knowledgeGetService struct {
	knowledgeRepo rknowledge.KnowledgeGetRepository
}

func NewKnowledgeGetService(knowledgeRepo rknowledge.KnowledgeGetRepository) KnowledgeGetService {
	return &knowledgeGetService{
		knowledgeRepo: knowledgeRepo,
	}
}

func (qs *knowledgeGetService) GetAllKnowledge(limit, page int) ([]model.Knowledge, int64, error) {
	offset := (page - 1) * limit
	return qs.knowledgeRepo.GetAllKnowledge(limit, offset)
}

func (qs *knowledgeGetService) GetByIdKnowledge(IDKnowledge int) (model.Knowledge, error) {
	return qs.knowledgeRepo.GetByIdKnowledge(IDKnowledge)
}
