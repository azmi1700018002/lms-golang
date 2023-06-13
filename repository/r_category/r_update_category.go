package rcategory

import (
	"lms/config/db"
	"lms/model"
)

func MUpdateCategoryByID(idcategory int, input model.Category) (category model.Category, err error) {
	category = model.Category{
		Categoryname: input.Categoryname,
	}
	if err = db.Server().Where("idcategory = ?", idcategory).Updates(&category).Error; err != nil {
		return category, err
	}
	return category, nil
}
