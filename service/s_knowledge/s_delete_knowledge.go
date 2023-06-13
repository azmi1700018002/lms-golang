package s_knowledge

import (
	rknowledge "lms/repository/r_knowledge"
)

type KnowledgeDeleteService interface {
	DeleteKnowledge(id int) error
}

type knowledgeDeleteService struct {
	knowledgeDeleteRepo rknowledge.KnowledgeDeleteRepository
}

func DeleteKnowledgeService(knowledgeDeleteRepo rknowledge.KnowledgeDeleteRepository) KnowledgeDeleteService {
	return &knowledgeDeleteService{knowledgeDeleteRepo}
}

func (s *knowledgeDeleteService) DeleteKnowledge(id int) error {
	err := s.knowledgeDeleteRepo.DeleteKnowledgeByID(id)
	if err != nil {
		return err
	}
	return nil
}
