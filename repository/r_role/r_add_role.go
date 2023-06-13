package rrole

import (
	"lms/config/db"
	"lms/model"
)

func MaddRoles(input model.Role) (roles model.Role, err error) {

	roles = model.Role{
		Rolename: input.Rolename,
	}

	if err = db.Server().Create(&roles).Error; err != nil {
		return roles, err
	}
	return roles, nil
}
