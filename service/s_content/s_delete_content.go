package s_content

import (
	rcontent "lms/repository/r_content"
)

type ContentDeleteService interface {
	DeleteContent(id int) error
}

type contentDeleteService struct {
	contentDeleteRepo rcontent.ContentDeleteRepository
}

func DeleteContentService(contentDeleteRepo rcontent.ContentDeleteRepository) ContentDeleteService {
	return &contentDeleteService{contentDeleteRepo}
}

func (s *contentDeleteService) DeleteContent(id int) error {
	err := s.contentDeleteRepo.DeleteContentByID(id)
	if err != nil {
		return err
	}
	return nil
}
