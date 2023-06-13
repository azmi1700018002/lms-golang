package rcategory

import (
	"lms/config/db"
	"lms/model"
)

func MDeleteCategoryByID(idcategory int) (err error) {
	if err = db.Server().Unscoped().Where("idcategory = ?", idcategory).Delete(&model.Category{}).Error; err != nil {
		return err
	}
	return nil
}
