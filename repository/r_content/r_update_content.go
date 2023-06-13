package rcontent

import (
	"lms/config/db"
	"lms/model"
)

func MUpdateContentByID(idcontent int, input model.Content) (content model.Content, err error) {
	content = model.Content{
		IDsection:    input.IDsection,
		ContentTitle: input.ContentTitle,
		ContentType:  input.ContentType,
		Image:        input.Image,
		Link:         input.Link,
	}
	if err = db.Server().Where("idcontent = ?", idcontent).Updates(&content).Error; err != nil {
		return content, err
	}
	return content, nil
}
