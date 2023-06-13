package s_knowledge

import (
	"lms/model"
	rknowledge "lms/repository/r_knowledge"
)

type KnowledgeUpdateService interface {
	UpdateKnowledgeByID(id_knowledge int, input model.Knowledge) (knowledge model.Knowledge, err error)
}

type knowledgeUpdateService struct {
	knowledgeUpdateRepo rknowledge.KnowledgeUpdateRepository
}

func UpdateKnowledgeService(knowledgeUpdateRepo rknowledge.KnowledgeUpdateRepository) KnowledgeUpdateService {
	return &knowledgeUpdateService{
		knowledgeUpdateRepo: knowledgeUpdateRepo,
	}
}

func (s *knowledgeUpdateService) UpdateKnowledgeByID(id_knowledge int, input model.Knowledge) (knowledge model.Knowledge, err error) {
	return s.knowledgeUpdateRepo.UpdateKnowledgeByID(id_knowledge, input)
}
