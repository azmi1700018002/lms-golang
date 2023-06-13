package s_knowledge

import (
	"errors"
	"lms/model"
	rknowledge "lms/repository/r_knowledge"
	"lms/validation/v_knowledge"
)

type KnowledgeService interface {
	AddKnowledge(input model.Knowledge) (model.Knowledge, error)
}

type knowledgeService struct {
	knowledgeRepo rknowledge.KnowledgeRepository
}

func NewKnowledgeService(knowledgeRepo rknowledge.KnowledgeRepository) KnowledgeService {
	return &knowledgeService{
		knowledgeRepo: knowledgeRepo,
	}
}

func (s *knowledgeService) AddKnowledge(input model.Knowledge) (model.Knowledge, error) {
	if err := v_knowledge.ValidateKnowledge(input); err != nil {
		return model.Knowledge{}, err
	}
	knowledge, err := s.knowledgeRepo.Create(input)
	if err != nil {
		return model.Knowledge{}, errors.New("failed to add knowledge")
	}
	return knowledge, nil
}
