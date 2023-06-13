package rrole

import (
	"lms/config/db"
	"lms/model"
)

func MGetAllRole(limit, offset int) (roles []model.Role, count int64, err error) {

	querySelect := `
		SELECT id_role, rolename
		FROM roles
		LIMIT $1
		OFFSET $2
	`

	rows, err := db.Server().Raw(querySelect, limit, offset).Rows()
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var role model.Role
		if err := rows.Scan(&role.IDRole, &role.Rolename); err != nil {
			return nil, 0, err
		}
		roles = append(roles, role)
	}

	queryCount := `
		SELECT COUNT(*)
		FROM roles
	`

	if err := db.Server().Raw(queryCount).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	return roles, count, nil
}
