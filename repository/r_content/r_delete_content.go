package rcontent

import (
	"lms/config/db"
	"lms/model"
)

// func MDeleteContentByID(idcontent int) (err error) {
// 	if err = db.Server().Unscoped().Where("idcontent = ?", idcontent).Delete(&model.Content{}).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

type ContentDeleteRepository interface {
	DeleteContentByID(idcontent int) error
}

type contentDeleteRepository struct{}

func DeleteContentRepository() ContentDeleteRepository {
	return &contentDeleteRepository{}
}

func (*contentDeleteRepository) DeleteContentByID(idcontent int) error {
	if err := db.Server().Unscoped().Where("idcontent = ?", idcontent).Delete(&model.Content{}).Error; err != nil {
		return err
	}
	return nil
}
