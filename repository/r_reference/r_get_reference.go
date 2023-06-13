package rreference

import (
	"lms/config/db"
	"lms/model"
)

func MGetAllReference(limit, offset int) (references []model.Reference, count int64, err error) {

	querySelect := `SELECT idref1, desc1, groupref1, idref2, desc2, groupref2 FROM "references" LIMIT $1 OFFSET $2`

	rows, err := db.Server().Raw(querySelect, limit, offset).Rows()
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var reference model.Reference
		if err := rows.Scan(&reference.IDref1, &reference.Desc1, &reference.Groupref1, &reference.IDref2, &reference.Desc2, &reference.Groupref2); err != nil {
			return nil, 0, err
		}
		references = append(references, reference)
	}

	queryCount := `
		SELECT COUNT(*)
		FROM "references"
	`

	if err := db.Server().Raw(queryCount).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	return references, count, nil
}
