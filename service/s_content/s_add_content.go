package s_content

import (
	"lms/model"
)

type ContentRepository interface {
	CreateContent(input model.Content) (model.Content, error)
}

type ContentService struct {
	contentRepo ContentRepository
}

func NewContentService(contentRepo ContentRepository) *ContentService {
	return &ContentService{contentRepo: contentRepo}
}

func (s *ContentService) AddContent(input model.Content) (content model.Content, err error) {
	content, err = s.contentRepo.CreateContent(input)
	if err != nil {
		return model.Content{}, err
	}
	return content, nil
}
