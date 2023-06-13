package rreference

import (
	"lms/config/db"
	"lms/model"
)

func MaddReference(input model.Reference) (references model.Reference, err error) {

	references = model.Reference{
		Desc1:     input.Desc1,
		Groupref1: input.Groupref1,
		Desc2:     input.Desc2,
		Groupref2: input.Groupref2,
	}

	if err = db.Server().Create(&references).Error; err != nil {
		return references, err
	}
	return references, nil
}
