package rcategory

import (
	"lms/config/db"
	"lms/model"
)

type CategoryRepository interface {
	Create(category model.Category) (model.Category, error)
}

type categoryRepository struct{}

func NewCategoryRepository() CategoryRepository {
	return &categoryRepository{}
}

func (r *categoryRepository) Create(category model.Category) (model.Category, error) {
	if err := db.Server().Create(&category).Error; err != nil {
		return model.Category{}, err
	}
	return category, nil
}
