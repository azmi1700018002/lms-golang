package rcontent

import (
	"errors"
	"lms/config/db"
	"lms/model"
)

func MGetAllContent(limit, offset int) (contents []model.Content, count int64, err error) {

	querySelect := `
		SELECT idcontent, idsection, content_title, content_type, image, link
		FROM contents
		LIMIT $1
		OFFSET $2
	`

	rows, err := db.Server().Raw(querySelect, limit, offset).Rows()
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var content model.Content
		if err := rows.Scan(&content.IDcontent, &content.IDsection, &content.ContentTitle, &content.ContentType, &content.Image, &content.Link); err != nil {
			return nil, 0, err
		}
		contents = append(contents, content)
	}

	queryCount := `
		SELECT COUNT(*)
		FROM contents
	`

	if err := db.Server().Raw(queryCount).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	return contents, count, nil
}

func MGetContentByIDOrm(contentID string) (content model.Content, err error) {
	if err = db.Server().First(&content, contentID).Error; err != nil {
		return model.Content{}, err
	}
	return content, nil
}

func MGetContentByID(contentID string) (content model.Content, err error) {
	if contentID == "" {
		return model.Content{}, errors.New("contentID cannot be empty")
	}

	stmt := `
		SELECT idcontent, idsection, content_title, content_type, image, link
		FROM contents

		WHERE idcontent = $1
		LIMIT 1;
	`
	row := db.Server().Raw(stmt, contentID).Row()

	if err = row.Scan(&content.IDcontent, &content.IDsection, &content.ContentTitle, &content.ContentType, &content.Image, &content.Link); err != nil {
		return model.Content{}, err
	}

	return content, nil
}
