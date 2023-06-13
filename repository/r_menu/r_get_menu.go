package rmenu

import (
	"lms/config/db"
	"lms/model"
)

func MGetAllMenu(limit, offset int) (menus []model.Menu, count int64, err error) {

	querySelect := `
		SELECT id_menu, id_role
		FROM menus
		LIMIT $1
		OFFSET $2
	`

	rows, err := db.Server().Raw(querySelect, limit, offset).Rows()
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var menu model.Menu
		if err := rows.Scan(&menu.IDMenu, &menu.IDRole); err != nil {
			return nil, 0, err
		}
		menus = append(menus, menu)
	}

	queryCount := `
		SELECT COUNT(*)
		FROM menus
	`

	if err := db.Server().Raw(queryCount).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	return menus, count, nil
}
