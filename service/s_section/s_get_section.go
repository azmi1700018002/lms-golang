package s_section

import (
	"lms/model"
	rsection "lms/repository/r_section"
)

type SectionGetService interface {
	GetAllSection(limit, page int) ([]model.Section, int64, error)
	GetByIdSection(id int) (model.Section, error)
}

type sectionGetService struct {
	sectionRepo rsection.SectionGetRepository
}

func NewSectionGetService(sectionRepo rsection.SectionGetRepository) SectionGetService {
	return &sectionGetService{
		sectionRepo: sectionRepo,
	}
}

func (qs *sectionGetService) GetAllSection(limit, page int) ([]model.Section, int64, error) {
	offset := (page - 1) * limit
	return qs.sectionRepo.GetAllSection(limit, offset)
}

func (qs *sectionGetService) GetByIdSection(IDSection int) (model.Section, error) {
	return qs.sectionRepo.GetByIdSection(IDSection)
}
