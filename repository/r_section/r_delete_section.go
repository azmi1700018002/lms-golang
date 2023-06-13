package rsection

import (
	"lms/config/db"
	"lms/model"
)

func MDeleteSectionByID(idsection int) (err error) {
	if err = db.Server().Unscoped().Where("idsection = ?", idsection).Delete(&model.Section{}).Error; err != nil {
		return err
	}
	return nil
}
