package s_category

import (
	"errors"
	"lms/model"
	rcategory "lms/repository/r_category"
	"lms/validation/v_category"
)

type CategoryService interface {
	AddCategory(input model.Category) (model.Category, error)
}

type categoryService struct {
	categoryRepo rcategory.CategoryRepository
}

func NewCategoryService(categoryRepo rcategory.CategoryRepository) CategoryService {
	return &categoryService{
		categoryRepo: categoryRepo,
	}
}

func (s *categoryService) AddCategory(input model.Category) (model.Category, error) {
	if err := v_category.ValidateCategory(input); err != nil {
		return model.Category{}, err
	}
	category, err := s.categoryRepo.Create(input)
	if err != nil {
		return model.Category{}, errors.New("failed to add knowledge")
	}
	return category, nil
}
