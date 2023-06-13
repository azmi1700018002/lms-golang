package s_section

import (
	"lms/model"
	rsection "lms/repository/r_section"
)

type SectionService interface {
	AddSection(input *model.Section) (*model.Section, error)
}

type sectionService struct {
	sectionRepo rsection.SectionRepository
}

func NewSectionService(repo rsection.SectionRepository) SectionService {
	return &sectionService{
		sectionRepo: repo,
	}
}

func (s *sectionService) AddSection(input *model.Section) (*model.Section, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	section, err := s.sectionRepo.AddSection(input)
	if err != nil {
		return nil, err
	}

	return section, nil
}
