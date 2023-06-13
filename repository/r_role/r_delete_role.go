package rrole

import (
	"lms/config/db"
	"lms/model"
)

func MDeleteRolesByID(id_role int) (err error) {
	if err = db.Server().Unscoped().Where("id_role = ?", id_role).Delete(&model.Role{}).Error; err != nil {
		return err
	}
	return nil
}
