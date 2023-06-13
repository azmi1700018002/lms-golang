package rcategory

import (
	"lms/config/db"
	"lms/model"
)

func MGetAllCategory(limit, offset int) (categorys []model.Category, count int64, err error) {

	querySelect := `
		SELECT idcategory, categoryname
		FROM categories
		LIMIT $1
		OFFSET $2
	`

	rows, err := db.Server().Raw(querySelect, limit, offset).Rows()
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var category model.Category
		if err := rows.Scan(&category.IDcategory, &category.Categoryname); err != nil {
			return nil, 0, err
		}
		categorys = append(categorys, category)
	}

	queryCount := `
		SELECT COUNT(*)
		FROM categories
	`

	if err := db.Server().Raw(queryCount).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	return categorys, count, nil
}
