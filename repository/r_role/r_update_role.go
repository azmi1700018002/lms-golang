package rrole

import (
	"lms/config/db"
	"lms/model"
)

func MUpdateRolesByID(id_role int, input model.Role) (role model.Role, err error) {
	role = model.Role{
		Rolename: input.Rolename,
	}
	if err = db.Server().Where("id_role = ?", id_role).Updates(&role).Error; err != nil {
		return role, err
	}
	return role, nil
}
