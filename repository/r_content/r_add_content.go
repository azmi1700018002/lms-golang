package rcontent

import (
	"lms/config/db"
	"lms/model"
)

type ContentRepo struct{}

func NewContentRepository() *ContentRepo {
	return &ContentRepo{}
}

func (repo *ContentRepo) CreateContent(input model.Content) (content model.Content, err error) {
	content = model.Content{
		IDsection:    input.IDsection,
		ContentTitle: input.ContentTitle,
		ContentType:  input.ContentType,
		Image:        input.Image,
		Link:         input.Link,
	}

	if err = db.Server().Create(&content).Error; err != nil {
		return content, err
	}
	return content, nil
}
